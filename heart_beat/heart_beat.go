package heart_beat

import (
	"github.com/meson-network/peer_common/api"
)

const HEART_BEAT_INTERVAL_SECS = 30

// @Description Msg_Req_HeartBeat
type Msg_Req_HeartBeat struct {
	Node_id      string `json:"node_id"`      //required
	Port         string `json:"port"`         //required
	Storage_port string `json:"storage_port"` //required
	Version      string `json:"version"`      //required
	Access_key   string `json:"access_key"`   //required
	Initial      bool   `json:"initial"`      //required
}

// @Description Msg_Resp_HeartBeat
type Msg_Resp_HeartBeat struct {
	Server_unixtime int64 `json:"server_unixtime"`
	api.API_META_STATUS
}

// @Description Msg_Resp_HeartBeatCallback
type Msg_Resp_HeartBeatCallback struct {
	api.API_META_STATUS

	Stor_total_bytes int64 `json:"stor_total_bytes"`
	Stor_used_bytes  int64 `json:"stor_used_bytes"`

	HardwareInfo
}

// @Description HardwareInfo
type HardwareInfo struct {
	Cpu              string  `json:"cpu" redis:"cpu" structs:"cpu"`
	Cpu_count        int64   `json:"cpu_count" redis:"cpu_count" structs:"cpu_count"`
	Op_sys           string  `json:"op_sys" redis:"op_sys" structs:"op_sys"`
	Cpu_percentage   float64 `json:"cpu_percentage" redis:"cpu_percentage" structs:"cpu_percentage"`
	Mem_total_bytes  int64   `json:"mem_total_bytes" redis:"mem_total_bytes" structs:"mem_total_bytes"`
	Mem_used_bytes   int64   `json:"mem_used_bytes" redis:"mem_used_bytes" structs:"mem_used_bytes"`
	Disk_total_bytes int64   `json:"disk_total_bytes" redis:"disk_total_bytes" structs:"disk_total_bytes"`
	Disk_used_bytes  int64   `json:"disk_used_bytes" redis:"disk_used_bytes" structs:"disk_used_bytes"`
}
