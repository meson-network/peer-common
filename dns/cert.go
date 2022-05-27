package dns

import "github.com/meson-network/peer_common/api"

// @Description Msg_Resp_HeartBeat
type Msg_Resp_Cert struct {
	Crt string `json:"crt"`
	Key string `json:"key"`
	api.API_META_STATUS
}

// @Description Msg_Resp_HeartBeat
type Msg_Resp_CertHash struct {
	Hash string `json:"hash"`
	api.API_META_STATUS
}
