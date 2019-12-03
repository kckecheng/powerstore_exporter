package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kckecheng/powerstore_exporter/collector"
)

func main() {
	var box *collector.PowerStore
	var bytes []byte
	var err error

	box, err = collector.New("fnm0876.drm.lab.emc.com", 443, "admin", "Password123!")
	defer box.Close()
	if err != nil {
		log.Fatal("Fail to connect to Powerstore")
	}

	// List nodes
	bytes, err = box.List("node")
	if err != nil {
		fmt.Println(err)
	}

	var ids []collector.ResourceID
	err = json.Unmarshal(bytes, &ids)
	if err != nil {
		fmt.Println(err)
	}
	for _, id := range ids {
		fmt.Printf("%s\n", id.ID)
	}

	// Query node metrics
	var metrics []collector.NodeMetric
	bytes, err = box.CollectMetrics("performance_metrics_by_node", "N1", collector.OneDay)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(bytes, &metrics)
	fmt.Printf("%v\n", metrics)
}
