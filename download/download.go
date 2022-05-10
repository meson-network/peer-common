package download

import "github.com/meson-network/peer_common/api"

const DOWNLOAD_THREADS_MAX = 10 //max number of parellel downloading threads

// @Description Msg_Resp_Download
type Msg_Resp_Download struct {
	Origin_url string `json:"origin_url"`
	Url_hash   string `json:"url_hash"`
	api.API_META_STATUS
}
