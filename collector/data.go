package collector

import "time"

// ResourceID resouce id
type ResourceID struct {
	ID string `json:"id"`
}

// NodeMetric metric for node
type NodeMetric struct {
	NodeID                     string    `json:"node_id"`
	TimeStamp                  time.Time `json:"timestamp"`
	AvgReadLatency             float64   `json:"avg_read_latency"`
	AvgReadSize                float64   `json:"avg_read_size"`
	AvgLatency                 float64   `json:"avg_latency"`
	AvgWriteLatency            float64   `json:"avg_write_latency"`
	AvgWriteSize               float64   `json:"avg_write_size"`
	AvgReadIops                float64   `json:"avg_read_iops"`
	AvgReadBandwidth           float64   `json:"avg_read_bandwidth"`
	AvgTotalIops               float64   `json:"avg_total_iops"`
	AvgTotalBandwidth          float64   `json:"avg_total_bandwidth"`
	AvgWriteIops               float64   `json:"avg_write_iops"`
	AvgWriteBandwidth          float64   `json:"avg_write_bandwidth"`
	MaxAvgReadLatency          float64   `json:"max_avg_read_latency"`
	MaxAvgReadSize             float64   `json:"max_avg_read_size"`
	MaxAvgLatency              float64   `json:"max_avg_latency"`
	MaxAvgWriteLatency         float64   `json:"max_avg_write_latency"`
	MaxAvgWriteSize            float64   `json:"max_avg_write_size"`
	MaxReadIops                float64   `json:"max_read_iops"`
	MaxReadBandwidth           float64   `json:"max_read_bandwidth"`
	MaxTotalIops               float64   `json:"max_total_iops"`
	MaxTotalBandwidth          float64   `json:"max_total_bandwidth"`
	MaxWriteIops               float64   `json:"max_write_iops"`
	MaxWriteBandwidth          float64   `json:"max_write_bandwidth"`
	AvgCurrentLogins           float64   `json:"avg_current_logins"`
	AvgUnalignedWriteBandwidth float64   `json:"avg_unaligned_write_bandwidth"`
	AvgUnalignedReadBandwidth  float64   `json:"avg_unaligned_read_bandwidth"`
	AvgUnalignedReadIops       float64   `json:"avg_unaligned_read_iops"`
	AvgUnalignedWriteIops      float64   `json:"avg_unaligned_write_iops"`
	AvgUnalignedBandwidth      float64   `json:"avg_unaligned_bandwidth"`
	AvgUnalignedIops           float64   `json:"avg_unaligned_iops"`
	MaxCurrentLogins           float64   `json:"max_current_logins"`
	MaxUnalignedWriteBandwidth float64   `json:"max_unaligned_write_bandwidth"`
	MaxUnalignedReadBandwidth  float64   `json:"max_unaligned_read_bandwidth"`
	MaxUnalignedReadIops       float64   `json:"max_unaligned_read_iops"`
	MaxUnalignedWriteIops      float64   `json:"max_unaligned_write_iops"`
	MaxUnalignedBandwidth      float64   `json:"max_unaligned_bandwidth"`
	MaxUnalignedIops           float64   `json:"max_unaligned_iops"`
	AvgIoSize                  float64   `json:"avg_io_size"`
	MaxAvgIoSize               float64   `json:"max_avg_io_size"`
	ApplianceID                string    `json:"appliance_id"`
	RepeatCount                float64   `json:"repeat_count"`
	Entity                     string    `json:"entity"`
}
