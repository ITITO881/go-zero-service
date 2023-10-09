package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "success"

	/**(API 文件上传 upload服务 相关错误码)**/
	message[XORM_FILE_DB_CREATE_FAIL] = "xorm文件登记数据库新增错误"
	message[XORM_SPU_FIELD_IS_REQUIRED] = "缺少必填的SPU字段"
	message[SOURCE_FIELD_REQUIRED] = "Source字段必传"
	message[SOURCE_VALUE_IS_NOT_VALID] = "Source字段值不存在"
	message[JSCODE_FIELD_REQUIRED] = "JsCode字段必传"
	message[XORM_FILE_LIST_QUERY_FAIL] = "xorm文件登记数据库列表查询错误"
	message[XORM_ENGINE_OPERATE_FAIL] = "XORM数据库读写错误"

	/**(客户认证 错误码)**/
	message[AUTHORIZATION_FIELD_REQUIRED] = "Authorization字段必传"

	// 微信注册相关错误
	message[WECHAT_JS_CODE_FAIL] = "微信小程序js_code验证失败"
	message[USER_NO_EXIST] = "用户未注册"
	message[GET_TOKEN_FAIL] = "用户获取TOKEN失败"
	message[WECHAT_PHONE_CODE_FAIL] = "微信小程序手机授权code验证失败"
	message[TOKEN_PARSE_FAIL] = "Token解析值错误"

	// 客户数据库相关错误
	message[CTM_BASE_USER_EXIST] = "微信会员已存在，不能二次注册"
	message[CTM_BASE_USER_CREATE_FAIL] = "微信会员数据库，创建失败"
	message[CTM_BASE_USER_RETRIEVE_FAIL] = "微信会员数据库，会员信息查询失败"
	message[CTM_BASE_USER_INFO_UPDATE_FAIL] = "微信会员数据库，会员信息更新失败"

	//	minio客户端相关错误
	message[FILEURL_FIELD_IS_NOT_VALID] = "fileUrl字段不合法"
	message[SOURCE_CODE_IS_NOT_VALID] = "sourceCode不合法"
	message[MINIO_CLIENT_OPERATE_FAIL] = "minio数据库读写错误"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func NewCodeMsg(errCode uint32) *CodeError {
	return &CodeError{
		code: errCode,
		msg:  MapErrMsg(errCode),
	}
}
