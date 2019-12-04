package main

import (
	"log"
	"net/http"

	"github.com/kckecheng/powerstore_exporter/collector"
	"github.com/kckecheng/powerstore_exporter/powerstore"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	box, err := powerstore.New("fnm0876.drm.lab.emc.com", 443, "admin", "Password123!")
	defer box.Close()
	if err != nil {
		log.Fatal("Fail to connect to Powerstore")
	}

	collector.RecordMetrics(box, powerstore.FiveMins)

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
