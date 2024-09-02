package types

type (
	VipStockOutsideWmsQueryReq struct {
		ApiOuterNo  string `form:"api_outer_no,optional,omitempty" mapstructure:"api_outer_no,omitempty"`
		WarehouseNo string `form:"warehouse_no,optional,omitempty" mapstructure:"warehouse_no,omitempty"`
		Status      string `form:"status,optional,omitempty" mapstructure:"status,omitempty"`
		OrderType   string `form:"order_type,optional,omitempty" mapstructure:"order_type,omitempty"`
		OuterNo     string `form:"outer_no,optional,omitempty" mapstructure:"outer_no,omitempty"`
		OrderNo     string `form:"order_no,optional,omitempty" mapstructure:"order_no,omitempty"`
		StartTime   string `form:"start_time,optional,omitempty" mapstructure:"start_time,omitempty"`
		EndTime     string `form:"end_time,optional,omitempty" mapstructure:"end_time,omitempty"`
		PageSize    string `form:"page_size,optional,omitempty" mapstructure:"page_size,omitempty"`
		PageNo      string `form:"page_no,optional,omitempty" mapstructure:"page_no,omitempty"`
	}
	VipStockOutsideWmsQueryResp struct {
		TotalCount int               `json:"total_count"`
		OrderList  []StockOutsideWms `json:"order_list"`
	}
	StockOutsideWms struct {
		OrderID           int                     `mapstructure:"order_id" json:"order_id,omitempty"`
		OrderNo           string                  `mapstructure:"order_no" json:"order_no,omitempty"`
		WarehouseNo       string                  `mapstructure:"warehouse_no" json:"warehouse_no,omitempty"`
		ReceiverCountry   int                     `mapstructure:"receiver_country" json:"receiver_country,omitempty"`
		NoteCount         int                     `mapstructure:"note_count" json:"note_count,omitempty"`
		Status            int                     `mapstructure:"status" json:"status,omitempty"`
		WmsStatus         int                     `mapstructure:"wms_status" json:"wms_status,omitempty"`
		ErrorInfo         string                  `mapstructure:"error_info" json:"error_info,omitempty"`
		OperatorID        int                     `mapstructure:"operator_id" json:"operator_id,omitempty"`
		CheckOperatorID   int                     `mapstructure:"check_operator_id" json:"check_operator_id,omitempty"`
		WarehouseID       int                     `mapstructure:"warehouse_id" json:"warehouse_id,omitempty"`
		WarehouseType     int                     `mapstructure:"warehouse_type" json:"warehouse_type,omitempty"`
		TransportMode     int                     `mapstructure:"transport_mode" json:"transport_mode,omitempty"`
		LogisticsID       int                     `mapstructure:"logistics_id" json:"logistics_id,omitempty"`
		LogisticsFee      float64                 `mapstructure:"logistics_fee" json:"logistics_fee,omitempty"`
		Created           *string                 `mapstructure:"created" json:"created,omitempty"`
		Remark            string                  `mapstructure:"remark" json:"remark,omitempty"`
		OrderType         int                     `mapstructure:"order_type" json:"order_type,omitempty"`
		OtherFee          float64                 `mapstructure:"other_fee" json:"other_fee,omitempty"`
		GoodsCount        float64                 `mapstructure:"goods_count" json:"goods_count,omitempty"`
		GoodsTypeCount    float64                 `mapstructure:"goods_type_count" json:"goods_type_count,omitempty"`
		GoodsRealCount    float64                 `mapstructure:"goods_real_count" json:"goods_real_count,omitempty"`
		CheckTime         *string                 `mapstructure:"check_time" json:"check_time,omitempty"`
		Modified          *string                 `mapstructure:"modified" json:"modified,omitempty"`
		ReceiverName      string                  `mapstructure:"receiver_name" json:"receiver_name,omitempty"`
		ReceiverProvince  int                     `mapstructure:"receiver_province" json:"receiver_province,omitempty"`
		ReceiverCity      int                     `mapstructure:"receiver_city" json:"receiver_city,omitempty"`
		ReceiverDistrict  int                     `mapstructure:"receiver_district" json:"receiver_district,omitempty"`
		ReceiverZip       string                  `mapstructure:"receiver_zip" json:"receiver_zip,omitempty"`
		ReceiverArea      string                  `mapstructure:"receiver_area" json:"receiver_area,omitempty"`
		ReceiverAddress   string                  `mapstructure:"receiver_address" json:"receiver_address,omitempty"`
		ReceiverMobile    string                  `mapstructure:"receiver_mobile" json:"receiver_mobile,omitempty"`
		ReceiverTelno     string                  `mapstructure:"receiver_telno" json:"receiver_telno,omitempty"`
		Prop1             string                  `mapstructure:"prop1" json:"prop1,omitempty"`
		Prop2             string                  `mapstructure:"prop2" json:"prop2,omitempty"`
		Prop3             string                  `mapstructure:"prop3" json:"prop3,omitempty"`
		Prop4             string                  `mapstructure:"prop4" json:"prop4,omitempty"`
		ApiOuterNo        string                  `mapstructure:"api_outer_no" json:"api_outer_no,omitempty"`
		OuterNo           string                  `mapstructure:"outer_no" json:"outer_no,omitempty"`
		WmsOuterNo        string                  `mapstructure:"wms_outer_no" json:"wms_outer_no,omitempty"`
		LogisticsNo       string                  `mapstructure:"logistics_no" json:"logistics_no,omitempty"`
		CheckOperatorName string                  `mapstructure:"check_operator_name" json:"check_operator_name,omitempty"`
		OperatorName      string                  `mapstructure:"operator_name" json:"operator_name,omitempty"`
		DetailsList       []StockOutsideWmsDetail `mapstructure:"details_list" json:"details_list,omitempty"`
	}

	StockOutsideWmsDetail struct {
		RecID       int     `mapstructure:"rec_id" json:"rec_id,omitempty"`
		OrderNo     string  `mapstructure:"order_no" json:"order_no,omitempty"`
		SpecID      int     `mapstructure:"spec_id" json:"spec_id,omitempty"`
		Num2        float64 `mapstructure:"num2" json:"num2,omitempty"`
		BaseUnitID  int     `mapstructure:"base_unit_id" json:"base_unit_id,omitempty"`
		UnitRatio   float64 `mapstructure:"unit_ratio" json:"unit_ratio,omitempty"`
		BatchID     int     `mapstructure:"batch_id" json:"batch_id,omitempty"`
		Created     *string `mapstructure:"created" json:"created,omitempty"`
		Modified    *string `mapstructure:"modified" json:"modified,omitempty"`
		GoodsID     int     `mapstructure:"goods_id" json:"goods_id,omitempty"`
		ShortName   string  `mapstructure:"short_name" json:"short_name,omitempty"`
		SpecNo      string  `mapstructure:"spec_no" json:"spec_no,omitempty"`
		Barcode     string  `mapstructure:"barcode" json:"barcode,omitempty"`
		UnitID      int     `mapstructure:"unit_id" json:"unit_id,omitempty"`
		BrandID     int     `mapstructure:"brand_id" json:"brand_id,omitempty"`
		BatchNo     string  `mapstructure:"batch_no" json:"batch_no,omitempty"`
		BatchRemark string  `mapstructure:"batch_remark" json:"batch_remark,omitempty"`
		PositionID  string  `mapstructure:"position_id" json:"position_id,omitempty"`
		PositionNo  string  `mapstructure:"position_no" json:"position_no,omitempty"`
		Num         float64 `mapstructure:"num" json:"num,omitempty"`
		InoutNum    float64 `mapstructure:"inout_num" json:"inout_num,omitempty"`
		Price       float64 `mapstructure:"price" json:"price,omitempty"`
		Remark      string  `mapstructure:"remark" json:"remark,omitempty"`
		LogisticsID int     `mapstructure:"logistics_id" json:"logistics_id,omitempty"`
		LogisticsNo string  `mapstructure:"logistics_no" json:"logistics_no,omitempty"`
		BrandName   string  `mapstructure:"brand_name" json:"brand_name,omitempty"`
		GoodsName   string  `mapstructure:"goods_name" json:"goods_name,omitempty"`
		GoodsNo     string  `mapstructure:"goods_no" json:"goods_no,omitempty"`
		SpecName    string  `mapstructure:"spec_name" json:"spec_name,omitempty"`
		SpecCode    string  `mapstructure:"spec_code" json:"spec_code,omitempty"`
		GoodsUnit   string  `mapstructure:"goods_unit" json:"goods_unit,omitempty"`
		ExpireDate  *string `mapstructure:"expire_date" json:"expire_date,omitempty"`
	}
)
