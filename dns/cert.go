package dns

import "github.com/meson-network/peer_common/api"

// @Description Msg_Resp_HeartBeat
type Msg_Resp_Cert struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
	api.API_META_STATUS
}
