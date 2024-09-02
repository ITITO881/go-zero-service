package types

// RefundQueryReq 退换单查询请求
type (
	RefundQueryReq struct {
		RefundNo      string  `form:"refund_no" mapstructure:"refund_no"`
		SrcRefundNo   string  `form:"src_refund_no" mapstructure:"src_refund_no"`
		TradeNo       string  `form:"trade_no" mapstructure:"trade_no"`
		Tid           string  `form:"tid" mapstructure:"tid"`
		LogisticsNo   string  `form:"logistics_no" mapstructure:"logistics_no"`
		ShopNos       string  `form:"shop_nos" mapstructure:"shop_nos"`
		ShopNo        string  `form:"shop_no" mapstructure:"shop_no"`
		ProcessStatus string  `form:"process_status" mapstructure:"process_status"`
		TimeType      string  `form:"time_type" mapstructure:"time_type"`
		StartTime     *string `form:"start_time" mapstructure:"start_time"`
		EndTime       *string `form:"end_time" mapstructure:"end_time"`
		PageNo        string  `form:"page_no" mapstructure:"page_no"`
		PageSize      string  `form:"page_size" mapstructure:"page_size"`
	}
	RefundQueryResp struct {
		TotalCount int          `json:"total_count"`
		Refunds    []RefundData `json:"refunds"`
	}
	RefundData struct {
		RefundId             int                  `json:"refund_id" mapstructure:"refund_id"`                           //退换单主键
		RefundNo             string               `json:"refund_no" mapstructure:"refund_no"`                           //erp退换单号，默认TK开头
		ApiOuterNo           string               `json:"api_outer_no" mapstructure:"api_outer_no"`                     //接口创建退换单的外部单号
		PlatformId           int                  `json:"platform_id" mapstructure:"platform_id"`                       //平台类型，响应值为代表平台的ID数字，ID对应的平台名称单击这里
		SrcNo                string               `json:"src_no" mapstructure:"src_no"`                                 //平台退款单号
		Type                 int                  `json:"type" mapstructure:"type"`                                     //1退款(未发货，退款申请)2退货3换货4退款不退货5订单补款
		ProcessStatus        int                  `json:"process_status" mapstructure:"process_status"`                 //5 补款 10已取消 20待审核 30已同意 40已拒绝 50待财审 60待收货 63待推送 64推送失败 65委外待收货 69待收货(已结算) 70部分到货 71部分到货(已结算) 80待结算 90已完成
		Status               int                  `json:"status" mapstructure:"status"`                                 //平台状态 0表示为空，1取消退款，2已申请退款，3等待退货，4等待收货，5退款成功，10换货待处理，20待买家退货，30买家已退货 待收货，40待买家修改，50待发出换货商品，55待买家收货，60换货改退货，70换货关闭，80换货成功
		GuaranteeMode        int                  `json:"guarantee_mode" mapstructure:"guarantee_mode"`                 //担保方式1担保交易2非担保交易3非担保在线交易（ecshop）
		CsStatus             int                  `json:"cs_status" mapstructure:"cs_status"`                           //退款阶段0售中1售后
		ShopNo               string               `json:"shop_no" mapstructure:"shop_no"`                               //代表店铺所有属性的唯一编码，用于店铺区分，ERP内支持自定义（ERP店铺界面设置）
		SwapTradeId          int                  `json:"swap_trade_id" mapstructure:"swap_trade_id"`                   //换出原始单id
		PayAccount           string               `json:"pay_account" mapstructure:"pay_account"`                       //买家支付帐号
		PayNo                string               `json:"pay_no" mapstructure:"pay_no"`                                 //支付订单号
		GoodsAmount          float64              `json:"goods_amount" mapstructure:"goods_amount"`                     //退货金额
		RefundAmount         float64              `json:"refund_amount" mapstructure:"refund_amount"`                   //买家申请退款金额
		DirectRefundAmount   float64              `json:"direct_refund_amount" mapstructure:"direct_refund_amount"`     //订单确认收货之后的退款金额
		GuaranteRefundAmount float64              `json:"guarante_refund_amount" mapstructure:"guarante_refund_amount"` //平台退款金额
		ActualRefundAmount   float64              `json:"actual_refund_amount" mapstructure:"actual_refund_amount"`     //实际退款/收款金额
		ExchangeAmount       float64              `json:"exchange_amount" mapstructure:"exchange_amount"`               //换货金额
		PostAmount           float64              `json:"post_amount" mapstructure:"post_amount"`                       //邮费金额
		OtherAmount          float64              `json:"other_amount" mapstructure:"other_amount"`                     //其它金额
		Paid                 float64              `json:"paid" mapstructure:"paid"`                                     //已退款/已收款
		Tid                  string               `json:"tid" mapstructure:"tid"`                                       //订单的原始单号
		SalesTradeId         int                  `json:"sales_trade_id" mapstructure:"sales_trade_id"`                 //系统订单ID
		SalesTid             string               `json:"sales_tid" mapstructure:"sales_tid"`                           //系统订单编号，默认单号为JY开头，ERP内支持自定义（设置路径：设置——编码设置）
		BuyerNick            string               `json:"buyer_nick" mapstructure:"buyer_nick"`                         //昵称
		ReceiverName         string               `json:"receiver_name" mapstructure:"receiver_name"`                   //收件人姓名
		ReceiverAddress      string               `json:"receiver_address" mapstructure:"receiver_address"`             //收件人地址
		ReceiverTelno        string               `json:"receiver_telno" mapstructure:"receiver_telno"`                 //收件人电话
		ReturnName           string               `json:"return_name" mapstructure:"return_name"`                       //退货人姓名
		ReturnMobile         string               `json:"return_mobile" mapstructure:"return_mobile"`                   //退货人手机号
		ReturnTelno          string               `json:"return_telno" mapstructure:"return_telno"`                     //退货人电话
		ReturnAddress        string               `json:"return_address" mapstructure:"return_address"`                 //退回仓库地址
		ReturnLogisticsName  string               `json:"return_logistics_name" mapstructure:"return_logistics_name"`   //退回物流公司名称
		ReturnLogisticsNo    string               `json:"return_logistics_no" mapstructure:"return_logistics_no"`       //退回物流单号
		SwapReceiver         string               `json:"swap_receiver" mapstructure:"swap_receiver"`                   //换货收件人姓名
		SwapMobile           string               `json:"swap_mobile" mapstructure:"swap_mobile"`                       //换收发件人姓名手机
		SwapTelno            string               `json:"swap_telno" mapstructure:"swap_telno"`                         //换货收件人姓名电话
		SwapProvince         int                  `json:"swap_province" mapstructure:"swap_province"`                   //换货收件人省份
		SwapCity             int                  `json:"swap_city" mapstructure:"swap_city"`                           //换货收件人城市
		SwapDistrict         int                  `json:"swap_district" mapstructure:"swap_district"`                   //换货收件人区县
		SwapArea             string               `json:"swap_area" mapstructure:"swap_area"`                           //换货收件人省市区
		SwapWarehouseNo      string               `json:"swap_warehouse_no" mapstructure:"swap_warehouse_no"`           //换出货品出库仓库编号
		SwapAddress          string               `json:"swap_address" mapstructure:"swap_address"`                     //换货收件人姓名地址
		WarehouseNo          string               `json:"warehouse_no" mapstructure:"warehouse_no"`                     //期望退回的仓库编号
		RefundTime           string               `json:"refund_time" mapstructure:"refund_time"`                       //退款创建时间格式：yyyy-MM-ddHH:mm:ss
		Remark               string               `json:"remark" mapstructure:"remark"`                                 //备注
		RefundReason         string               `json:"refund_reason" mapstructure:"refund_reason"`                   //退款的具体原因说明
		SwapTradeNo          string               `json:"swap_trade_no" mapstructure:"swap_trade_no"`                   //ERP生成的换出订单的原始订单编号
		Modified             string               `json:"modified" mapstructure:"modified"`                             //最后修改时间格式：yyyy-MM-ddHH:mm:ss
		FinishTime           string               `json:"finish_time" mapstructure:"finish_time"`                       //旺店通内操作结算的时间格式：yyyy-MM-ddHH:mm:ss
		CustomerNo           string               `json:"customer_no" mapstructure:"customer_no"`                       //客户档案编号
		CustomerName         string               `json:"customer_name" mapstructure:"customer_name"`                   //收件人姓名（如需唯一值，以receiver_name为准）
		SwapZip              string               `json:"swap_zip" mapstructure:"swap_zip"`                             //收件人邮编
		SwapLogisticsType    int                  `json:"swap_logistics_type" mapstructure:"swap_logistics_type"`       //响应值为代表换货新订单物流公司类别的数字，数字对应的物流公司类别名称单击这里
		SwapLogisticsName    string               `json:"swap_logistics_name" mapstructure:"swap_logistics_name"`       //换货新订单物流公司名称
		CreatorName          string               `json:"creator_name" mapstructure:"creator_name"`                     //建单员姓名
		ErrorMsg             string               `json:"error_msg" mapstructure:"error_msg"`                           //逆向物流错误信息
		Created              string               `json:"created" mapstructure:"created"`                               //创建时间
		OuterNo              string               `json:"outer_no" mapstructure:"outer_no"`
		ShopName             string               `json:"shop_name" mapstructure:"shop_name"`                         //店铺名称
		WarehouseId          int                  `json:"warehouse_id" mapstructure:"warehouse_id"`                   //期望退回的仓库
		ShopId               int                  `json:"shop_id" mapstructure:"shop_id"`                             //店铺id
		AdvanceStatus        int                  `json:"advance_status" mapstructure:"advance_status"`               //0; 退款先行垫付申请 1; 退款先行垫付，垫付完成 2; 退款先行垫付，卖家拒绝收货 3; 退款先行垫付，垫付关闭 4; 退款先行垫付，垫付分账成功 5;
		IsGoodsReceived      int                  `json:"is_goods_received" mapstructure:"is_goods_received"`         //买家是否收到货
		CustomerId           int                  `json:"customer_id" mapstructure:"customer_id"`                     //客户id
		SwapWarehouseId      int                  `json:"swap_warehouse_id" mapstructure:"swap_warehouse_id"`         //换出货品出库仓库
		SwapLogisticsId      int                  `json:"swap_logistics_id" mapstructure:"swap_logistics_id"`         //换货新订单物流公司ID
		WarehouseType        int                  `json:"warehouse_type" mapstructure:"warehouse_type"`               //期望退回仓库类型（0不限 1普通仓库 2自动流传外部 3京东仓储 4科捷 5百世物流 6SKU360 7通天晓 8中联网仓 9顺丰仓储 10网仓2号 11奇门仓储 12旺店通仓储 13心怡仓储 14力威仓储15京东沧海16云集仓储17POS18虎符20外部链路型22抖店云仓126分销委外仓127其它）
		WmsStatus            int                  `json:"wms_status" mapstructure:"wms_status"`                       //外部状态,与wms对接使用 0初始化 1失败 2成功
		WmsResult            string               `json:"wms_result" mapstructure:"wms_result"`                       //外部结果,与wms对接使用
		WmsOuterNo           string               `json:"wms_outer_no" mapstructure:"wms_outer_no"`                   //对应外部仓库的单号
		FromType             int                  `json:"from_type" mapstructure:"from_type"`                         //建立方式 1api抓单 2手式建单 3外部导入 4天猫换货 5菜鸟回传自动创建
		ReturnGoodsCount     float64              `json:"return_goods_count" mapstructure:"return_goods_count"`       //退货数量
		OpConstraint         int                  `json:"op_constraint" mapstructure:"op_constraint"`                 //操作约束 0无约束 1不允许修改 2不允许取消 3不允许修改和取消
		RefundVersion        string               `json:"refund_version" mapstructure:"refund_version"`               //退款版本,大于0后表明已经从平台下载完成，退款单不可修改
		ConsignMode          int                  `json:"consign_mode" mapstructure:"consign_mode"`                   //发货方式：0默认 1审核退款单后立即生成换货订单 2退货完成后生成换货订单
		LogisticsName        string               `json:"logistics_name" mapstructure:"logistics_name"`               //物流公司名称
		LogisticsNo          string               `json:"logistics_no" mapstructure:"logistics_no"`                   //物流单号
		PushNo               string               `json:"push_no" mapstructure:"push_no"`                             //物流递交单号
		LogisticsStyle       int                  `json:"logistics_style" mapstructure:"logistics_style"`             //取件方式 :0客户取件，1顺丰取件，2拒收退回
		ReasonId             int                  `json:"reason_id" mapstructure:"reason_id"`                         //原因
		RevertReason         int                  `json:"revert_reason" mapstructure:"revert_reason"`                 //驳回原因
		BadReason            int                  `json:"bad_reason" mapstructure:"bad_reason"`                       //错误原因
		ReturnMask           int                  `json:"return_mask" mapstructure:"return_mask"`                     //标志位 4-原始单唯一 8-增加退换在途量 16-确认退款金额标识(为了保持兼容，0表示已确认，1表示未确认) 32-对应订单是否开了电子发票 64天猫换货单关闭需要取消委外单
		IsTradeCharged       int                  `json:"is_trade_charged" mapstructure:"is_trade_charged"`           //是否结算后生成订单
		StockinPreNo         string               `json:"stockin_pre_no" mapstructure:"stockin_pre_no"`               //预入库单
		RevertReasonInfo     string               `json:"revert_reason_info" mapstructure:"revert_reason_info"`       //驳回原因说明
		FenxiaoNickName      string               `json:"fenxiao_nick_name" mapstructure:"fenxiao_nick_name"`         //分销商昵称
		FenxiaoRefundNo      string               `json:"fenxiao_refund_no" mapstructure:"fenxiao_refund_no"`         //分销商退款单号
		FenxiaoTid           string               `json:"fenxiao_tid" mapstructure:"fenxiao_tid"`                     //分销原始单号
		FlagName             string               `json:"flag_name" mapstructure:"flag_name"`                         //标记名称
		RefundOrderList      []RefundOrderList    `json:"refund_order_list" mapstructure:"refund_order_list"`         //退款子订单信息, refund_order_list结构
		RefundOutGoodsList   []RefundOutGoodsList `json:"refund_out_goods_list" mapstructure:"refund_out_goods_list"` //退货商品信息, refund_out_goods_list结构
	}
	RefundOrderList struct {
		OrderId           int     `json:"order_id" mapstructure:"order_id"`                       //退换单退回货品明细主键
		RefundId          int     `json:"refund_id" mapstructure:"refund_id"`                     //退换单主键
		Oid               string  `json:"oid" mapstructure:"oid"`                                 //原始子订单编号
		Tid               string  `json:"tid" mapstructure:"tid"`                                 //订单子订单的原始单号
		SrcNo             string  `json:"src_no" mapstructure:"src_no"`                           //平台退款单号
		ProcessStatus     int     `json:"process_status" mapstructure:"process_status"`           //处理状态0已取消 10待递交 20待审核 30已同意 40已拒绝 50待财审 60待收货 70部分到货 80已完成(业务意义不完整，请使用refunds节点下的process_status)
		SalesTid          string  `json:"sales_tid" mapstructure:"sales_tid"`                     //销售订单编号
		OrderNum          float64 `json:"order_num" mapstructure:"order_num"`                     //实际数量
		Price             float64 `json:"price" mapstructure:"price"`                             //价格
		PlatformId        int     `json:"platform_id" mapstructure:"platform_id"`                 //平台类型，响应值为代表平台的ID数字，ID对应的平台名称单击这里
		CostPrice         float64 `json:"cost_price" mapstructure:"cost_price"`                   //成本价
		OriginalPrice     float64 `json:"original_price" mapstructure:"original_price"`           //原价
		Discount          float64 `json:"discount" mapstructure:"discount"`                       //优惠
		Paid              float64 `json:"paid" mapstructure:"paid"`                               //已付金额
		MarketPrice       float64 `json:"market_price" mapstructure:"market_price"`               //单品市场价
		RefundNum         float64 `json:"refund_num" mapstructure:"refund_num"`                   //退款数量
		RefundOrderAmount float64 `json:"refund_order_amount" mapstructure:"refund_order_amount"` //明细退款金额，由主单退款金额分摊
		TotalAmount       float64 `json:"total_amount" mapstructure:"total_amount"`               //退货金额
		SpecId            int     `json:"spec_id" mapstructure:"spec_id"`                         //erp商品主键
		SpecNo            string  `json:"spec_no" mapstructure:"spec_no"`                         //代表单品(sku)所有属性的唯一编码，用于系统单品区分，SKU概念介绍,单击这里
		GoodsName         string  `json:"goods_name" mapstructure:"goods_name"`                   //货品名称
		GoodsNo           string  `json:"goods_no" mapstructure:"goods_no"`                       //代表货品(spu)所有属性的唯一编号，用于系统货品区分，，SPU概念介绍,单击这里
		SpecName          string  `json:"spec_name" mapstructure:"spec_name"`                     //规格名称
		SpecCode          string  `json:"spec_code" mapstructure:"spec_code"`                     //规格编码
		SuiteNo           string  `json:"suite_no" mapstructure:"suite_no"`                       //组合装编号
		SuiteName         string  `json:"suite_name" mapstructure:"suite_name"`                   //组合装名称
		SuiteNum          float64 `json:"suite_num" mapstructure:"suite_num"`                     //组合装数量
		StockinNum        float64 `json:"stockin_num" mapstructure:"stockin_num"`                 //退货入库数量
		Remark            string  `json:"remark" mapstructure:"remark"`                           //退换明细的备注
		IsSnEnable        int     `json:"is_sn_enable" mapstructure:"is_sn_enable"`               //默认为'0'，0不启用序列号 1强序列号 2弱序列号
		PositionNo        string  `json:"position_no" mapstructure:"position_no"`                 //货位编号
		PositionId        int     `json:"position_id" mapstructure:"position_id"`                 //货位ID，入库时默认ID
		SalesTradeId      int     `json:"sales_trade_id" mapstructure:"sales_trade_id"`           //系统订单ID
		SalesOrderId      int     `json:"sales_order_id" mapstructure:"sales_order_id"`           //默认为'0'
		GoodsId           int     `json:"goods_id" mapstructure:"goods_id"`                       //货品ID
		SuiteId           int     `json:"suite_id" mapstructure:"suite_id"`                       //默认为'0'，组合装ID
		StockinAmount     float64 `json:"stockin_amount" mapstructure:"stockin_amount"`           //默认为'0.0000'，退货入库总额
		Modified          string  `json:"modified" mapstructure:"modified"`                       //最后修改时间格式：yyyy-MM-ddHH:mm:ss
		Created           string  `json:"created" mapstructure:"created"`                         //创建时间
		EditMask          int     `json:"edit_mask" mapstructure:"edit_mask"`                     //默认为'0'，编辑掩码
		Tag               int     `json:"tag" mapstructure:"tag"`                                 //默认为'0'，标签
		Barcode           string  `json:"barcode" mapstructure:"barcode"`                         //主条码
		BrandName         string  `json:"brand_name" mapstructure:"brand_name"`                   //品牌名称
	}
	RefundOutGoodsList struct {
		RecId       int     `json:"rec_id" mapstructure:"rec_id"`             //主键id
		RefundId    int     `json:"refund_id" mapstructure:"refund_id"`       //退换单主键id
		TargetType  int     `json:"target_type" mapstructure:"target_type"`   //商品类型 1表示单品2表示组合装
		TargetId    int     `json:"target_id" mapstructure:"target_id"`       //组合装或者单品的id
		GoodsName   string  `json:"goods_name" mapstructure:"goods_name"`     //商品名称
		SpecName    string  `json:"spec_name" mapstructure:"spec_name"`       //规格名称
		MerchantNo  string  `json:"merchant_no" mapstructure:"merchant_no"`   //代表单品(sku)所有属性的唯一编码，用于系统单品区分，SKU概念介绍,单击这里
		RetailPrice float64 `json:"retail_price" mapstructure:"retail_price"` //零售价
		Num         float64 `json:"num" mapstructure:"num"`                   //数量
		Remark      string  `json:"remark" mapstructure:"remark"`             //备注
		SpecCode    string  `json:"spec_code" mapstructure:"spec_code"`       //规格编码
		IsSnEnable  int     `json:"is_sn_enable" mapstructure:"is_sn_enable"` //默认为'0'，0不启用序列号 1强序列号 2弱序列号
		Oid         string  `json:"oid" mapstructure:"oid"`                   //原始子订单编号
		Created     string  `json:"created" mapstructure:"created"`           //创建时间
		Modified    string  `json:"modified" mapstructure:"modified"`         //最后修改时间
		Barcode     string  `json:"barcode" mapstructure:"barcode"`           //主条码
	}
)

