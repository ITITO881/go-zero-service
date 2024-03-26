package conf

type JwtAuthConf struct {
	AppKey       int32
	AccessSecret string
	AccessExpire int64
}

type MinIOClientConf struct {
	Endpoint          string
	AccessKeyID       string
	SecretAccessKey   string
	UseSSL            bool
	BucketName        string
	UploadContentType string
	ExpireDuration    int64
	Location          string
}

type RedisConf struct {
	Addr     string
	Password string
	DB       int
}

type TokenParam struct {
	UserId     int32  `json:"userId"`
	Username   string `json:"username,omitempty"`
	OpenId     string `json:"openId,omitempty"`
	UserMobile string `json:"userMobile,omitempty"`
}

//type MinIOClaimConf struct {
//	Endpoint          string
//	AccessKeyID       string
//	SecretAccessKey   string
//	UseSSL            bool
//	BucketName        string
//	UploadContentType string
//	ExpireDuration    int64
//	Location          string
//}
