package purchase

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/convert"
	"github.com/ITITO881/go-zero-service/wdt_client/erp"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/mitchellh/mapstructure"
)

type ApiPurchaseQueryFunc struct {
	erpClient *erp.WdtClient
}

func NewApiPurchaseQueryFunc(conf map[string]string) *ApiPurchaseQueryFunc {
	return &ApiPurchaseQueryFunc{erpClient: erp.NewWdtClient(conf)}
}

func (f *ApiPurchaseQueryFunc) PurchaseQuery(data *types.PurchaseOrderQueryReq) *types.PurchaseOrderQueryResp {
	method := "purchase_order_query.php"
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println(method, " mapstructure.Decode error:", err)
		return nil
	}
	result, err := f.erpClient.Execute(method, businessParams)
	if err != nil {
		fmt.Println("请求奇门接口 ,"+method+" 错误：", err.Error())
		return nil
	}
	resp := *result
	if resp["code"].(float64) != 0 {
		fmt.Println("查询采购单失败 ", method, ": ", resp["code"], " 错误信息：", resp["message"])
		return nil
	}
	dataList := make([]types.PurchaseData, 0)
	err = mapstructure.WeakDecode(resp["purchase_list"], &dataList)
	if err != nil {
		fmt.Println(method, " resp mapstructure.Decode error:", err)
		return nil
	}
	return &types.PurchaseOrderQueryResp{
		TotalCount:   int(resp["total_count"].(float64)),
		PurchaseList: dataList,
	}
}
