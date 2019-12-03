package collector

import (
	"crypto/tls"
	"fmt"
	"log"

	"github.com/parnurzeal/gorequest"
)

// PowerStore the array
/*
	prefix: the REST API prefix
	baseReq: the base gorequest.SuperAgent object, all requests should be a clone
	of it as the starting point
*/
type PowerStore struct {
	Address  string
	Port     uint
	User     string
	Password string
	baseReq  *gorequest.SuperAgent
	prefix   string
	token    string
}

// Interval PowerStore supported metric query interval
type Interval string

// Common used interval definitions supported by PowerStore
const (
	TwentySec Interval = "Twenty_Sec"
	FiveMins  Interval = "Five_Mins"
	OneHour   Interval = "One_Hour"
	OneDay    Interval = "One_Day"
)

// reqBody: payload for collecting metrics
type reqBody struct {
	Entity   string   `json:"entity"`
	EntityID string   `json:"entity_id"`
	Interval Interval `json:"interval"`
}

// New init the PowerStore object
func New(address string, port uint, user string, password string) (*PowerStore, error) {
	box := PowerStore{
		Address:  address,
		Port:     port,
		User:     user,
		Password: password,
		prefix:   "/api/rest",
		baseReq:  nil,
		token:    "",
	}
	err := box.connect()
	if err != nil {
		return nil, err
	}

	return &box, nil
}

func logErrors(errs []error) bool {
	if len(errs) > 0 {
		for _, err := range errs {
			log.Println("Hit errors: ")
			log.Println(err)
		}
		return true
	}
	return false
}

// BuildURL create the full url for requests
func (box *PowerStore) buildURL(resource string) string {
	url := fmt.Sprintf("https://%s:%d%s%s", box.Address, box.Port, box.prefix, resource)
	return url
}

func (box *PowerStore) updateHeader() {
	box.baseReq.Set("Accept", "application/json")
	box.baseReq.Set("Content-Type", "application/json")
	if box.token != "" {
		box.baseReq.Set("DELL-EMC-TOKEN", box.token)
	}
}

// Connect to the box
func (box *PowerStore) connect() error {
	baseReq := gorequest.New().SetBasicAuth(box.User, box.Password)
	baseReq.TLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	box.baseReq = baseReq
	box.updateHeader()

	request := baseReq.Clone()
	resp, _, errs := request.Get(box.buildURL("/login_session")).End()
	ok := logErrors(errs)
	if ok {
		return fmt.Errorf("Fail to connect to the PowerStore array")
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Internal error, incorrect response")
	}

	// fmt.Println("Dump headers")
	// for k, v := range resp.Header {
	//   fmt.Printf("%s: %s\n", k, v)
	// }

	token, ok := resp.Header["Dell-Emc-Token"]
	if !ok {
		return fmt.Errorf("Cannot locate Dell-Emc-Token in the login response header")
	}
	box.token = token[0]
	box.updateHeader()
	return nil
}

// CollectMetrics collect metrics from PowerStore
func (box *PowerStore) CollectMetrics(entity string, id string, interval Interval) ([]byte, error) {
	request := box.baseReq.Clone()
	payload := reqBody{entity, id, interval}
	resp, body, errs := request.Post(box.buildURL("/metrics/generate")).SendStruct(payload).End()
	ok := logErrors(errs)
	if ok {
		return nil, fmt.Errorf("Fail to connect to the PowerStore array")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Internal error, incorrect response")
	}

	// fmt.Println("Dump response body")
	// fmt.Printf("%T\n", body)
	// fmt.Printf("%v\n", body)
	return []byte(body), nil
}

// List resouces such as appliance, cluster, volume
func (box *PowerStore) List(resource string) ([]byte, error) {
	request := box.baseReq.Clone()
	resp, body, errs := request.Get(box.buildURL(resource)).End()
	ok := logErrors(errs)
	if ok {
		return nil, fmt.Errorf("Fail to list resouce")
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("Internal error, incorrect response")
	}
	return []byte(body), nil
}

// Close the connection to the PowerStore
func (box *PowerStore) Close() error {
	request := box.baseReq.Clone()
	resp, _, errs := request.Get(box.buildURL("/logout")).End()
	ok := logErrors(errs)
	if ok {
		return fmt.Errorf("Fail to logout due to internal errors")
	}

	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		return fmt.Errorf("Invalid logout response")
	}
	box.baseReq = nil
	box.token = ""
	return nil
}
