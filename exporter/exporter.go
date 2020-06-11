package exporter

import (
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/kckecheng/powerstore_exporter/common"
	"github.com/kckecheng/powerstore_exporter/powerstore"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	labels []string = []string{"id", "name"}
)

// Exporter Cluster metrics exporter
type Exporter struct {
	resType   string              // Resource type: cluster, appliance, node, fc_port, eth_port, volume, file_system
	resources []map[string]string // Available resources
	mutex     sync.Mutex          // Mutext control resource updates
	metrics   []string            // metric names
	descs     map[string]*prometheus.Desc
	ps        *powerstore.PowerStore
}

func (e *Exporter) initDesc() {
	for _, name := range e.metrics {
		common.Logger.Infof("Init metric definition: %s", e.metrics)
		desc := prometheus.NewDesc(name, strings.ReplaceAll(name, "_", " "), labels, nil)
		e.descs[name] = desc
	}
}

// Describe define metric desc
func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	e.initDesc()
	for _, desc := range e.descs {
		ch <- desc
	}
}

// Collect logic to collect metrics
func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	resType := e.resType
	if resType == "volume" || resType == "file_system" {
		e.mutex.Lock()
		defer e.mutex.Unlock()
	}

	var wg sync.WaitGroup
	wg.Add(len(e.resources))

	common.Logger.Infof("Start collecting metrics for %s", resType)
	for _, res := range e.resources {
		go func(resID string, resName string) {
			defer wg.Done()

			metric := e.ps.GetLatestMetric(resType, resID, common.Config.Exporter.Rollup)
			if metric == nil || len(metric) == 0 {
				return
			}
			for k, v := range metric {
				desc := e.descs[fmt.Sprintf("%s_%s", resType, k)]
				ch <- prometheus.MustNewConstMetric(
					desc,
					prometheus.GaugeValue,
					v,
					resID,
					resName,
				)
			}
		}(res["id"], res["name"])
	}

	wg.Wait()
	common.Logger.Infof("Complete collecting metrics for %s", resType)
}

func detectMetricNames(ps *powerstore.PowerStore, resType, resID string) []string {
	names := ps.DetectMetricNames(resType, resID, common.Config.Exporter.Rollup)
	if names == nil || len(names) == 0 {
		common.Logger.Warnf("No metric has been defined for %s", resType)
		return nil
	}

	common.Logger.Debugf("Get original metric names: %v", names)
	for index, name := range names {
		names[index] = fmt.Sprintf("%s_%s", resType, name)
	}
	common.Logger.Debugf("Get new metric names: %v", names)
	return names
}

// New init an exporter
func New(ps *powerstore.PowerStore, resType string) *Exporter {
	common.Logger.Infof("Init resources dynamically for %s", resType)
	resources := ps.ListResources(resType)
	if resources == nil || len(resources) == 0 {
		common.Logger.Warnf("No %s resource exists", resType)
	}

	// when metric names cannot be determined during initial, it needs to be refreshed periodicaly
	names := []string{}
	for _, res := range resources {
		names = detectMetricNames(ps, resType, res["id"])
		if names != nil && len(names) > 0 {
			break
		}
	}

	e := Exporter{
		resType:   resType,
		resources: resources,
		mutex:     sync.Mutex{},
		metrics:   names,
		descs:     map[string]*prometheus.Desc{},
		ps:        ps,
	}

	// Refresh resource type periodically for volume and file system
	if e.resType == "volume" || e.resType == "file_system" {
		common.Logger.Infof("Register periodically resource update hook for %s every %d seconds", e.resType, common.ResUpdateInterval)
		ticker := time.NewTicker(common.ResUpdateInterval * time.Second)
		go func() {
			for {
				select {
				case <-ticker.C:
					e.mutex.Lock()
					common.Logger.Infof("Update resource %s", resType)
					newResources := ps.ListResources(resType)
					if newResources == nil || len(newResources) == 0 {
						common.Logger.Errorf("No %s resource is found during periodical update, skip update", resType)
					} else {
						e.resources = newResources
						common.Logger.Infof("Successfully update %s resource during periodical update", resType)

						// when metric names are not initialized previously, try to initialize them
						if e.metrics == nil || len(e.metrics) == 0 {
							common.Logger.Infof("Since there is no metric found for %s previously, try to reinitialize", e.resType)
							names := []string{}
							for _, res := range resources {
								names = detectMetricNames(ps, resType, res["id"])
								if names != nil && len(names) > 0 {
									break
								}
							}
							if names != nil && len(names) > 0 {
								e.metrics = names
								e.initDesc()
							}
						}
					}
					e.mutex.Unlock()
				}
			}
		}()
	}

	return &e
}