// VipApiRefundQueryReq 查询原始退款单请求
type (
	VipApiRefundQueryReq struct {
		LogisticsNo string  `form:"logistics_no" mapstructure:"logistics_no"`
		RefundNo    string  `form:"refund_no" mapstructure:"refund_no"`     // 平台退款单号
		Tid         string  `form:"tid" mapstructure:"tid"`                 // 平台订单号
		PlatformId  int     `form:"platform_id" mapstructure:"platform_id"` // 平台id（必须与tid或refund_no同时使用）
		ShopNo      string  `form:"shop_no" mapstructure:"shop_no"`         // 店铺编号（暂不支持批量查询）
		StartTime   *string `form:"start_time" mapstructure:"start_time"`   // 按时间增量获取数据,start_time作为开始时间，格式：yyyy-MM-ddHH:mm:ss
		EndTime     *string `form:"end_time" mapstructure:"end_time"`       // 按时间增量获取数据,end_time作为开始时间，格式：yyyy-MM-ddHH:mm:ss
		PageNo      string  `form:"page_no" mapstructure:"page_no"`         // 不传值默认从0页开始
		PageSize    string  `form:"page_size" mapstructure:"page_size"`     // 不传值默认从0页开始每页返回的数据条数，输入值范围1~100，不传本参数，输入值默认为40
		TimeType    string  `form:"time_type" mapstructure:"time_type"`     // 0:最后修改时间，1:退款时间。默认为0
	}
	VipApiRefundQueryResp struct {
		TotalCount int                 `json:"total_count" mapstructure:"total_count"` // 总记录数
		TradeList  []VipApiRefundTrade `json:"trade_list" mapstructure:"trade_list"`   // 退款单列表
	}
	VipApiRefundTrade struct {
		PlatformId          int                  `json:"platform_id" mapstructure:"platform_id"`                     // 平台id
		ShopNo              string               `json:"shop_no" mapstructure:"shop_no"`                             // 店铺编号
		RefundNo            string               `json:"refund_no" mapstructure:"refund_no"`                         // 平台退款编号
		Tid                 string               `json:"tid" mapstructure:"tid"`                                     // 平台订单号
		Type                int                  `json:"type" mapstructure:"type"`                                   // 类型：0.取消订单 1.退款（未发货退款） 2.退货 3.换货 4.退款不退货
		Status              int                  `json:"status" mapstructure:"status"`                               // 平台状态 0表示为空，1取消退款，2已申请退款，3等待退货，4等待收货，5退款成功，10换货待处理，20待买家退货，30买家已退货 待收货，40待买家修改，50待发出换货商品，55待买家收货，60换货改退货，70换货关闭，80换货成功
		ProcessStatus       int                  `json:"process_status" mapstructure:"process_status"`               // 系统状态：0.待递交 10.已取消 15.未递交 20.已递交
		OpConstraint        int                  `json:"op_constraint" mapstructure:"op_constraint"`                 // AG操作状态：0.未回传 3.回传AG成功 4.回传AG失败
		IsAftersale         int                  `json:"is_aftersale" mapstructure:"is_aftersale"`                   // 是否是售后退款单据
		CsStatus            int                  `json:"cs_status" mapstructure:"cs_status"`                         // 客服介入：1.需要客服介入 2.客服已经介入 3.客服初审完成 4.客服主管复审失败 5.客服处理完成
		GuaranteeMode       int                  `json:"guarantee_mode" mapstructure:"guarantee_mode"`               // 交易类型：1.担保交易 2.非担保交易 3.非担保在线交易
		RefundAmount        float64              `json:"refund_amount" mapstructure:"refund_amount"`                 // 申请退款金额
		ActualRefundAmount  float64              `json:"actual_refund_amount" mapstructure:"actual_refund_amount"`   // 实际退款金额
		Reason              string               `json:"reason" mapstructure:"reason"`                               // 退款原因
		Remark              string               `json:"remark" mapstructure:"remark"`                               // 备注
		RefundMask          int                  `json:"refund_mask" mapstructure:"refund_mask"`                     // 退款标记掩码 1024代表在途退仅限抖店
		LogisticsName       string               `json:"logistics_name" mapstructure:"logistics_name"`               // 物流公司名称
		LogisticsNo         string               `json:"logistics_no" mapstructure:"logistics_no"`                   // 物流单号
		RefundTime          string               `json:"refund_time" mapstructure:"refund_time"`                     // 退款时间格式：yyyy-MM-ddHH:mm:ss
		Modified            string               `json:"modified" mapstructure:"modified"`                           // 修改时间格式：yyyy-MM-ddHH:mm:ss
		Created             string               `json:"created" mapstructure:"created"`                             // 创建时间格式：yyyy-MM-ddHH:mm:ss
		CurrentPhaseTimeout string               `json:"current_phase_timeout" mapstructure:"current_phase_timeout"` // 退款成功时间格式：yyyy-MM-ddHH:mm:ss
		ShopId              int                  `json:"shop_id" mapstructure:"shop_id"`                             // 默认为'0'，店铺
		Title               string               `json:"title" mapstructure:"title"`                                 // 货品名称
		IsExternal          int                  `json:"is_external" mapstructure:"is_external"`                     // 默认为'0'，外部退款单，未经ERP处理过的订单
		RefundId            int                  `json:"refund_id" mapstructure:"refund_id"`                         // 自动增量
		PayNo               string               `json:"pay_no" mapstructure:"pay_no"`                               // 默认为''，支付订单号
		GoodsList           []RefundGoodsList    `json:"goods_list" mapstructure:"goods_list"`                       // 货品节点
		DiscountList        []RefundDiscountList `json:"discount_list" mapstructure:"discount_list"`                 // 优惠节点
	}
	RefundGoodsList struct {
		PlatformId  int     `json:"platform_id" mapstructure:"platform_id"`   // 平台id
		RefundNo    string  `json:"refund_no" mapstructure:"refund_no"`       // 平台退款编号
		Oid         string  `json:"oid" mapstructure:"oid"`                   // 子订单编号
		Status      int     `json:"status" mapstructure:"status"`             // 平台状态 0表示为空，1取消退款，2已申请退款，3等待退货，4等待收货，5退款成功，10换货待处理，20待买家退货，30买家已退货 待收货，40待买家修改，50待发出换货商品，55待买家收货，60换货改退货，70换货关闭，80换货成功
		GoodsId     int     `json:"goods_id" mapstructure:"goods_id"`         // 货品ID
		SpecId      int     `json:"spec_id" mapstructure:"spec_id"`           // 规格ID
		GoodsNo     string  `json:"goods_no" mapstructure:"goods_no"`         // 货品编码
		SpecNo      string  `json:"spec_no" mapstructure:"spec_no"`           // 规格编码
		Num         float64 `json:"num" mapstructure:"num"`                   // 数量
		Price       float64 `json:"price" mapstructure:"price"`               // 单价
		TotalAmount float64 `json:"total_amount" mapstructure:"total_amount"` // 总金额
		Remark      string  `json:"remark" mapstructure:"remark"`             // 备注
		Modified    string  `json:"modified" mapstructure:"modified"`         // 修改时间格式：yyyy-MM-ddHH:mm:ss
		Created     string  `json:"created" mapstructure:"created"`           // 创建时间格式：yyyy-MM-ddHH:mm:ss
		GoodsName   string  `json:"goods_name" mapstructure:"goods_name"`     // 默认为''，货品名称
		SpecName    string  `json:"spec_name" mapstructure:"spec_name"`       // 规格名
		OrderID     int     `json:"order_id" mapstructure:"order_id"`         // 默认为'0'，单据id
		ReturnTime  string  `json:"return_time" mapstructure:"return_time"`   // 退货时间

	}
	RefundDiscountList struct {
		RecId      int     `json:"rec_id" mapstructure:"rec_id"`           // 明细主键
		PlatformId int     `json:"platform_id" mapstructure:"platform_id"` // 平台id
		Tid        string  `json:"tid" mapstructure:"tid"`                 // 原始单号
		Oid        string  `json:"oid" mapstructure:"oid"`                 // 原始子订单编号
		RefundNo   string  `json:"refund_no" mapstructure:"refund_no"`     // 原始退款单号
		Name       string  `json:"name" mapstructure:"name"`               // 优惠名称
		Amount     float64 `json:"amount" mapstructure:"amount"`           // 金额
	}
)
