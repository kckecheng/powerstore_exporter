package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kckecheng/powerstore_exporter/powerstore"
)

func main() {
	var box *powerstore.PowerStore
	var bytes []byte
	var err error

	box, err = powerstore.New("fnm0876.drm.lab.emc.com", 443, "admin", "Password123!")
	defer box.Close()
	if err != nil {
		log.Fatal("Fail to connect to Powerstore")
	}

	// List nodes
	bytes, err = box.List("node")
	if err != nil {
		fmt.Println(err)
	}

	var ids []powerstore.ResourceID
	err = json.Unmarshal(bytes, &ids)
	if err != nil {
		fmt.Println(err)
	}
	for _, id := range ids {
		fmt.Printf("%s\n", id.ID)
	}

	// Query node metrics
	var metrics []powerstore.NodeMetric
	bytes, err = box.CollectMetrics("performance_metrics_by_node", "N1", powerstore.OneDay)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &metrics)
	for _, metric := range metrics {
		fmt.Printf("%v\n", metric)
	}
}
