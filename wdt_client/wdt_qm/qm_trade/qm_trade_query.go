package qm_trade

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/convert"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/ITITO881/go-zero-service/wdt_client/wdt_qm"
	"github.com/mitchellh/mapstructure"
)

type TradeQueryFunc struct {
	erpClient *wdt_qm.QmClient
}

func NewTradeQueryFunc(conf map[string]string) *TradeQueryFunc {
	return &TradeQueryFunc{
		erpClient: wdt_qm.NewQmClient(conf),
	}
}

func (f *TradeQueryFunc) TradeQuery(data *types.TradeQueryReq) *types.TradeQueryResp {
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	f.erpClient.Method = "wdt.trade.query"
	result, err := f.erpClient.Execute(businessParams)
	if err != nil {
		fmt.Println("请求奇门接口 wdt.trade.query 错误：", err.Error())
		return nil
	}
	resp := (*result)["response"].(map[string]interface{})
	if resp["errorcode"].(float64) != 0 {
		fmt.Println("查询订单管理失败 wdt.trade.query：", resp["code"], " 错误信息：", resp["message"])
		return nil
	}
	dataList := make([]types.TradeOrder, 0)
	err = mapstructure.WeakDecode(resp["trades"], &dataList)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	return &types.TradeQueryResp{
		TotalCount: int(resp["total_count"].(float64)),
		Trades:     dataList,
	}
}
