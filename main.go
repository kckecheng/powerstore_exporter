package main

import (
	"fmt"
	"log"

	"github.com/kckecheng/powerstore_exporter/powerstore"
)

func main() {
	box, err := powerstore.New("fnm0876.drm.lab.emc.com", 443, "admin", "Password123!")
	defer box.Close()
	if err != nil {
		log.Fatal("Fail to connect to Powerstore")
	}

	ids, err := box.ListNodes()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", ids)

	metric, err := box.GetLatestNodeMetric("N1", powerstore.FiveMins)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", *metric)
}
