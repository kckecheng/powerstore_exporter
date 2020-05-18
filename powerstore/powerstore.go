package powerstore

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/kckecheng/powerstore_exporter/common"
)

// PowerStore array object
type PowerStore struct {
	APIEndPoint string
	User        string
	Password    string
	client      *resty.Client // Not exposed
}

// NewPowerStore init a PowerStore object
func NewPowerStore(address, user, password string) *PowerStore {
	ps := PowerStore{}
	ps.APIEndPoint = fmt.Sprintf("https://%s/api/rest", address)
	ps.User = user
	ps.Password = password

	// Populate Conn
	common.Logger.Infof("Connect to PowerStore %s", address)
	ps.login()
	return &ps
}

func (ps *PowerStore) login() {
	client := newClient(ps.User, ps.Password)

	common.Logger.Debug("Access /login_session to get token")
	resp, err := get(client, ps.url("/login_session"), nil, nil)
	if err != nil {
		common.Logger.Fatalf("Fail to login: %s", err.Error())
	}

	common.Logger.Debugf("Locate token with header key Dell-Emc-Token")
	token, exists := resp.Header()["Dell-Emc-Token"]
	if exists {
		common.Logger.Debugf("Get login session token: %s", token[0])
		ps.client = client.SetHeader("DELL-EMC-TOKEN", token[0])
		common.Logger.Debugf("Current header: %+v", client.Header)
	} else {
		common.Logger.Fatalf("Fail to login since no token can be found")
	}
}

// Logout Log out the connection
func (ps *PowerStore) Logout() {
	resp, err := ps.Post("/logout", nil)
	if err != nil {
		common.Logger.Errorf("Fail to logout, skip silently")
	} else {
		if resp.IsError() {
			common.Logger.Errorf("Logout with an incorrect return status: %s - skip silently", resp.Status())
		} else {
			common.Logger.Info("Successfully logout PowerStore")
		}
	}
}

func (ps PowerStore) url(uri string) string {
	url := fmt.Sprintf("%s%s", ps.APIEndPoint, uri)
	common.Logger.Debugf("Full URL: %s", url)
	return url
}

// Get PowerStore HTTP GET encapsulation
func (ps PowerStore) Get(uri string, headers map[string]string, params map[string]string) (*resty.Response, error) {
	common.ReqCounter <- 1
	resp, err := get(ps.client, ps.url(uri), headers, params)
	<-common.ReqCounter
	return resp, err
}

// Post PowerStore HTTP POST encapsulation
func (ps PowerStore) Post(uri string, body map[string]string) (*resty.Response, error) {
	common.ReqCounter <- 1
	resp, err := post(ps.client, ps.url(uri), body)
	<-common.ReqCounter
	return resp, err
}

// ListResources List PowerStore resources
func (ps PowerStore) ListResources(resType string) []map[string]string {
	res := map[string]string{
		"cluster":     "id,name",
		"appliance":   "id,name",
		"node":        "appliance_id,id,name",
		"fc_port":     "appliance_id,node_id,id,name,wwn,current_speed",         // current_speed will be "" when link is down
		"eth_port":    "appliance_id,node_id,id,name,mac_address,current_speed", // current_speed will be "" when link is down
		"volume":      "id,name,type",                                           // type: Primary, Clone, Snapshot
		"file_system": "id,name,filesystem_type",                                // filesystem_type: Primary, Snapshot
	}
	fields, exists := res[resType]
	if !exists {
		common.Logger.Fatalf("Resource type %s does not exist", resType)
	}
	params := map[string]string{"select": fields}

	// To be enhanced - implement the same within other module/func???
	if resType == "fc_port" || resType == "eth_port" {
		params["current_speed"] = "not.is.null"
	}
	if resType == "volume" {
		params["type"] = "eq.Primary"
	}
	if resType == "file_system" {
		params["filesystem_type"] = "eq.Primary"
	}

	common.Logger.Infof("List resource %s", resType)
	ret := []map[string]string{}
	var headers map[string]string
	remain := true
	for remain {
		resp, err := ps.Get("/"+resType, headers, params)
		if err != nil {
			common.Logger.Fatalf("Fail to get %s information: %s", resType, err.Error())
		}

		if resp.IsError() {
			common.Logger.Fatalf("Unexpected HTTP return code: %s", resp.Status())
		}

		if resp.StatusCode() != 206 {
			remain = false
		} else {
			common.Logger.Debugf("Resource %s has more than 100 instances", resType)
			headers, remain = generateRangeHeader(resp.Header())
		}

		var partial []map[string]string
		err = json.Unmarshal(resp.Body(), &partial)
		if err != nil {
			common.Logger.Fatalf("Fail to extract %s information: %s", resType, err.Error())
		}
		ret = append(ret, partial...)
	}

	common.Logger.Debugf("Get resources %v", ret)
	return ret
}

