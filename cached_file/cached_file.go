package cached_file

// @Description Msg_Req_FileExpire
type Msg_Req_FileExpire struct {
	Expired_files []string `json:"expired_files"` //required
}

//// @Description Msg_Resp_HeartBeat
//type Msg_Resp_FileExpire struct {
//	api.API_META_STATUS
//}

// @Description Msg_Req_FileExpire
type Msg_Req_FileMissing struct {
	Missing_files []string `json:"missing_files"` //required
}

//// @Description Msg_Resp_HeartBeat
//type Msg_Resp_FileMissing struct {
//	api.API_META_STATUS
//}
