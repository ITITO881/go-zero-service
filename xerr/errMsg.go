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
	message[XORM_QUERY_EMPTY] = "查询结果为 空"

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
	message[MINIO_OBJECT_GET_FAIL] = "minio数据库获取对象数据失败"

	message[REDIS_ERROR] = "redis error"

	message[COUPON_STOCK_NOT_ENOUGH] = "优惠券库存不足"
	message[ACTIVITY_NOT_AVALIABLE] = "活动已失效"
	message[ACTIVITY_NOT_STARTED] = "活动尚未开启"
	message[COUPON_TAKE_LIMIT] = "优惠券领取数量已达上限"

	// 订单 配置表 相关错误
	message[ORDER_CONF_TRADE_TYPE_FAIL] = "订单类型获取错误"
	message[ORDER_CONF_STOCK_LIST_FAIL] = "旺店通仓库列表获取错误"
	message[ORDER_CONF_SHOP_LIST_FAIL] = "旺店通店铺列表获取错误"
	message[ORDER_CONF_UPDATE_FAIL] = "订单设置更新错误"
	message[ORDER_CONF_GET_LIST_FAIL] = "订单类型列表查询错误"

	// 订单配置 相关 拒绝原因 相关错误
	message[REFUND_REASON_TRADE_TYPE_FAIL] = "指定订单类型下 拒绝原因列表查询错误"
	message[REFUND_REASON_CREATE_FAIL] = "新增 拒绝原因失败"
	message[REFUND_REASON_UPDATE_FAIL] = "更新 拒绝原因失败"
	message[REFUND_REASON_DESTROY_FAIL] = "删除 拒绝原因失败"
	message[REFUND_REASON_CHILDREN_BATCH_UPDATE_FAIL] = "根据父类是否启用，批量更新子类是否启用失败"

	// 订单主表 相关错误
	message[ORDER_MODEL_OPERATE_FAIL] = "订单主表 操作错误"
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
