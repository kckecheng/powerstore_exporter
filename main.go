package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kckecheng/powerstore_exporter/collector"
	"github.com/kckecheng/powerstore_exporter/powerstore"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gopkg.in/yaml.v2"
)

type powerStoreCfg struct {
	PowerStore struct {
		Address  string `yaml:"address"`
		Port     int    `yaml:"port"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
	} `yaml:"powerstore"`
	Interval string `yaml:"interval"`
}

func main() {
	// Read and parse PowerStore configuration file
	var cfg powerStoreCfg
	handler, err := ioutil.ReadFile("powerstore_exporter.yml")
	if err != nil {
		log.Fatal("Fail to read configuration file")
	}

	err = yaml.Unmarshal(handler, &cfg)
	if err != nil {
		log.Fatal("Fail to parse the configuration file")
	}
	// fmt.Printf("%+v\n", cfg)

	// Connect to the PowerStore
	box, err := powerstore.New(cfg.PowerStore.Address, cfg.PowerStore.Port, cfg.PowerStore.User, cfg.PowerStore.Password)
	defer box.Close()
	if err != nil {
		log.Fatal("Fail to connect to Powerstore")
	}

	// Collect the latest node metrics
	// collector.RecordMetrics(box, powerstore.FiveMins)
	collector.RecordMetrics(box, powerstore.Interval(cfg.Interval))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `<html><body><h1>Metrics for PowerStore, check <a href="/metrics">metrics</a> for details<h1></body></html>`)
	})

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
