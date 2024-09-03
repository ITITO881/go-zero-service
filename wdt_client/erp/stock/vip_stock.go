package stock

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/convert"
	"github.com/ITITO881/go-zero-service/wdt_client/erp"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/mitchellh/mapstructure"
)

type VipStockOutsideWmsFunc struct {
	erpClient *erp.WdtClient
	method    string
}

func NewVipStockOutsideWmsFunc(conf map[string]string) *VipStockOutsideWmsFunc {
	return &VipStockOutsideWmsFunc{
		erpClient: erp.NewWdtClient(conf),
		method:    "vip_stock_outside_wms_query.php",
	}
}

func (f *VipStockOutsideWmsFunc) VipStockOutsideWmsQuery(data *types.VipStockOutsideWmsQueryReq) *types.VipStockOutsideWmsQueryResp {
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	resp, err := f.erpClient.Execute(f.method, businessParams)
	if err != nil {
		fmt.Println("请求旺店通接口 vip_stock_outside_wms_query.php 错误：", err.Error())
		return nil
	}
	if (*resp)["code"].(float64) != 0 {
		fmt.Println("查询委外出入口单失败：", (*resp)["code"], " 错误信息：", (*resp)["message"])
		return nil
	}
	dataList := make([]types.StockOutsideWms, 0)
	err = mapstructure.WeakDecode((*resp)["order_list"], &dataList)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	return &types.VipStockOutsideWmsQueryResp{
		TotalCount: int((*resp)["total_count"].(float64)),
		OrderList:  dataList,
	}
}
