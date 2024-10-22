package types

type (
	WarehouseQueryReq struct {
		WarehouseNo string `form:"warehouse_no" mapstructure:"warehouse_no"` // 仓库编码
		Type        *int   `form:"type" mapstructure:"type"`                 // 仓库类型 0不限 1普通仓库 2自动流传外部 3京东仓储 4科捷 5百世物流 6SKU360 7通天晓 8中联网仓 9顺丰仓储 10网仓2号 11奇门仓储 12旺店通仓储 13心怡仓储 14力威仓储 15京东沧海 16云集仓储 17POS18虎符 20外部链路型 22抖店云仓 126分销委外仓 127其它
		SubType     *int   `form:"sub_type" mapstructure:"sub_type"`         // 仓库子类型（针对其它仓储 1米氏实体店 2 线下虚拟店）
		IsDisabled  *bool  `form:"is_disabled" mapstructure:"is_disabled"`   // 是否停用 0未停用 1停用 默认为0 返回未停用仓库
		PageNo      *int   `form:"page_no" mapstructure:"page_no"`           // 页码 从0页开始
		PageSize    *int   `form:"page_size" mapstructure:"page_size"`       // 每页返回的数据条数 值范围1~100，默认值为40
	}
	WarehouseQueryResp struct {
		TotalCount    int             `json:"total_count"`
		WarehouseList []WarehouseInfo `json:"warehouses"`
	}
	WarehouseInfo struct {
		WarehouseNo      string `json:"warehouse_no" mapstructure:"warehouse_no"`             // 仓库编码
		Name             string `json:"name" mapstructure:"name"`                             // 仓库名称
		WarehouseType    *int   `json:"warehouse_type" mapstructure:"warehouse_type"`         // 仓库类型 0不限 1普通仓库 2自动流传外部 3京东仓储 4科捷 5百世物流 6SKU360 7通天晓 8中联网仓 9顺丰仓储 10网仓2号 11奇门仓储 12...
		ExtWarehouseNo   string `json:"ext_warehouse_no" mapstructure:"ext_warehouse_no"`     // 外部仓库编号
		Province         string `json:"province" mapstructure:"province"`                     // 省份
		City             string `json:"city" mapstructure:"city"`                             // 城市
		District         string `json:"district" mapstructure:"district"`                     // 区县
		Address          string `json:"address" mapstructure:"address"`                       // 详细地址
		Contact          string `json:"contact" mapstructure:"contact"`                       // 联系人
		Telno            string `json:"telno" mapstructure:"telno"`                           // 固话
		Zip              string `json:"zip" mapstructure:"zip"`                               // 邮编
		Mobile           string `json:"mobile" mapstructure:"mobile"`                         // 手机
		IsDefect         *bool  `json:"is_defect" mapstructure:"is_defect"`                   // 是否残次仓 0不是 1是
		Remark           string `json:"remark" mapstructure:"remark"`                         // 备注
		Prop1            string `json:"prop1" mapstructure:"prop_1"`                          // 自定义1
		Prop2            string `json:"prop2" mapstructure:"prop_2"`                          // 自定义2
		WarehouseId      string `json:"warehouse_id" mapstructure:"warehouse_id"`             // 仓库id
		Type             *int   `json:"type" mapstructure:"type"`                             // 仓库类型 0不限 1普通仓库 2自动流传外部 3京东仓储 4科捷 5百世物流 6SKU360 7通天晓 8中联网仓 9顺丰仓储 10网仓2号 11奇门仓储 12...
		SubType          *int   `json:"sub_type" mapstructure:"sub_type"`                     // 仓库子类型 1米氏实体店 2 线下虚拟店
		ApiKey           string `json:"api_key" mapstructure:"api_key"`                       // API中商家的唯一标识
		ApiObjectId      string `json:"api_object_id" mapstructure:"api_object_id"`           // 和ext_warehouse_no一起保证唯一
		Tag              string `json:"tag" mapstructure:"tag"`                               // 委外仓禁止使用
		DivisionId       string `json:"division_id" mapstructure:"division_id"`               // 区编号
		Flag             string `json:"flag" mapstructure:"flag"`                             // 用于下载标记
		ShopId           string `json:"shop_id" mapstructure:"shop_id"`                       // 实体店仓库店铺id
		IsOuterStock     *bool  `json:"is_outer_stock" mapstructure:"is_outer_stock"`         // 使用外部库存
		IsOuterGoods     *bool  `json:"is_outer_goods" mapstructure:"is_outer_goods"`         // 代销仓库 货品没有所有权
		PickerNum        string `json:"picker_num" mapstructure:"picker_num"`                 // 仓库拣货员数量 智选批次使用
		CoordinatesX     string `json:"coordinates_x" mapstructure:"coordinates_x"`           // 地理坐标_x
		CoordinatesY     string `json:"coordinates_y" mapstructure:"coordinates_y"`           // 地理坐标_y
		MatchWarehouseId string `json:"match_warehouse_id" mapstructure:"match_warehouse_id"` // 对应的残次品/正品仓id
		CodLogisticsId   string `json:"cod_logistics_id" mapstructure:"cod_logistics_id"`     // 货到付款物流公司
		IsDisabled       *bool  `json:"is_disabled" mapstructure:"is_disabled"`               // 是否停用 0未停用 1停用
		Priority         string `json:"priority" mapstructure:"priority"`                     // 仓库的优先级
		Created          string `json:"created" mapstructure:"created"`                       // 创建时间 格式：yyyy-MM-ddHH:mm:ss
		Modified         string `json:"modified" mapstructure:"modified"`                     // 最后修改时间 格式：yyyy-MM-ddHH:mm:ss
	}
)
