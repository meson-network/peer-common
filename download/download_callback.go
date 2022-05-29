package download

import (
	"github.com/meson-network/peer_common/api"
)

// @Description Msg_Req_Download_Callback_Success
type Msg_Req_Download_Callback_Success struct {
	//Origin_url string `json:"origin_url"` //required
	File_hash string `json:"file_hash"` //required
	File_size int64  `json:"file_size"` //required
}

// @Description Msg_Req_Download_Callback_Failed
type Msg_Req_Download_Callback_Failed struct {
	//Origin_url string `json:"origin_url"` //required
	File_hash string `json:"file_hash"` //required
}

type Msg_Resp_Download_Callback struct {
	api.API_META_STATUS
}
