package heart_beat

import (
	"github.com/meson-network/peer_common/api"
	"github.com/meson-network/peer_common/info"
)

const HEART_BEAT_INTERVAL_SECS = 30

//tasks
const TASK_DOWNLOAD = "task_download"

// @Description Msg_Req_HeartBeat
type Msg_Req_HeartBeat struct {
	Node_id             string `json:"node_id"`             //required
	Downloading_threads uint64 `json:"downloading_threads"` //required
	Port                uint64 `json:"port"`
	Stor_total_bytes    int64  `json:"stor_total_bytes"`
	Stor_used_bytes     int64  `json:"stor_used_bytes"`
	info.HardwareInfo
}

// @Description Msg_Resp_HeartBeat
type Msg_Resp_HeartBeat struct {
	Server_unixtime int64    `json:"server_unixtime"`
	Tasks           []string `json:"tasks"`
	api.API_META_STATUS
}
