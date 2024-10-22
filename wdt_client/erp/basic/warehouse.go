package basic

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/convert"
	"github.com/ITITO881/go-zero-service/wdt_client/erp"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/mitchellh/mapstructure"
)

type WarehouseQueryFunc struct {
	erpClient *erp.WdtClient
	method    string
}

func NewWarehouseQueryFunc(conf map[string]string) *WarehouseQueryFunc {
	return &WarehouseQueryFunc{
		erpClient: erp.NewWdtClient(conf),
		method:    "warehouse_query.php",
	}
}

func (f *WarehouseQueryFunc) WarehouseQuery(data *types.WarehouseQueryReq) *types.WarehouseQueryResp {
	data.PageSize = new(int)
	*data.PageSize = 100
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println("WarehouseQuery mapstructure.Decode error:", err)
		return nil
	}
	resp, err := f.erpClient.Execute(f.method, businessParams)
	if err != nil {
		fmt.Println("请求旺店通接口 warehouse_query.php 错误：", err.Error())
		return nil
	}
	if (*resp)["code"].(float64) != 0 {
		fmt.Println("查询查询仓库失败：", (*resp)["code"], " 错误信息：", (*resp)["message"])
		return nil
	}
	dataList := make([]types.WarehouseInfo, 0)
	err = mapstructure.WeakDecode((*resp)["warehouses"], &dataList)
	if err != nil {
		fmt.Println("mapstructure.Decode WarehouseQueryResp error:", err)
		return nil
	}
	return &types.WarehouseQueryResp{
		TotalCount:    int((*resp)["total_count"].(float64)),
		WarehouseList: dataList,
	}
}
