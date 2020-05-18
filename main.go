package main

import (
	"fmt"
	"net/http"

	"github.com/kckecheng/powerstore_exporter/common"
	"github.com/kckecheng/powerstore_exporter/exporter"
	"github.com/kckecheng/powerstore_exporter/powerstore"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var ps *powerstore.PowerStore
var exporters []*exporter.Exporter

func main() {
	common.CfgInit()

	ps = powerstore.NewPowerStore(
		common.Config.PowerStore.Address,
		common.Config.PowerStore.User,
		common.Config.PowerStore.Password,
	)

	for _, t := range common.Config.Exporter.Resources {
		exporters = append(exporters, exporter.New(ps, t))
	}

	reg := prometheus.NewRegistry()
	for _, e := range exporters {
		reg.MustRegister(e)
	}

	// Add process and go internal stats information
	// reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}), prometheus.NewGoCollector())

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "PowerStore Exporter: access /metrics for data")
	})

	http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{}))

	common.Logger.Infof("Start PowerStore Exporter at port %d", common.Config.Exporter.Port)
	common.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", common.Config.Exporter.Port), nil))
}
