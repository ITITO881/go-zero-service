package types

type PurchaseOrderQueryReq struct {
	PurchaseNo  string `form:"purchase_no,omitempty" mapstructure:"purchase_no"`   // ERP系统采购单编号，传该字段可以不传开始时间和结束时间
	OuterNo     string `form:"outer_no,omitempty" mapstructure:"outer_no"`         // ERP推送采购单外部返回单号（对应外部仓储单号） 传该字段可以不传开始时间和结束时间
	Status      int    `form:"status,omitempty" mapstructure:"status"`             // 状态 10 已取消,20 编辑中,30 待审核,35 待财审,40 已审核,43 待推送,45 推送失败,48 待收货,50 部分到货,60已到货,70 待结算,80 部分结算,90 已完成（不传默认查询全部状态采购单）
	StartTime   string `form:"start_time,omitempty" mapstructure:"start_time"`     // 按最后修改时间增量获取数据，start_time作为开始时间，格式：yyyy-MM-ddHH:mm:ss
	EndTime     string `form:"end_time,omitempty" mapstructure:"end_time"`         // 按最后修改时间增量获取数据，end_time作为结束时间，格式：yyyy-MM-ddHH:mm:ss
	WarehouseNo string `form:"warehouse_no,omitempty" mapstructure:"warehouse_no"` // 代表仓库所有属性的唯一编码，用于仓库区分，ERP内支持自定义（ERP仓库界面设置）
	PageSize    int    `form:"page_size,omitempty" mapstructure:"page_size"`       // 每页返回的数据条数，输入值范围1~100，不传本参数，输入值默认为40，使用举例单击这里
	PageNo      int    `form:"page_no,omitempty" mapstructure:"page_no"`           // 页码，从0开始
}

