package version

import (
	"github.com/meson-network/peer_common/api"
)

// @Description Msg_Resp_NodeVersion
type Msg_Resp_NodeVersion struct {
	api.API_META_STATUS
	Allow_version  string `json:"allow_version"`
	Latest_version string `json:"latest_version"`
	Download_host  string `json:"download_host"`
}
