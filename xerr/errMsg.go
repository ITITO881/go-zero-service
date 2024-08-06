package xerr

var message map[uint32]string
var payState map[string]string

func init() {
	message = make(map[uint32]string)
	payState = make(map[string]string)
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
	message[WECHAT_SECURITY_FAIL] = "文案中包含不良内容，请修改后重新提交。"

	// 客户数据库相关错误
	message[CTM_BASE_USER_EXIST] = "微信会员已存在，不能二次注册"
	message[CTM_BASE_USER_CREATE_FAIL] = "微信会员数据库，创建失败"
	message[CTM_BASE_USER_RETRIEVE_FAIL] = "微信会员数据库，会员信息查询失败"
	message[CTM_BASE_USER_INFO_UPDATE_FAIL] = "微信会员数据库，会员信息更新失败"
	message[CTM_BASE_USER_GET_LIST_FAIL] = "微信用户列表 获取失败"

	//	minio客户端相关错误
	message[FILEURL_FIELD_IS_NOT_VALID] = "fileUrl字段不合法"
	message[SOURCE_CODE_IS_NOT_VALID] = "sourceCode不合法"
	message[MINIO_CLIENT_OPERATE_FAIL] = "minio数据库读写错误"
	message[MINIO_OBJECT_GET_FAIL] = "minio数据库获取对象数据失败"

	message[REDIS_ERROR] = "redis error"
	message[DS_API_ERROR] = "数仓API请求出错"
	message[JAVA_API_ERROR] = "JAVA API返回错误码"

	message[COUPON_STOCK_NOT_ENOUGH] = "优惠券库存不足"
	message[ACTIVITY_NOT_AVALIABLE] = "活动已失效"
	message[ACTIVITY_NOT_STARTED] = "活动尚未开启"
	message[COUPON_TAKE_LIMIT] = "优惠券领取数量已达上限"
	message[ACTIVITY_NEED_ATTEND] = "请先参与征集活动"

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

	// 用户 记名 优惠券表 相关错误
	message[CUSTOMER_COUPON_OPERATE_FAIL] = "用户记名优惠券 操作错误"
	message[CUSTOMER_COUPON_REDUCE_AMOUNT_MORE_THAN_RAW_PRICE] = "错误：折扣价 高于 原价"
	message[CUSTOMER_COUPON_DISCOUNT_AMOUNT_PARSE_FAIL] = "优惠券 折扣金额 解析错误"
	message[CUSTOMER_COUPON_DISCOUNT_PARSE_FAIL] = "优惠券 折扣 解析错误"

	message[CTM_COUPON_CRATE_ITEM_FAIL] = "用户记名优惠券 新增错误"

	message[SPU_COUPON_INFORMATION_ERROR] = "单品级 优惠券信息错误"
	message[ORDER_COUPON_INFORMATION_ERROR] = "订单级 优惠券信息错误"

	message[ORDER_COUPON_OPERATE_FAIL] = "用户不记名通用优惠券 操作错误"
	message[ORDER_COUPON_SPLIT_BY_DETAIL_FAIL] = "订单优惠 按明细分摊 操作错误"

	// 不记名 优惠券管理 相关错误
	message[CARD_COUPON_OPTIONS_GET_FAIL] = "优惠券管理 预定义项 查询错误"
	message[CARD_COUPON_RPC_GET_LIST_FAIL] = "优惠券RPC 列表查询失败"
	message[CARD_COUPON_RPC_CREATE_ITEM_FAIL] = "优惠券RPC 新增失败"
	message[CARD_COUPON_RPC_UPDATE_ITEM_FAIL] = "优惠券RPC 部分更新失败"
	message[CARD_COUPON_RPC_RETRIEVE_ITEM_FAIL] = "优惠券RPC 单查失败"
	message[CARD_COUPON_RPC_DESTROY_ITEM_FAIL] = "优惠券RPC 单删失败"

	// 优惠活动管理 相关错误
	message[ACT_DISCOUNT_OPTIONS_GET_FAIL] = "优惠活动管理 预定义项 查询错误"
	message[ACT_DISCOUNT_RPC_GET_LIST_FAIL] = "优惠活动RPC 列表查询失败"
	message[ACT_DISCOUNT_RPC_CREATE_ITEM_FAIL] = "优惠活动RPC 新增失败"
	message[ACT_DISCOUNT_RPC_UPDATE_ITEM_FAIL] = "优惠活动RPC 部分更新失败"
	message[ACT_DISCOUNT_RPC_RETRIEVE_ITEM_FAIL] = "优惠活动RPC 单查失败"
	message[ACT_DISCOUNT_RPC_DESTROY_ITEM_FAIL] = "优惠活动RPC 单删失败"

	// 优惠活动下 优惠券 相关错误
	message[VOUCHER_IN_ACT_CREATE_ITEM_FAIL] = "优惠活动下 优惠券RPC 新增失败"
	message[VOUCHER_IN_ACT_GET_LIST_FAIL] = "优惠活动下 优惠券RPC 列表查询失败"
	message[VOUCHER_IN_ACT_UPDATE_ITEM_FAIL] = "优惠活动下 优惠券RPC 部分更新失败"
	message[VOUCHER_IN_ACT_RETRIEVE_ITEM_FAIL] = "优惠活动下 优惠券RPC 单查失败"
	message[VOUCHER_IN_ACT_DESTROY_ITEM_FAIL] = "优惠活动下 优惠券RPC 单删失败"

	// 优惠活动下 优惠券 发卡记录 相关错误
	message[RECORD_IN_ACT_CREATE_ITEM_FAIL] = "优惠活动下 优惠券发卡记录RPC 新增失败"
	message[RECORD_IN_ACT_GET_LIST_FAIL] = "优惠活动下 优惠券发卡记录RPC 列表查询失败"
	message[RECORD_IN_ACT_UPDATE_ITEM_FAIL] = "优惠活动下 优惠券发卡记录RPC 部分更新失败"
	message[RECORD_IN_ACT_RETRIEVE_ITEM_FAIL] = "优惠活动下 优惠券发卡记录RPC 单查失败"
	message[RECORD_IN_ACT_DESTROY_ITEM_FAIL] = "优惠活动下 优惠券发卡记录RPC 单删失败"

	// 订单主表 相关错误
	message[ORDER_MODEL_OPERATE_FAIL] = "订单主表 操作错误"
	message[ORDER_CONF_MODEL_OPERATION_FAIL] = "订单主表之配置子表 操作错误"
	message[CTM_BASE_MODEL_OPERATION_FAIL] = "订单基础表 操作错误"

	// 用户关联产品 相关错误
	message[PRODUCT_NOT_EXITS] = "产品不存在"
	message[PRODUCT_ALREADY_BIND] = "产品已被绑定"
	message[PRODUCT_NOT_MATCH] = "产品信息不匹配"
	message[PRODUCT_NOT_SUPPORT_AFTERSALE] = "当前产品不支持售后"
	message[PRODUCT_ON_AFTERSALE] = "当前产品维修/售后中"
	message[PRODUCT_IS_INVALID] = "产品已下架，请重新选择"

	// 服务单 相关错误
	message[SERVICE_RPC_FAIL] = "服务单RPC 失败"
	message[TRADE_TYPE_EMPTY_FAIL] = "服务单 业务类型字段 缺少"
	message[TYPE_EMPTY_FAIL] = "退换单 申请子类型TYPE 缺少"
	message[OBJECT_REPAIR_RPC_FAIL] = "维修单RPC 失败"
	message[OBJECT_SWAP_RPC_FAIL] = "换货单RPC 失败"
	message[OBJECT_TAILOR_RPC_FAIL] = "定制单RPC 失败"
	message[OBJECT_COMPLAIN_RPC_FAIL] = "投诉单RPC 失败"
	message[OBJECT_FEEDBACK_RPC_FAIL] = "反馈单RPC 失败"
	message[OBJECT_TRADE_IN_RPC_FAIL] = "以旧换新RPC 失败"
	message[SERVICE_GOODS_RETURN_RPC_FAIL] = "退货单RPC 失败"
	message[SERVICE_GOODS_DELIVERY_RPC_FAIL] = "发货单RPC 失败"
	message[SERVICE_GOODS_PAY_RPC_FAIL] = "支付单RPC 失败"

	message[RPC_PARSE_TO_API_FAIL] = "RPC解析转换API 失败"
	message[PRODUCT_RAW_QUERY_FAIL] = "产品原始信息 查询失败"

	message[WDT_TRADE_PUSH_FAIL] = "旺店通 下推订单推送失败"
	message[WDT_TRADE_QUERY_FAIL] = "旺店通 查询订单失败"
	message[WDT_QM_TRADE_QUERY_FAIL] = "旺店通奇门 查询订单失败"
	// 定制相关错误
	message[WDT_TRADE_NOT_EXIST] = "查询订单不存在"
	message[WDT_TRADE_STATUS_ERROR_CANNOT_TAILOR] = "订单状态错误，已出库不能定制"
	message[WDT_TAILOR_PERMISSION_SKU_REQUIRE] = "缺少定制许可SKU"
	message[WDT_NO_CAN_TAILOR_PRODUCT] = "无可定制产品"
	message[WDT_ALREADY_TAILORED] = "已定制|定制次数不足"

	message[WDT_STOCK_IN_PUSH_FAIL] = "旺店通 下推入库单推送失败"
	message[EXP_CLIENT_QUERY_FAIL] = "快递单号 客户端 查询失败"

	payState[WX_PAY_DETAIL_NOT_PAY] = "待支付"

	// 品牌活动相关
	message[BRAND_ACTIVITY_NOT_EXIST] = "活动不存在"
	message[BRAND_ACTIVITY_NOT_START] = "活动尚未开始，敬请期待"
	message[BRAND_ACTIVITY_NOT_AVAILABLE] = "活动已停止报名"
	message[BRAND_ACTIVITY_EXPIRED] = "活动已结束"
	message[BRAND_BACTIVITY_ACCEPTED] = "您已领取该活动卡券，，请前往权益卡列表查看"
	message[BRAND_ACTIVITY_FINISHED] = "手慢了，卡券已被领取完"

	// 产品相关
	message[PRODUCT_STOCK_NOT_ENOUGH] = "库存不足"
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
