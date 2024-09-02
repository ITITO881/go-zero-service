package types

type (
	TradeQueryReq struct {
		Status         *int    `form:"status" mapstructure:"status"`                     // 订单状态 5已取消 10待付款 12待尾款 13待选仓 15等未付16延时审核 19预订单前处理 20前处理(赠品，合并，拆分)21委外前处理22抢单前处理 25预订单 27待抢单 30待客审 35待财审 40待递交仓库 45递交仓库中 50已递交仓库 53未确认 55已确认(已审核) 95已发货 105部分打款 110已完成 113异常发货
		StartTime      *string `form:"start_time" mapstructure:"start_time"`             // 按最后修改时间增量获取数据,start_time作为开始时间，时间格式：yyyy-MM-ddHH:mm:ss
		EndTime        *string `form:"end_time" mapstructure:"end_time"`                 // 按最后修改时间增量获取数据,end_time作为结束时间，时间格式：yyyy-MM-ddHH:mm:ss
		PageSize       *int    `form:"page_size" mapstructure:"page_size"`               // 每页返回的数据条数，输入值范围1~100，不传本参数，输入值默认为40，使用举例单击这里
		PageNo         *int    `form:"page_no" mapstructure:"page_no"`                   // 不传值默认从0页开始
		SrcTid         string  `form:"src_tid" mapstructure:"src_tid"`                   // 如果使用原始单号，其余参数不起效。其余参数可以不传
		TradeNo        string  `form:"trade_no" mapstructure:"trade_no"`                 // 系统订单编号，默认单号为JY开头，ERP内支持自定义（设置路径：设置——编码设置）如果使用订单编号，其余参数不起效，其余参数可以不传
		ShopNo         string  `form:"shop_no" mapstructure:"shop_no"`                   // 代表店铺所有属性的唯一编码，用于店铺区分，ERP内支持自定义（ERP店铺界面设置），用于获取指定店铺单据数据信息
		WarehouseNo    string  `form:"warehouse_no" mapstructure:"warehouse_no"`         // 代表仓库所有属性的唯一编码，用于仓库区分，ERP内支持自定义（ERP仓库界面设置），用于获取指定仓库单据数据信息(不支持一次推送多个仓库编码)
		Goodstax       *int    `form:"goodstax" mapstructure:"goodstax"`                 // 0 使用订单中的税率 1 使用单品中的税率(默认0)
		HasLogisticsNo *int    `form:"has_logistics_no" mapstructure:"has_logistics_no"` // 0 没有任何限制(默认值) 1 物流单号不为空才返回 2 只返回物流单号为空的
		IsFuzzy        *int    `form:"is_fuzzy" mapstructure:"is_fuzzy"`                 // 0精确1模糊默认为0 （仅在原始单号src_tid查询时生效）
		Src            *int    `form:"src" mapstructure:"src"`                           // 1 （返回 ）,0 (不返回)（不传默认值0）
		LogisticsNo    string  `form:"logistics_no" mapstructure:"logistics_no"`         // 如果使用物流单号，其余参数不起效。其余参数可以不传
		ShopNos        string  `form:"shop_nos" mapstructure:"shop_nos"`                 // 批量指定店铺获取（多个店铺编号之间用英文逗号隔开，最多指定20个店铺）
	}
	TradeQueryResp struct {
		TotalCount int          `json:"total_count" mapstructure:"total_count"` // 总记录数
		Trades     []TradeOrder `json:"trades" mapstructure:"trades"`           // 订单列表
	}
	TradeOrder struct {
		TradeId             *int        `json:"trade_id" mapstructure:"trade_id"`                             // erp订单表的主键
		TradeNo             string      `json:"trade_no" mapstructure:"trade_no"`                             // 系统订单编号，默认单号为JY开头，ERP内支持自定义订单编号生成规则（设置路径：设置——编码设置）
		PlatformId          *int        `json:"platform_id" mapstructure:"platform_id"`                       // 响应值为代表平台的ID数字，ID对应的平台名称单击这里
		ShopPlatformId      *int        `json:"shop_platform_id" mapstructure:"shop_platform_id"`             // 响应值为代表平台的ID数字，ID对应的平台名称单击这里
		ShopNo              string      `json:"shop_no" mapstructure:"shop_no"`                               // 代表店铺所有属性的唯一编码，用于店铺区分，ERP内支持自定义（ERP店铺界面设置）
		ShopName            string      `json:"shop_name" mapstructure:"shop_name"`                           // 店铺名称
		ShopRemark          string      `json:"shop_remark" mapstructure:"shop_remark"`                       // 店铺备注
		WarehouseType       *int        `json:"warehouse_type" mapstructure:"warehouse_type"`                 // （0不限 1普通仓库 2自动流传外部 3京东仓储 4科捷 5百世物流 6SKU360 7通天晓 8中联网仓 9顺丰仓储 10网仓2号 11奇门仓储 12旺店通仓储 13心怡仓储 14力威仓储15京东沧海16云集仓储17POS18虎符20外部链路型22抖店云仓126分销委外仓127其它）
		WarehouseNo         string      `json:"warehouse_no" mapstructure:"warehouse_no"`                     // 代表仓库所有属性的唯一编码，用于仓库区分，ERP内支持自定义（ERP仓库界面设置）（根据编号可查询仓库名称）
		SrcTids             string      `json:"src_tids" mapstructure:"src_tids"`                             // 商城或电商平台的订单编号,合并订单的多个订单编号有逗号隔开
		TradeStatus         *int        `json:"trade_status" mapstructure:"trade_status"`                     // 订单状态 5已取消 10待付款 12待尾款 13待选仓 15等未付16延时审核 19预订单前处理 20前处理(赠品，合并，拆分)21委外前处理22抢单前处理 25预订单 27待抢单 30待客审 35待财审 40待递交仓库 45递交仓库中 50已递交仓库 53未确认 55已确认（已审核） 95已发货 105部分打款 110已完成 113异常发货
		ConsignStatus       *int        `json:"consign_status" mapstructure:"consign_status"`                 // 出库状态：0表示无出库状态，1验货 2称重 4出库 8物流同步16分拣32档口 64拣货 128供销回传成功 256供销回传失败（注意：如果是3，则表示完成了验货和称重“1+2”，如果是15，则表示四个过程都完成了“1+2+4+8”，其他数字以此类推）
		TradeType           *int        `json:"trade_type" mapstructure:"trade_type"`                         // 1网店销售 2线下零售 3售后换货 4批发业务 5保修换新 6保修完成 7订单补发8供销补发 101自定义类型1 102自定义类型2 103自定义类型3 104自定义类型4 105自定义类型5 106自定义类型6 107自定义属性7 108自定义属性8
		DeliveryTerm        *int        `json:"delivery_term" mapstructure:"delivery_term"`                   // 发货条件 1款到发货 2货到付款(包含部分货到付款)，4挂账
		FreezeReason        *int        `json:"freeze_reason" mapstructure:"freeze_reason"`                   // 冻结原因
		RefundStatus        *int        `json:"refund_status" mapstructure:"refund_status"`                   // 退款状态 0无退款 1申请退款 2部分退款 3全部退款 4未付款关闭或手工关闭
		FenxiaoType         *int        `json:"fenxiao_type" mapstructure:"fenxiao_type"`                     // 0非分销订单 1转供销 2代销 3经销
		FenxiaoNick         string      `json:"fenxiao_nick" mapstructure:"fenxiao_nick"`                     // 分销商信息，不同平台单据分销商会存不同格式情况，具体以返回信息为准
		TradeTime           *string     `json:"trade_time" mapstructure:"trade_time"`                         // 下单时间时间格式：yyyy-MM-ddHH:mm:ss
		PayTime             *string     `json:"pay_time" mapstructure:"pay_time"`                             // 付款时间时间格式：yyyy-MM-ddHH:mm:ss
		CustomerName        string      `json:"customer_name" mapstructure:"customer_name"`                   // 客户名称
		CustomerNo          string      `json:"customer_no" mapstructure:"customer_no"`                       // 客户编码
		PayAccount          string      `json:"pay_account" mapstructure:"pay_account"`                       // 买家付款账号
		BuyerNick           string      `json:"buyer_nick" mapstructure:"buyer_nick"`                         // 客户网名
		ReceiverName        string      `json:"receiver_name" mapstructure:"receiver_name"`                   // 收件人
		ReceiverProvince    *int        `json:"receiver_province" mapstructure:"receiver_province"`           // 收件人的省份,按照城市代码表中对应城市代码进行返回
		ReceiverCity        *int        `json:"receiver_city" mapstructure:"receiver_city"`                   // 收件人的城市,按照城市代码表中对应城市代码进行返回
		ReceiverDistrict    *int        `json:"receiver_district" mapstructure:"receiver_district"`           // 收件人的地区,按照城市代码表中对应城市代码进行返回
		ReceiverAddress     string      `json:"receiver_address" mapstructure:"receiver_address"`             // 地址
		ReceiverMobile      string      `json:"receiver_mobile" mapstructure:"receiver_mobile"`               // 手机
		ReceiverTelno       string      `json:"receiver_telno" mapstructure:"receiver_telno"`                 // 电话
		ReceiverZip         string      `json:"receiver_zip" mapstructure:"receiver_zip"`                     // 收件人邮编
		ReceiverArea        string      `json:"receiver_area" mapstructure:"receiver_area"`                   // 省市县空格分隔
		ReceiverRing        string      `json:"receiver_ring" mapstructure:"receiver_ring"`                   // 收件人区域
		ReceiverDtb         string      `json:"receiver_dtb" mapstructure:"receiver_dtb"`                     // 大头笔
		ToDeliverTime       *string     `json:"to_deliver_time" mapstructure:"to_deliver_time"`               // 派送时间 如周一至周五，上午
		BadReason           *int        `json:"bad_reason" mapstructure:"bad_reason"`                         // 异常订单(bit位)， 2地址变化 8仓库变化 16客服备注变化 128拦截赠品 64退款 2048买家留言发生变更
		LogisticsId         *int        `json:"logistics_id" mapstructure:"logistics_id"`                     // 物流公司ID（系统自增主键id）
		LogisticsName       string      `json:"logistics_name" mapstructure:"logistics_name"`                 // 物流公司名称
		LogisticsCode       string      `json:"logistics_code" mapstructure:"logistics_code"`                 // 代表物流所有属性的唯一编码，用于物流区分，ERP内支持自定义（ERP物流界面设置）
		LogisticsType       *int        `json:"logistics_type" mapstructure:"logistics_type"`                 // 响应值为代表物流方式的数字，数字对应的物流方式名称单击这里
		LogisticsNo         string      `json:"logistics_no" mapstructure:"logistics_no"`                     // 物流单号
		ShopId              *int        `json:"shop_id" mapstructure:"shop_id"`                               // 店铺列表主键
		WarehouseId         *int        `json:"warehouse_id" mapstructure:"warehouse_id"`                     // 仓库ID
		CheckStep           *int        `json:"check_step" mapstructure:"check_step"`                         // 用于多级审核,特殊值：-100根据预售策略自动转入特殊单，-101人工转入
		UnmergeMask         *int        `json:"unmerge_mask" mapstructure:"unmerge_mask"`                     // 未合并标记,1有未付款订单,2有同名未合并订单
		DelayToTime         *string     `json:"delay_to_time" mapstructure:"delay_to_time"`                   // 延时此进一步处理，等未付或延时审核 激活时间 秒级时间戳格式返回
		CustomerType        *int        `json:"customer_type" mapstructure:"customer_type"`                   // 0普通客户1经销商
		CustomerId          string      `json:"customer_id" mapstructure:"customer_id"`                       // 买家ID
		ReceiverCountry     *int        `json:"receiver_country" mapstructure:"receiver_country"`             // 收件人的国家,按照城市代码表中对应城市代码进行返回
		PreChargeTime       string      `json:"pre_charge_time" mapstructure:"pre_charge_time"`               // 配送中心,未使用
		IsPrevNotify        *int        `json:"is_prev_notify" mapstructure:"is_prev_notify"`                 // 是否京配(为1时,只能发京邦达）
		NoteCount           *int        `json:"note_count" mapstructure:"note_count"`                         // 便签条数
		BuyerMessageCount   *int        `json:"buyer_message_count" mapstructure:"buyer_message_count"`       // 买家留言条数
		CsRemarkCount       *int        `json:"cs_remark_count" mapstructure:"cs_remark_count"`               // 客服备注条数
		CsRemarkChangeCount *int        `json:"cs_remark_change_count" mapstructure:"cs_remark_change_count"` // 客服备注变化 0,未变化1平台变化,2手工修改,4发票手工修改
		DiscountChange      *float64    `json:"discount_change" mapstructure:"discount_change"`               // 优惠变化金额,更新货品和数量
		TradePrepay         *float64    `json:"trade_prepay" mapstructure:"trade_prepay"`                     // 客户使用的预存款
		PiAmount            *float64    `json:"pi_amount" mapstructure:"pi_amount"`                           // 分期付款金额
		OtherCost           *float64    `json:"other_cost" mapstructure:"other_cost"`                         // 其它成本(目前用作记录货到付款单据的物流佣金）
		Volume              *float64    `json:"volume" mapstructure:"volume"`                                 // 体积（单位：cm³）
		SalesScore          string      `json:"sales_score" mapstructure:"sales_score"`                       // 销售积分,未使用
		FlagId              *int        `json:"flag_id" mapstructure:"flag_id"`                               // 背景色标记
		IsSealed            *int        `json:"is_sealed" mapstructure:"is_sealed"`                           // 不可合并拆分
		GiftMask            *int        `json:"gift_mask" mapstructure:"gift_mask"`                           // 赠品标记1已处理过赠品,但没有匹配任何策略2自动赠送4手工赠送6即有自动也有手工
		SplitFromTradeId    string      `json:"split_from_trade_id" mapstructure:"split_from_trade_id"`       // 拆分订单，原单ID，用于避免自动合并,大件拆分为（原订单的id值），自动拆分为负值（原订单的-id值）
		LogisticsTemplateId string      `json:"logistics_template_id" mapstructure:"logistics_template_id"`   // 未使用
		SendbillTemplateId  string      `json:"sendbill_template_id" mapstructure:"sendbill_template_id"`     // 未使用
		RevertReason        *int        `json:"revert_reason" mapstructure:"revert_reason"`                   // 驳回原因
		CancelReason        *int        `json:"cancel_reason" mapstructure:"cancel_reason"`                   // 取消原因
		IsUnpaymentSms      *int        `json:"is_unpayment_sms" mapstructure:"is_unpayment_sms"`             // 催未付款订单消息发送标记
		PackageId           string      `json:"package_id" mapstructure:"package_id"`                         // 包装id
		TradeMask           string      `json:"trade_mask" mapstructure:"trade_mask"`                         // 订单标记位 1使用智选物流 2 航空禁运 4 预订单自动转审核失败 8 预占用待发货库存 16 订单货品指定批次 32 自动流转仓库 64 部分发货 128 全部发货 256 已发过签收消息 512 大单锁定仓库 1024 人工转入预订单 2048因配置先占用待发货库存 4096 顺丰前置发货 8192 订单批量合并后标记16384 达到原始单最大合并限制32768 物流升级原始单指定物流订单 65536 BIC订单标记 131072 标记货品指定物流 3554432 前N有礼订单 67108864 预售下沉
		Reserve             string      `json:"reserve" mapstructure:"reserve"`                               // 保留
		LargeType           *int        `json:"large_type" mapstructure:"large_type"`                         // 包含大件类型，1普通套件2独立套件3分组单发,未使用-1非单发件 取子单中的最大值
		BuyerMessage        string      `json:"buyer_message" mapstructure:"buyer_message"`                   // 买家留言
		CsRemark            string      `json:"cs_remark" mapstructure:"cs_remark"`                           // 客服备注
		RemarkFlag          *int        `json:"remark_flag" mapstructure:"remark_flag"`                       // 标旗 0至5对应的标旗颜色依次为灰（无标旗）、红、黄、绿、蓝、紫
		PrintRemark         string      `json:"print_remark" mapstructure:"print_remark"`                     // 打印备注
		GoodsTypeCount      *int        `json:"goods_type_count" mapstructure:"goods_type_count"`             // 货品种类数
		GoodsCount          *float64    `json:"goods_count" mapstructure:"goods_count"`                       // 货品总数
		GoodsAmount         *float64    `json:"goods_amount" mapstructure:"goods_amount"`                     // 货品总额（未扣除优惠），sum（share_amount+discount）所得
		PostAmount          *float64    `json:"post_amount" mapstructure:"post_amount"`                       // 邮费
		OtherAmount         *float64    `json:"other_amount" mapstructure:"other_amount"`                     // 其它从买家的收费（非订单支付金额以及服务费），从原始订单列表继承
		Discount            *float64    `json:"discount" mapstructure:"discount"`                             // 订单优惠，系统子订单“优惠”求合所得
		Receivable          *float64    `json:"receivable" mapstructure:"receivable"`                         // 应收金额，系统订单的（“货品总额”+“邮资”-“折扣”）所得
		DapAmount           *float64    `json:"dap_amount" mapstructure:"dap_amount"`                         // 款到发货金额,paid>=dap_amount才可发货
		CodAmount           *float64    `json:"cod_amount" mapstructure:"cod_amount"`                         // 货到付款订单金额，系统子订单的（“分摊后总价”+“分摊邮费”-“已付”）再求和
		ExtCodFee           *float64    `json:"ext_cod_fee" mapstructure:"ext_cod_fee"`                       // 货到付款非订单金额，从原始订单继承
		GoodsCost           *float64    `json:"goods_cost" mapstructure:"goods_cost"`                         // 货款预估成本
		PostCost            *float64    `json:"post_cost" mapstructure:"post_cost"`                           // 预估邮费成本
		Paid                *float64    `json:"paid" mapstructure:"paid"`                                     // 已付金额，系统子订单“已付”求合所得
		Weight              *float64    `json:"weight" mapstructure:"weight"`                                 // 预估重量（单位：kg）
		Profit              *float64    `json:"profit" mapstructure:"profit"`                                 // 预估毛利
		Tax                 *float64    `json:"tax" mapstructure:"tax"`                                       // 税额
		TaxRate             *float64    `json:"tax_rate" mapstructure:"tax_rate"`                             // 税率
		Commission          *float64    `json:"commission" mapstructure:"commission"`                         // 佣金
		InvoiceType         *int        `json:"invoice_type" mapstructure:"invoice_type"`                     // 发票类别 0 不需要，1普通发票，2增值普通税发票，3增值专用税发票
		InvoiceTitle        string      `json:"invoice_title" mapstructure:"invoice_title"`                   // 发票抬头
		InvoiceContent      string      `json:"invoice_content" mapstructure:"invoice_content"`               // 发票内容
		SalesmanId          *int        `json:"salesman_id" mapstructure:"salesman_id"`                       // 业务员ID
		CheckerId           *int        `json:"checker_id" mapstructure:"checker_id"`                         // 审核员工ID
		Fullname            string      `json:"fullname" mapstructure:"fullname"`                             // 业务员姓名
		CheckerName         string      `json:"checker_name" mapstructure:"checker_name"`                     // 审核员工姓名
		FcheckerId          *int        `json:"fchecker_id" mapstructure:"fchecker_id"`                       // 财审操作员ID
		CheckouterId        *int        `json:"checkouter_id" mapstructure:"checkouter_id"`                   // 签出员工id
		StockoutNo          string      `json:"stockout_no" mapstructure:"stockout_no"`                       // 出库单号,内部或外部仓库的订单号
		FlagName            string      `json:"flag_name" mapstructure:"flag_name"`                           // 背景色标记名称
		TradeFrom           *int        `json:"trade_from" mapstructure:"trade_from"`                         // 订单来源 1API抓单，2手工建单 3excel导入 4现款销售
		SingleSpecNo        string      `json:"single_spec_no" mapstructure:"single_spec_no"`                 // 货品商家编码
		RawGoodsCount       *float64    `json:"raw_goods_count" mapstructure:"raw_goods_count"`               // 原始货品数量
		RawGoodsTypeCount   *int        `json:"raw_goods_type_count" mapstructure:"raw_goods_type_count"`     // 原始货品种类数
		Currency            string      `json:"currency" mapstructure:"currency"`                             // 币种
		SplitPackageNum     *int        `json:"split_package_num" mapstructure:"split_package_num"`           // 已拆分包裹数
		InvoiceId           *int        `json:"invoice_id" mapstructure:"invoice_id"`                         // 0表示未开发票，＞0表示已开发票
		VersionId           *int        `json:"version_id" mapstructure:"version_id"`                         // 订单每修改一次，版本号做一次变更
		Modified            string      `json:"modified" mapstructure:"modified"`                             // 最后修改时间 时间格式：yyyy-MM-ddHH:mm:ss
		Created             string      `json:"created" mapstructure:"created"`                               // 系统单生成时间 时间格式：yyyy-MM-ddHH:mm:ss
		IdCardType          *int        `json:"id_card_type" mapstructure:"id_card_type"`                     // 证件类别
		IdCard              string      `json:"id_card" mapstructure:"id_card"`                               // 淘系平台不返回，其他平台正常返回
		FcheckerName        string      `json:"fchecker_name" mapstructure:"fchecker_name"`                   // 财审人名称
		CheckouterName      string      `json:"checkouter_name" mapstructure:"checkouter_name"`               // 签出人名称
		FreezeReasonInfo    string      `json:"freeze_reason_info" mapstructure:"freeze_reason_info"`         // 冻结原因名称
		FenxiaoTid          string      `json:"fenxiao_tid" mapstructure:"fenxiao_tid"`                       // 分销原始单号
		Tags                string      `json:"tags" mapstructure:"tags"`                                     // 订单标签/异常id，具体标签/异常名称查询，需使用查询标签/异常名称接口
		GoodsList           []GoodsData `json:"goods_list" mapstructure:"goods_list"`                         // 响应参数的2级数据节点，包含系统订单货品明细所有属性信息的数据节点,节点下数据字段详见下述"goods_list"
	}
	GoodsData struct {
		RecId           *int     `json:"rec_id" mapstructure:"rec_id"`                       // erp子订单主键
		TradeId         *int     `json:"trade_id" mapstructure:"trade_id"`                   // erp订单主键
		SpecId          *int     `json:"spec_id" mapstructure:"spec_id"`                     // erp内商品主键
		PlatformId      *int     `json:"platform_id" mapstructure:"platform_id"`             // 平台ID
		SrcOid          string   `json:"src_oid" mapstructure:"src_oid"`                     // 原始子订单号
		PlatformGoodsId string   `json:"platform_goods_id" mapstructure:"platform_goods_id"` // 平台货品ID
		PlatformSpecId  string   `json:"platform_spec_id" mapstructure:"platform_spec_id"`   // 平台商品ID
		SuiteId         *int     `json:"suite_id" mapstructure:"suite_id"`                   // 如果货品是由组合装拆分的，这里组合装ID
		Flag            *int     `json:"flag" mapstructure:"flag"`                           // 子单拆分时，单品明细都一致，用来区分子单的唯一性
		SrcTid          string   `json:"src_tid" mapstructure:"src_tid"`                     // 原始订单号
		GiftType        *int     `json:"gift_type" mapstructure:"gift_type"`                 // 是否是赠品 0非赠品 1自动赠送 2手工赠送 3回购自动送赠品 4前N有礼送赠品 6天猫优仓赠品 7淘宝CRM会员送赠
		RefundStatus    *int     `json:"refund_status" mapstructure:"refund_status"`         // 退款状态 0无退款,1取消退款,2已申请退款,3等待退货,4等待收货,5退款成功 6已关闭
		GuaranteeMode   *int     `json:"guarantee_mode" mapstructure:"guarantee_mode"`       // 1担保 2非担保 3在线非担保
		DeliveryTerm    *int     `json:"delivery_term" mapstructure:"delivery_term"`         // 1款到发货 2货到付款(包含部分货到付款) 3分期付款
		BindOid         string   `json:"bind_oid" mapstructure:"bind_oid"`                   // 关联发货
		Num             *float64 `json:"num" mapstructure:"num"`                             // 货品数量
		Price           *float64 `json:"price" mapstructure:"price"`                         // 销售单价，手工新建时使用货品属性中的“零售价”
		ActualNum       *float64 `json:"actual_num" mapstructure:"actual_num"`               // 实发数量,此数量为发货数量,删除操作等于将此值设置为0
		RefundNum       *float64 `json:"refund_num" mapstructure:"refund_num"`               // 无实际含义
		OrderPrice      *float64 `json:"order_price" mapstructure:"order_price"`             // 成交价,原始单折扣及分摊之后的价格
		SharePrice      *float64 `json:"share_price" mapstructure:"share_price"`             // 进入ERP后再次调整的价格，默认值与order_price一致
		Adjust          *float64 `json:"adjust" mapstructure:"adjust"`                       // 手工调整价,正数为加价,负数为减价
		Discount        *float64 `json:"discount" mapstructure:"discount"`                   // 总折扣金额
		ShareAmount     *float64 `json:"share_amount" mapstructure:"share_amount"`           // 分摊后合计应收=share_price * num , share_price 是根据share_amount反推的,因此share_price可能有精度损失
		SharePost       *float64 `json:"share_post" mapstructure:"share_post"`               // 分摊邮费
		Paid            *float64 `json:"paid" mapstructure:"paid"`                           // 已支付金额
		GoodsName       string   `json:"goods_name" mapstructure:"goods_name"`               // 货品名称
		ClassName       string   `json:"class_name" mapstructure:"class_name"`               // 分类名称
		Prop2           string   `json:"prop2" mapstructure:"prop2"`                         // 自定义属性2
		TcOrderId       string   `json:"tc_order_id" mapstructure:"tc_order_id"`             // 分销商id
		Barcode         string   `json:"barcode" mapstructure:"barcode"`                     // 主条码
		GoodsId         *int     `json:"goods_id" mapstructure:"goods_id"`                   // 货品id
		GoodsNo         string   `json:"goods_no" mapstructure:"goods_no"`                   // 代表货品(spu)所有属性的唯一编号，用于系统货品区分，，SPU概念介绍单击这里
		SpecName        string   `json:"spec_name" mapstructure:"spec_name"`                 // 规格名
		SpecNo          string   `json:"spec_no" mapstructure:"spec_no"`                     // 代表单品(sku)所有属性的唯一编码，用于系统单品区分，SKU概念介绍,单击这里
		SpecCode        string   `json:"spec_code" mapstructure:"spec_code"`                 // 规格码
		SuiteNo         string   `json:"suite_no" mapstructure:"suite_no"`                   // 代表组合装商品所有属性的唯一编码，用于系统组合装商品的区分
		SuiteName       string   `json:"suite_name" mapstructure:"suite_name"`               // 如果是组合装拆分的，此为组合装名称
		SuiteNum        *float64 `json:"suite_num" mapstructure:"suite_num"`                 // 组合装数量
		SuiteAmount     *float64 `json:"suite_amount" mapstructure:"suite_amount"`           // 组合装分摊后总价
		SuiteDiscount   *float64 `json:"suite_discount" mapstructure:"suite_discount"`       // 组合装优惠
		ShareAmount2    *float64 `json:"share_amount2" mapstructure:"share_amount2"`         // share_amount备份值,退款恢复使用,!可回收 目前会存放分销订单发货回传的分销价
		IsPrintSuite    *int     `json:"is_print_suite" mapstructure:"is_print_suite"`       // 0:组合装明细/1:组合装以及明细/2:组合装
		IsZeroCost      *int     `json:"is_zero_cost" mapstructure:"is_zero_cost"`           // 是否允许0成本
		StockReserved   *int     `json:"stock_reserved" mapstructure:"stock_reserved"`       // 库存保留情况 0未保留（取消的订单或完成）1无库存记录 2未付款 3已保留待审核 4待发货 5预订单库存
		IsConsigned     *int     `json:"is_consigned" mapstructure:"is_consigned"`           // 平台已发货
		IsReceived      *int     `json:"is_received" mapstructure:"is_received"`             // 线上订单，标记是否打款
		IsMaster        *int     `json:"is_master" mapstructure:"is_master"`                 // 是否主子订单,为发货算法使用
		ApiGoodsName    string   `json:"api_goods_name" mapstructure:"api_goods_name"`       // 平台货品名称
		ApiSpecName     string   `json:"api_spec_name" mapstructure:"api_spec_name"`         // 平台规格名称
		Weight          *float64 `json:"weight" mapstructure:"weight"`                       // 预估单个货品重量（单位：kg）
		Commission      *float64 `json:"commission" mapstructure:"commission"`               // 佣金
		GoodsType       *int     `json:"goods_type" mapstructure:"goods_type"`               // 1销售商品 2原材料 3包装 4周转材料5虚拟商品 0其它
		LargeType       *int     `json:"large_type" mapstructure:"large_type"`               // 大件类型 0非大件 1普通大件 2独立大件
		InvoiceType     *int     `json:"invoice_type" mapstructure:"invoice_type"`           // 发票类别，0 不需要，1普通发票，2增值税发票
		InvoiceContent  string   `json:"invoice_content" mapstructure:"invoice_content"`     // 发票内容
		FromMask        *int     `json:"from_mask" mapstructure:"from_mask"`                 // 订单内部来源1手机，2聚划算，4服务子订单，8家装，16二次付款，32开具电子发票，128指定批次，2048当日达，4096次日达，8192预计时效，262144天猫直送/唯品仓中仓，524288‘3PL时效/jitx’，2097152区域零售，4194304预售单，8388608周期购，16777216同城购(淘宝)，268435456物流升级，536870912 抖店抽奖订单
		Cid             *int     `json:"cid" mapstructure:"cid"`                             // 类目id
		Remark          string   `json:"remark" mapstructure:"remark"`                       // 货品明细备注
		Modified        string   `json:"modified" mapstructure:"modified"`                   // 最后修改时间 时间格式：yyyy-MM-ddHH:mm:ss
		Created         string   `json:"created" mapstructure:"created"`                     // 创建时间 时间格式：yyyy-MM-ddHH:mm:ss
		TaxRate         *float64 `json:"tax_rate" mapstructure:"tax_rate"`                   // 税率（根据条件使用订单中或单品中税率）
		BaseUnitId      *int     `json:"base_unit_id" mapstructure:"base_unit_id"`           // 基本单位ID
		UnitName        string   `json:"unit_name" mapstructure:"unit_name"`                 // 基本单位名称
		PayId           string   `json:"pay_id" mapstructure:"pay_id"`                       // 交易流水单号  请求参数src传为1时返回
		PayStatus       *int     `json:"pay_status" mapstructure:"pay_status"`               // 0(未付款),1(部分付款),2(已付款)如需响应该字段，请将ERP升级到V2.3.8.3及以上
		PayTime         *string  `json:"pay_time" mapstructure:"pay_time"`                   // 时间格式：yyyy-MM-ddHH:mm:ss，如需响应该字段，请将ERP升级到V2.3.8.3及以上
	}
)
