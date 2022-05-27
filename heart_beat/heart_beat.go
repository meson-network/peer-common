package heart_beat

import (
	"github.com/meson-network/peer_common/api"
	"github.com/meson-network/peer_common/info"
)

const HEART_BEAT_INTERVAL_SECS = 30

// @Description Msg_Req_HeartBeat
type Msg_Req_HeartBeat struct {
	Node_id string `json:"node_id"` //required
	//Downloading_threads uint64 `json:"downloading_threads"` //required
	Port    int    `json:"port"`
	Version string `json:"version"`
	//Stor_total_bytes    int64  `json:"stor_total_bytes"`
	//Stor_used_bytes     int64  `json:"stor_used_bytes"`
}

// @Description Msg_Resp_HeartBeat
type Msg_Resp_HeartBeat struct {
	Server_unixtime int64 `json:"server_unixtime"`
	api.API_META_STATUS
}

type Msg_Req_HeartBeatCallback struct {
	Token_md5 string `json:"token_md5"` //required
}

type Msg_Resp_HeartBeatCallback struct {
	api.API_META_STATUS
	Node_id             string `json:"node_id"`
	Downloading_threads uint64 `json:"downloading_threads"`
	Port                int    `json:"port"`
	Version             string `json:"version"`
	Stor_total_bytes    int64  `json:"stor_total_bytes"`
	Stor_used_bytes     int64  `json:"stor_used_bytes"`
	info.HardwareInfo
}
