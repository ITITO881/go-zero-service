package xerr

// OK 成功返回
const OK uint32 = 0

/**(前3位代表业务,后三位代表具体功能)**/

const (
	/**(API 文件上传 upload服务 相关错误码)**/
	XORM_FILE_DB_CREATE_FAIL   uint32 = 135001
	XORM_SPU_FIELD_IS_REQUIRED uint32 = 135002
	SOURCE_VALUE_IS_NOT_VALID  uint32 = 135003
	JSCODE_FIELD_REQUIRED      uint32 = 135004
	SOURCE_FIELD_REQUIRED      uint32 = 135006
	XORM_FILE_LIST_QUERY_FAIL  uint32 = 135007
	XORM_ENGINE_OPERATE_FAIL   uint32 = 135009
	XORM_QUERY_EMPTY           uint32 = 135011
)

const (
	/**(RPC Go 客户认证 错误码)**/
	AUTHORIZATION_FIELD_REQUIRED uint32 = 120001 // Authorization字段必传
	WECHAT_JS_CODE_FAIL          uint32 = 120006 // 微信小程序js_code验证失败
	WECHAT_PHONE_CODE_FAIL       uint32 = 120009 // 微信小程序手机授权code验证失败
	GET_TOKEN_FAIL               uint32 = 120008 // 用户获取TOKEN失败
	TOKEN_PARSE_FAIL             uint32 = 120010 // Token解析值错误
	USER_NO_EXIST                uint32 = 120007 // 用户未注册
)

const (
	/**(RPC Python 客户数据库)**/
	CTM_BASE_USER_EXIST            uint32 = 130001
	CTM_BASE_USER_CREATE_FAIL      uint32 = 130002
	CTM_BASE_USER_RETRIEVE_FAIL    uint32 = 130003
	CTM_BASE_USER_INFO_UPDATE_FAIL uint32 = 130004
)

const (
	FILEURL_FIELD_IS_NOT_VALID uint32 = 131021
	SOURCE_CODE_IS_NOT_VALID   uint32 = 131022
	MINIO_CLIENT_OPERATE_FAIL  uint32 = 131023
	MINIO_OBJECT_GET_FAIL      uint32 = 131024
)

const (
	ORM_ERROR uint32 = 200000
)
