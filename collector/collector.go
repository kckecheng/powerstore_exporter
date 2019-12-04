package collector

import (
	"fmt"
	"log"
	"time"

	"github.com/kckecheng/powerstore_exporter/powerstore"
	"github.com/prometheus/client_golang/prometheus"
)

// RecordMetrics collect and publish metrics
func RecordMetrics(box *powerstore.PowerStore, interval powerstore.Interval) {
	go func() {
		ids, err := box.ListNodes()
		if err != nil {
			log.Fatal("Cannot get any node on the PowerStore")
		}

		for {
			// Increase the counter each time a query is performed against the backend
			queryTotal.Inc()

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

			// Sleep - TBD
			// time.Sleep()
			var h int
			switch interval {
			case powerstore.TwentySec:
				h = 20
			case powerstore.FiveMins:
				h = 300
			case powerstore.OneHour:
				h = 3600
			case powerstore.OneDay:
				h = 3600 * 24
			default:
				h = 300
			}
			time.Sleep(time.Second * time.Duration(h))
		}
	}()
}
