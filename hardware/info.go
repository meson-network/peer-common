package hardware

// @Description HardwareInfo
type HardwareInfo struct {
	Cpu_cores      uint64  `json:"cpu_cores"`
	Cpu_percentage float64 `json:"cpu_percentage"`
	Mem_total      uint64  `json:"mem_total"`
	Mem_used       uint64  `json:"mem_used"`
	Op_sys         string  `json:"op_sys"`
	Disk_total     uint64  `json:"disk_total"`
	Disk_available uint64  `json:"disk_available"`
}
