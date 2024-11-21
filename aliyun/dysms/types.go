package dysms

//	type AliyunSmsConf struct {
//		AccessKeyId     string
//		AccessKeySecret string
//		Endpoint        string
//	}
type DysmsResponse struct {
	Code int32
	Msg  string
}

type TemplateParamReq struct {
	Title   string `json:"title,omitempty"`
	Result  string `json:"result,omitempty"`
	Scheme  string `json:"scheme,omitempty"`
	Express string `json:"express,omitempty"`
}
