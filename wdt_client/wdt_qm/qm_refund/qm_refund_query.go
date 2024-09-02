package qm_refund

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/convert"
	"github.com/ITITO881/go-zero-service/wdt_client/erp"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/mitchellh/mapstructure"
)

type ApiRefundFunc struct {
	erpClient *erp.WdtClient
}

func NewApiRefundFunc(conf map[string]string) *ApiRefundFunc {
	return &ApiRefundFunc{
		erpClient: erp.NewWdtClient(conf),
	}
}

func (f *ApiRefundFunc) VipApiRefundQuery(data *types.VipApiRefundQueryReq) *types.VipApiRefundQueryResp {
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	resp, err := f.erpClient.Execute("vip_api_refund_query.php", businessParams)
	if err != nil {
		fmt.Println("请求旺店通接口 vip_api_refund_query.php 错误：", err.Error())
		return nil
	}
	if (*resp)["code"].(float64) != 0 {
		fmt.Println("查询原始退款单失败：", (*resp)["code"], " 错误信息：", (*resp)["message"])
		return nil
	}
	dataList := make([]types.VipApiRefundTrade, 0)
	err = mapstructure.WeakDecode((*resp)["trade_list"], &dataList)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	return &types.VipApiRefundQueryResp{
		TotalCount: int((*resp)["total_count"].(float64)),
		TradeList:  dataList,
	}
}

func (f *ApiRefundFunc) RefundQuery(data *types.RefundQueryReq) *types.RefundQueryResp {
	params := convert.ReflectTag{Value: data}.GetMapOmitempty("form")
	businessParams := make(map[string]string)
	err := mapstructure.WeakDecode(params, &businessParams)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	resp, err := f.erpClient.Execute("refund_query.php", businessParams)
	if err != nil {
		fmt.Println("请求旺店通接口 refund_query.php 错误：", err.Error())
		return nil
	}
	if (*resp)["code"].(float64) != 0 {
		fmt.Println("查询退换管理失败：", (*resp)["code"], " 错误信息：", (*resp)["message"])
		return nil
	}
	dataList := make([]types.RefundData, 0)
	err = mapstructure.WeakDecode((*resp)["refunds"], &dataList)
	if err != nil {
		fmt.Println("mapstructure.Decode error:", err)
		return nil
	}
	return &types.RefundQueryResp{
		TotalCount: int((*resp)["total_count"].(float64)),
		Refunds:    dataList,
	}
}
