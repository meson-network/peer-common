package speed_tester

import (
	"github.com/meson-network/peer_common/api"
)

type SpeedTestTarget struct {
	Ip    string `json:"ip"`
	Token string `json:"token"`
	Port  string `json:"port"`
}

// @Description Msg_Resp_SpeedTestJob
type Msg_Resp_SpeedTestJob struct {
	api.API_META_STATUS
	Targets []*SpeedTestTarget `json:"targets"`
}

// @Description Msg_Req_Report
type Msg_Req_Report struct {
	Ip                  string `json:"ip"`                  //required
	Success             bool   `json:"success"`             //required
	Bandwidth_bytes_sec int64  `json:"bandwidth_bytes_sec"` //required
}

//// @Description Msg_Resp_Report
//type Msg_Resp_Report struct {
//	api.API_META_STATUS
//}
