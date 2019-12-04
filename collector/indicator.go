package collector

import (
	"github.com/prometheus/client_golang/prometheus"
)

var (
	avgReadLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_read_latency",
			Help: "Average read latency",
		},
		[]string{"appliance", "node"},
	)
	avgReadSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_read_size",
			Help: "Average read size",
		},
		[]string{"appliance", "node"},
	)
	avgLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_latency",
			Help: "Average latency",
		},
		[]string{"appliance", "node"},
	)
	avgWriteLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_write_latency",
			Help: "Average write latency",
		},
		[]string{"appliance", "node"},
	)
	avgWriteSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_write_size",
			Help: "Average write size",
		},
		[]string{"appliance", "node"},
	)
	avgReadIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_read_iops",
			Help: "Average read IOPS",
		},
		[]string{"appliance", "node"},
	)
	avgReadBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_read_bandwidth",
			Help: "Average read bandwidth",
		},
		[]string{"appliance", "node"},
	)
	avgTotalIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_total_iops",
			Help: "Average total IOPS",
		},
		[]string{"appliance", "node"},
	)
	avgTotalBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_total_bandwidth",
			Help: "Average total bandwidth",
		},
		[]string{"appliance", "node"},
	)
	avgWriteIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_write_iops",
			Help: "Average write IOPS",
		},
		[]string{"appliance", "node"},
	)
	avgWriteBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_write_bandwidth",
			Help: "Average write bandwidth",
		},
		[]string{"appliance", "node"},
	)
	maxAvgReadLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_avg_read_latency",
			Help: "Max average read latency",
		},
		[]string{"appliance", "node"},
	)
	maxAvgReadSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_avg_read_size",
			Help: "Max average read size",
		},
		[]string{"appliance", "node"},
	)
	maxAvgLatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_avg_latency",
			Help: "Max average latency",
		},
		[]string{"appliance", "node"},
	)
	maxAvgWritelatency = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_avg_write_latency",
			Help: "Max average write latency",
		},
		[]string{"appliance", "node"},
	)
	maxAvgWriteSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_avg_write_size",
			Help: "Max average write size",
		},
		[]string{"appliance", "node"},
	)
	maxReadIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_read_iops",
			Help: "Max read IOPS",
		},
		[]string{"appliance", "node"},
	)
	maxReadBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_read_bandwidth",
			Help: "Max read bandwidth",
		},
		[]string{"appliance", "node"},
	)
	maxTotalIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_total_iops",
			Help: "Max total IOPS",
		},
		[]string{"appliance", "node"},
	)
	maxTotalBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_total_bandwidth",
			Help: "Max total bandwidth",
		},
		[]string{"appliance", "node"},
	)
	maxWriteIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_write_iops",
			Help: "Max write IOPS",
		},
		[]string{"appliance", "node"},
	)
	maxWriteBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_write_bandwidth",
			Help: "Max write bandwidth",
		},
		[]string{"appliance", "node"},
	)
	avgCurrentLogins = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_current_logins",
			Help: "Average logins",
		},
		[]string{"appliance", "node"},
	)
	maxCurrentLogins = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_current_logins",
			Help: "Max logins",
		},
		[]string{"appliance", "node"},
	)
	avgUnalignedWriteBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_unaligned_write_bandwidth",
			Help: "Average write bandwidth",
		},
		[]string{"appliance", "node"},
	)
	avgUnalignedReadBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_unaligned_read_bandwidth",
			Help: "Average read bandwidth",
		},
		[]string{"appliance", "node"},
	)
	avgUnalignedReadIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_unaligned_read_iops",
			Help: "Average read IOPS",
		},
		[]string{"appliance", "node"},
	)
	avgUnalignedWriteIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_unaligned_write_iops",
			Help: "Average write IOPS",
		},
		[]string{"appliance", "node"},
	)
	avgUnalignedBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_unaligned_bandwidth",
			Help: "Average unaligned bandwidth",
		},
		[]string{"appliance", "node"},
	)
	avgUnalignedIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_unaligned_iops",
			Help: "Average unaligned IOPS",
		},
		[]string{"appliance", "node"},
	)
	maxUnalignedWriteBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_unaligned_write_bandwidth",
			Help: "Max write bandwidth",
		},
		[]string{"appliance", "node"},
	)
	maxUnalignedReadBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_unaligned_read_bandwidth",
			Help: "Max read bandwidth",
		},
		[]string{"appliance", "node"},
	)
	maxUnalignedReadIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_unaligned_read_iops",
			Help: "Max read IOPS",
		},
		[]string{"appliance", "node"},
	)
	maxUnalignedWriteIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_unaligned_write_iops",
			Help: "Max write IOPS",
		},
		[]string{"appliance", "node"},
	)
	maxUnalignedBandwidth = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_unaligned_bandwidth",
			Help: "Max unaligned bandwidth",
		},
		[]string{"appliance", "node"},
	)
	maxUnalignedIops = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_unaligned_iops",
			Help: "Max unaligned IOPS",
		},
		[]string{"appliance", "node"},
	)
	avgIoSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "avg_io_size",
			Help: "Average IO size",
		},
		[]string{"appliance", "node"},
	)
	maxAvgIoSize = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "max_avg_io_size",
			Help: "Max average IO size",
		},
		[]string{"appliance", "node"},
	)
)

func init() {
	prometheus.MustRegister(avgReadLatency)
	prometheus.MustRegister(avgReadSize)
	prometheus.MustRegister(avgLatency)
	prometheus.MustRegister(avgWriteLatency)
	prometheus.MustRegister(avgWriteSize)
	prometheus.MustRegister(avgReadIops)
	prometheus.MustRegister(avgReadBandwidth)
	prometheus.MustRegister(avgTotalIops)
	prometheus.MustRegister(avgTotalBandwidth)
	prometheus.MustRegister(avgWriteIops)
	prometheus.MustRegister(avgWriteBandwidth)
	prometheus.MustRegister(maxAvgReadLatency)
	prometheus.MustRegister(maxAvgReadSize)
	prometheus.MustRegister(maxAvgLatency)
	prometheus.MustRegister(maxAvgWritelatency)
	prometheus.MustRegister(maxAvgWriteSize)
	prometheus.MustRegister(maxReadIops)
	prometheus.MustRegister(maxReadBandwidth)
	prometheus.MustRegister(maxTotalIops)
	prometheus.MustRegister(maxTotalBandwidth)
	prometheus.MustRegister(maxWriteIops)
	prometheus.MustRegister(maxWriteBandwidth)
	prometheus.MustRegister(avgCurrentLogins)
	prometheus.MustRegister(maxCurrentLogins)
	prometheus.MustRegister(avgUnalignedWriteBandwidth)
	prometheus.MustRegister(avgUnalignedReadBandwidth)
	prometheus.MustRegister(avgUnalignedReadIops)
	prometheus.MustRegister(avgUnalignedWriteIops)
	prometheus.MustRegister(avgUnalignedBandwidth)
	prometheus.MustRegister(avgUnalignedIops)
	prometheus.MustRegister(maxUnalignedWriteBandwidth)
	prometheus.MustRegister(maxUnalignedReadBandwidth)
	prometheus.MustRegister(maxUnalignedReadIops)
	prometheus.MustRegister(maxUnalignedWriteIops)
	prometheus.MustRegister(maxUnalignedBandwidth)
	prometheus.MustRegister(maxUnalignedIops)
	prometheus.MustRegister(avgIoSize)
	prometheus.MustRegister(maxAvgIoSize)
}
