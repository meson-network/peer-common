package node_domain

import (
	"github.com/meson-network/peer_common/api"
)

// @Description Msg_Resp_NodeDomain
type Msg_Resp_NodeDomain struct {
	api.API_META_STATUS
	Node_domain string `json:"node_domain"`
}
