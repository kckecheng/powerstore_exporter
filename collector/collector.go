package collector

import (
	"fmt"
	"log"
	"net/http"

	"github.com/kckecheng/powerstore_exporter/powerstore"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// cpuTemp.Set(65.3)
	// hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	// The Handler function provides a default handler to expose metrics
	// via an HTTP server. "/metrics" is the usual endpoint for that.
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))

}

// RecordMetrics collect and publish metrics
func RecordMetrics(box *powerstore.PowerStore, interval powerstore.Interval) {
	ids, err := box.ListNodes()
	if err != nil {
		log.Fatal("Cannot get any node on the PowerStore")
	}

	go func() {
		for {
			// Sleep - TBD
			// time.Sleep()
			for _, id := range ids {
				metric, err := box.GetLatestNodeMetric(id, interval)
				if err != nil {
					log.Fatal(fmt.Sprintf("Cannot get any metric from node %s", id))
				}

				// Update gauge
				avgReadLatency.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgReadLatency)
				avgReadSize.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgReadSize)
				avgLatency.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgLatency)
				avgWriteLatency.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgWriteLatency)
				avgWriteSize.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgWriteSize)
				avgReadIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgReadIops)
				avgReadBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgReadBandwidth)
				avgTotalIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgTotalIops)
				avgTotalBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgTotalBandwidth)
				avgWriteIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgWriteIops)
				avgWriteBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgWriteBandwidth)
				maxAvgReadLatency.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxAvgReadLatency)
				maxAvgReadSize.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxAvgReadSize)
				maxAvgLatency.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxAvgLatency)
				maxAvgWritelatency.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxAvgWriteLatency)
				maxAvgWriteSize.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxAvgWriteSize)
				maxReadIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxReadIops)
				maxReadBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxReadBandwidth)
				maxTotalIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxTotalIops)
				maxTotalBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxTotalBandwidth)
				maxWriteIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxWriteIops)
				maxWriteBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxWriteBandwidth)
				avgCurrentLogins.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgCurrentLogins)
				maxCurrentLogins.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxCurrentLogins)
				avgUnalignedWriteBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgUnalignedWriteBandwidth)
				avgUnalignedReadBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgUnalignedReadBandwidth)
				avgUnalignedReadIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgUnalignedReadIops)
				avgUnalignedWriteIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgUnalignedWriteIops)
				avgUnalignedBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgUnalignedBandwidth)
				avgUnalignedIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgUnalignedIops)
				maxUnalignedWriteBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxUnalignedWriteBandwidth)
				maxUnalignedReadBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxUnalignedReadBandwidth)
				maxUnalignedReadIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxUnalignedReadIops)
				maxUnalignedWriteIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxUnalignedWriteIops)
				maxUnalignedBandwidth.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxUnalignedBandwidth)
				maxUnalignedIops.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxUnalignedIops)
				avgIoSize.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.AvgIoSize)
				maxAvgIoSize.With(prometheus.Labels{"appliance": metric.ApplianceID, "node": metric.NodeID}).Set(metric.MaxAvgIoSize)
			}
		}
	}()
}
