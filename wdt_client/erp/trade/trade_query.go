package trade

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/convert"
	"github.com/ITITO881/go-zero-service/wdt_client/erp"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/mitchellh/mapstructure"
)

type ApiTradeQueryFunc struct {
	erpClient *erp.WdtClient
}

func NewApiTradeQueryFunc(conf map[string]string) *ApiTradeQueryFunc {
	return &ApiTradeQueryFunc{
		erpClient: erp.NewWdtClient(conf),
	}
}

func (f *ApiTradeQueryFunc) TradeQuery(data *types.TradeQueryReq) *types.TradeQueryResp {
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println("trade_query.php mapstructure.Decode error:", err)
		return nil
	}

	result, err := f.erpClient.Execute("trade_query.php", businessParams)
	if err != nil {
		fmt.Println("请求奇门接口 trade_query.php 错误：", err.Error())
		return nil
	}
	resp := *result
	if resp["code"].(float64) != 0 {
		fmt.Println("查询订单管理失败 trade_query.php：", resp["code"], " 错误信息：", resp["message"])
		return nil
	}
	dataList := make([]types.TradeOrder, 0)
	err = mapstructure.WeakDecode(resp["trades"], &dataList)
	if err != nil {
		fmt.Println("trade_query.php resp mapstructure.Decode error:", err)
		return nil
	}
	return &types.TradeQueryResp{
		TotalCount: int(resp["total_count"].(float64)),
		Trades:     dataList,
	}
}