// GetLatestMetric Get the latest metric
func (ps PowerStore) GetLatestMetric(resType string, resID string, rollup bool) map[string]float64 {
	entities := map[string]string{
		"cluster":     "performance_metrics_by_cluster",
		"appliance":   "performance_metrics_by_appliance",
		"node":        "performance_metrics_by_node",
		"fc_port":     "performance_metrics_by_fe_fc_port",
		"eth_port":    "performance_metrics_by_fe_eth_port",
		"volume":      "performance_metrics_by_volume",
		"file_system": "performance_metrics_by_file_system",
	}

	entity, exists := entities[resType]
	if !exists {
		common.Logger.Fatalf("Resource type %s does not exist", resType)
	}

	interval := "Twenty_Sec"
	if rollup {
		interval = "Five_Mins"
	}

	body := map[string]string{
		"entity":    entity,
		"entity_id": resID,
		"interval":  interval,
	}

	common.Logger.Debugf("Get latest metrics: %v", body)
	resp, err := ps.Post("/metrics/generate", body)
	if err != nil {
		common.Logger.Errorf("Fail to get metrics for %s %s with interval %s: %s", resType, resID, interval, err.Error())
		return nil
	}

	if resp.IsError() {
		common.Logger.Warningf("HTTP error %s is hit while getting metric for %s with interval %s", resp.Status(), resID, interval)
		return nil
	}

	var metrics []map[string]interface{}
	common.Logger.Debugf("Decode metrics ...")
	err = json.Unmarshal(resp.Body(), &metrics)
	if err != nil {
		common.Logger.Warningf("Fail to decode metric from response: %s", err.Error())
		return nil
	}
	common.Logger.Debugf("Get metrics: %v", metrics)
	return extractLatestMetric(metrics)
}

// DetectMetricNames detect supported metrics
func (ps PowerStore) DetectMetricNames(resType string, resID string, rollup bool) []string {
	names := []string{}
	metrics := ps.GetLatestMetric(resType, resID, rollup)
	if metrics == nil {
		common.Logger.Debugf("No metric can be found for %s %s (rollup: %v)", resType, resID, rollup)
		return nil
	}
	for k := range metrics {
		names = append(names, k)
	}
	common.Logger.Debugf("Get metric names: %v", names)
	return names
}

func extractLatestMetric(metrics []map[string]interface{}) map[string]float64 {
	if metrics == nil || len(metrics) == 0 {
		common.Logger.Debugf("Metrics is empty")
		return nil
	}

	metricLatestRaw := metrics[len(metrics)-1]
	metric := map[string]float64{}

	for k, v := range metricLatestRaw {
		if k == "repeat_count" {
			continue
		}
		switch v.(type) {
		case int:
			metric[k] = float64(v.(int))
		case int32:
			metric[k] = float64(v.(int32))
		case int64:
			metric[k] = float64(v.(int64))
		case float32:
			metric[k] = float64(v.(float32))
		case float64:
			metric[k] = v.(float64)
		default:
			// Delete string values
		}
	}
	common.Logger.Debugf("Extract the latest metric: %v", metric)
	return metric
}

func newClient(user, password string) *resty.Client {
	client := resty.New().SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true}).SetBasicAuth(user, password)
	client = client.SetHeader("Content-Type", "application/json").SetHeader("Accept", "application/json")
	common.Logger.Debugf("Create HTTP client: %v", client)
	return client
}

func get(client *resty.Client, url string, headers map[string]string, params map[string]string) (*resty.Response, error) {
	common.Logger.Debugf("Access %s with headers %v and params %v", url, headers, params)
	req := client.R().SetHeaders(headers).SetQueryParams(params)
	return req.Get(url)
}

func post(client *resty.Client, url string, body map[string]string) (*resty.Response, error) {
	common.Logger.Debugf("Post body %v to %s", body, url)
	req := client.R()
	if body != nil {
		return req.SetBody(body).Post(url)
	}
	return req.Post(url)
}

func generateRangeHeader(respHeader http.Header) (map[string]string, bool) {
	crange, exists := respHeader["Content-Range"]
	if !exists {
		common.Logger.Fatalf("Response code indicates range based response, but there is no range informaton in headers")
	}

	seg := strings.Split(crange[0], "/")
	be := strings.Split(seg[0], "-")
	if len(be) == 1 {
		return nil, false
	}

	var step int64 = 100
	total, _ := strconv.ParseInt(seg[1], 10, 64)
	last, _ := strconv.ParseInt(be[1], 10, 64)

	if last >= total-1 {
		return nil, false
	}

	ns := last + 1
	ne := last + step

	if ns >= total-1 {
		ns = total - 1
		ne = total - 1
	} else if ne >= total-1 {
		ne = total - 1
	}

	headers := map[string]string{
		"Range": fmt.Sprintf("%d-%d", ns, ne),
	}
	common.Logger.Debugf("Create new range based header: %v", headers)
	return headers, true
}