type PurchaseData struct {
	PurchaseId          int                  `json:"purchase_id" mapstructure:"purchase_id"`                     // 采购单id， 采购单表主键
	PurchaseNo          string               `json:"purchase_no" mapstructure:"purchase_no"`                     // ERP系统采购单编号
	PurchaseOuterNo     string               `json:"purchase_outer_no" mapstructure:"purchase_outer_no"`         // ERP推送采购单外部返回单号（对应外部仓储单号）。
	OuterNo             string               `json:"outer_no" mapstructure:"outer_no"`                           // 通过接口创建采购单时传的api单号
	WarehouseName       string               `json:"warehouse_name" mapstructure:"warehouse_name"`               // 仓库名称
	WarehouseNo         string               `json:"warehouse_no" mapstructure:"warehouse_no"`                   // 代表仓库所有属性的唯一编码，用于仓库区分，ERP内支持自定义（ERP仓库界面设置）
	Status              int                  `json:"status" mapstructure:"status"`                               // 采购单状态：10 已取消,20 编辑中,30 待审核,35 待财审,40 已审核,43 待推送,45 推送失败,48 待收货,50 部分到货,60已到货,70 待结算,80 部分结算,90 已完成
	ProviderNo          string               `json:"provider_no" mapstructure:"provider_no"`                     // 供应商编号
	ProviderName        string               `json:"provider_name" mapstructure:"provider_name"`                 // 供应商名称
	Contact             string               `json:"contact" mapstructure:"contact"`                             // 联系人
	Telno               string               `json:"telno" mapstructure:"telno"`                                 // 电话
	PurchaserName       string               `json:"purchaser_name" mapstructure:"purchaser_name"`               // 采购员
	ReceiveAddress      string               `json:"receive_address" mapstructure:"receive_address"`             // 收货地址
	LogisticsType       int                  `json:"logistics_type" mapstructure:"logistics_type"`               // 货运方式
	CheckTime           string               `json:"check_time" mapstructure:"check_time"`                       // 审核时间，格式：yyyy-MM-dd HH:mm:ss
	GoodsAmount         string               `json:"goods_amount" mapstructure:"goods_amount"`                   // 货款
	PostFee             string               `json:"post_fee" mapstructure:"post_fee"`                           // 邮资
	OtherFee            string               `json:"other_fee" mapstructure:"other_fee"`                         // 其他费用
	TaxGoodsAmount      string               `json:"tax_goods_amount" mapstructure:"tax_goods_amount"`           // 税后货品总金额
	Remark              string               `json:"remark" mapstructure:"remark"`                               // 采购单备注
	Modified            string               `json:"modified" mapstructure:"modified"`                           // 单据的最后修改时间，格式：yyyy-MM-dd HH:mm:ss
	ExpectArriveTime    string               `json:"expect_arrive_time" mapstructure:"expect_arrive_time"`       // 采购预计到货时间，格式：yyyy-MM-dd HH:mm:ss
	Prop1               string               `json:"prop1" mapstructure:"prop1"`                                 // 自定义属性1
	Prop2               string               `json:"prop2" mapstructure:"prop2"`                                 // 自定义属性2
	Created             string               `json:"created" mapstructure:"created"`                             // 单据创建时间，格式：yyyy-MM-dd HH:mm:ss
	CreateId            int                  `json:"creator_id" mapstructure:"creator_id"`                       // 默认为'0'，创建人id
	CreatorName         string               `json:"creator_name" mapstructure:"creator_name"`                   // 默认为'0'，创建人姓名
	WarehouseId         int                  `json:"warehouse_id" mapstructure:"warehouse_id"`                   // 仓库id
	ApiOuterNo          string               `json:"api_outer_no" mapstructure:"api_outer_no"`                   // 默认为''，接口创建的单据的外部单号，同outer_no
	WmsOuterNo          string               `json:"wms_outer_no" mapstructure:"wms_outer_no"`                   // 默认为''，对应外部仓库的单号，同purchase_outer_no字段
	WmsStatus           int                  `json:"wms_status" mapstructure:"wms_status"`                       // 默认为'0'，外部状态,与wms对接使用 0初始化 1失败 2成功
	ErrorInfo           string               `json:"error_info" mapstructure:"error_info"`                       // 默认为''，返回错误信息
	PurchaserId         int                  `json:"purchaser_id" mapstructure:"purchaser_id"`                   // 默认为'0'，采购员id
	OrderType           int                  `json:"order_type" mapstructure:"order_type"`                       // 默认为'0'，0普通采购单，1档口采购单,2APP档口采购单
	FlagId              int                  `json:"flag_id" mapstructure:"flag_id"`                             // 默认为'0'，颜色标记
	ProviderId          int                  `json:"provider_id" mapstructure:"provider_id"`                     // 默认为'0'，供应商
	CheckOperatorId     int                  `json:"check_operator_id" mapstructure:"check_operator_id"`         // 默认为'0'，审核人
	FacheckOperatorId   int                  `json:"facheck_operator_id" mapstructure:"facheck_operator_id"`     // 默认为'0'，财审人
	FecheckTime         string               `json:"facheck_time" mapstructure:"facheck_time"`                   // 财审时间，格式：yyyy-MM-dd HH:mm:ss
	GoodsFee            string               `json:"goods_fee" mapstructure:"goods_fee"`                         // 货品总价
	GoodsCount          string               `json:"goods_count" mapstructure:"goods_count"`                     // 货品数量
	TaxFee              string               `json:"tax_fee" mapstructure:"tax_fee"`                             // 税后总价
	GoodsTypeCount      int                  `json:"goods_type_count" mapstructure:"goods_type_count"`           // 货品种类
	RevertReason        int                  `json:"revert_reason" mapstructure:"revert_reason"`                 // 默认为'0'，驳回原因
	VersionId           int                  `json:"version_id" mapstructure:"version_id"`                       // 默认为'0'，版本号
	NoteCount           int                  `json:"note_count" mapstructure:"note_count"`                       // 默认为'0'，便签条数
	PropTime1           string               `json:"prop_time1" mapstructure:"prop_time1"`                       // 默认为''，自定义时间1
	CheckOperatorName   string               `json:"check_operator_name" mapstructure:"check_operator_name"`     // 审核员名称
	FacheckOperatorName string               `json:"facheck_operator_name" mapstructure:"facheck_operator_name"` // 财审员名称
	DetailsList         []PurchaseDetailItem `json:"details_list" mapstructure:"details_list"`
}

