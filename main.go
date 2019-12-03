package main

import (
	"fmt"
	"log"

	"github.com/kckecheng/powerstore_exporter/collector"
)

func main() {
	box, err := collector.New("fnm0876.drm.lab.emc.com", 443, "admin", "Password123!")
	defer box.Close()
	if err != nil {
		log.Fatal("Fail to connect to Powerstore")
	}

	content, err := box.CollectMetrics("performance_metrics_by_cluster", "0", collector.OneDay)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%v\n", content)
}
