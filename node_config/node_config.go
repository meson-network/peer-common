package node_config

import (
	"github.com/meson-network/peer_common/api"
)

// @Description Msg_Resp_NodeConfig
type Msg_Resp_NodeConfig struct {
	api.API_META_STATUS
	Cached_file_expire_secs int64   `json:"cached_file_expire_secs"`
	Max_file_size_bytes     int64   `json:"max_file_size_bytes"`
	Free_space_line         int64   `json:"free_space_line"`
	Delete_trigger_rate     float64 `json:"delete_trigger_rate"`
}
