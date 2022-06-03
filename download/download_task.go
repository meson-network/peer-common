package download

const DOWNLOAD_THREADS_MAX = 10 //max number of parellel downloading threads

// @Description Msg_Req_Download_Task
type Msg_Req_Download_Task struct {
	Origin_url string `json:"origin_url"` //required
	File_hash  string `json:"file_hash"`  //required
}

//// @Description Msg_Resp_HeartBeat
//type Msg_Resp_Download_Task struct {
//	api.API_META_STATUS
//}