type PurchaseDetailItem struct {
	RecId           int    `json:"rec_id" mapstructure:"rec_id"`                       // 采购明细表主键id
	PurchaseId      int    `json:"purchase_id" mapstructure:"purchase_id"`             // 采购单id
	SpecNo          string `json:"spec_no" mapstructure:"spec_no"`                     // 代表单品（sku）所有属性的编码
	BrandName       string `json:"brand_name" mapstructure:"brand_name"`               // 品牌名称，推送此字段前需将版本升级至V2.3.9.1及以上的稳定版本
	Num             string `json:"num" mapstructure:"num"`                             // 数量
	Price           string `json:"price" mapstructure:"price"`                         // 税前单价
	Discount        string `json:"discount" mapstructure:"discount"`                   // 折扣
	Tax             string `json:"tax" mapstructure:"tax"`                             // 税率
	TaxPrice        string `json:"tax_price" mapstructure:"tax_price"`                 // 税后单价
	Amount          string `json:"amount" mapstructure:"amount"`                       // 税前金额
	TaxAmount       string `json:"tax_amount" mapstructure:"tax_amount"`               // 税后金额
	StockinAmount   string `json:"stockin_amount" mapstructure:"stockin_amount"`       // 采购到货入库总金额
	StockinNum      string `json:"stockin_num" mapstructure:"stockin_num"`             // 采购到货入库数量
	Prop1           string `json:"prop1" mapstructure:"prop1"`                         // 自定义属性1
	Prop2           string `json:"prop2" mapstructure:"prop2"`                         // 自定义属性2
	Remark          string `json:"remark" mapstructure:"remark"`                       // 明细备注
	Modified        string `json:"modified" mapstructure:"modified"`                   // 修改时间，格式：yyyy-MM-dd HH:mm:ss
	ProviderGoodsNo string `json:"provider_goods_no" mapstructure:"provider_goods_no"` // 供应商货品编号
	GoodsNo         string `json:"goods_no" mapstructure:"goods_no"`                   // 货品编号
	GoodsName       string `json:"goods_name" mapstructure:"goods_name"`               // 货品名称
	SpecName        string `json:"spec_name" mapstructure:"spec_name"`                 // 货品规格名称
	SpecCode        string `json:"spec_code" mapstructure:"spec_code"`                 // 规格码
	BaseUnit        string `json:"base_unit" mapstructure:"base_unit"`                 // 基本单位
	PurchaseUnit    string `json:"purchase_unit" mapstructure:"purchase_unit"`         // 采购单位
	ArriveNum       string `json:"arrive_num" mapstructure:"arrive_num"`               // 采购到货未入库数量
	StopwaitNum     string `json:"stopwait_num" mapstructure:"stopwait_num"`           // 停止等待数量
	PurchasePrice   string `json:"purchase_price" mapstructure:"purchase_price"`       // 采购价
	LastPrice       string `json:"last_price" mapstructure:"last_price"`               // 上一次采购价
	LastSecondPrice string `json:"last_second_price" mapstructure:"last_second_price"` // 上一次采购价
	LowestPrice     string `json:"lowest_price" mapstructure:"lowest_price"`           // 最低采购价
	Weight          string `json:"weight" mapstructure:"weight"`                       // 重量
	ArriveMoreNum   string `json:"arrive_more_num" mapstructure:"arrive_more_num"`     // 多到货量
	lackNum         string `json:"lack_num" mapstructure:"lack_num"`                   // 未到货量
	SpecId          int    `json:"spec_id" mapstructure:"spec_id"`                     // 规格id
	WarehouseId     int    `json:"warehouse_id" mapstructure:"warehouse_id"`           // 默认为'0'，仓库id
	Tag             int    `json:"tag" mapstructure:"tag"`                             // 默认为'0'，tag=0表示默认采购
	Num2            string `json:"num2" mapstructure:"num2"`                           // 默认为'0.0000'，辅助单位数量
	UnitRatio       string `json:"unit_ratio" mapstructure:"unit_ratio"`               // 默认为'1.0000'，基本单位和辅助单位的换算系数
	UnitId          int    `json:"unit_id" mapstructure:"unit_id"`                     // 辅助单位/采购单位
	BaseUnitId      int    `json:"base_unit_id" mapstructure:"base_unit_id"`           // 基本单位
	Created         string `json:"created" mapstructure:"created"`                     // 创建时间
}

type PurchaseOrderQueryResp struct {
	TotalCount   int            `json:"total_count" mapstructure:"total_count"`     // 符合条件的数据总条数，当page_no=0时返回，用来分页
	PurchaseList []PurchaseData `json:"purchase_list" mapstructure:"purchase_list"` // 采购单列表
}
