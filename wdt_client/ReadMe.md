# WDT 接口SDK 说明文档

## 1. 目录结构

```shell
wdt_client
├── erp # 旺店通erp标准接口
  ├── basic # 基础类 目录
  ├── goods # 货品类 目录
  ├── purchase # 采购类 目录
  ├── refund # 售后类 目录 已实现退换单和原始退款单查询
  ├── stock # 库存类 目录
  ├── trade # 订单类 目录 已实现订单查询
  ├── erpClient.go 
├── types # 旺店通接口数据结构
  ├── erpRefund.go # 售后接口 数据结构
  ├── erpStock.go # 库存接口 数据结构
  ├── erpTrade.go # 订单接口 数据结构
├── wdt_qm # 旺店通奇门接口
  ├── qm_refund  # 售后类 目录 已实现退换单和原始退款单查询
  ├── qm_trade # 订单类 目录 已实现订单查询
  ├── qmClient.go
├── wms # 旺店通wms接口
  ├── 暂无
├── ReadMe.md # 说明文档
```

## 2. 旺店通erp标准接口

接口文档：[旺店通erp标准接口](https://open.wangdian.cn/qyb/open/apidoc)

sdk使用示例：

```go
package main

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/wdt_client/erp/trade"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
)

func main() {
	client := trade.NewApiTradeQueryFunc(map[string]string{
		"Sid":       "sid",
		"AppKey":    "appKey",
		"AppSecret": "appSecret",
		"BaseUrl":   "http://api.wangdian.cn/openapi2/",
	})
	res := client.TradeQuery(&types.TradeQueryReq{
		SrcTid: "4028630709320530926",
	})
	if res != nil {
		fmt.Println(res.TotalCount)
	}
}
```

## 3. 旺店通奇门接口

接口文档：[参考旺店通erp标准接口](https://open.wangdian.cn/qyb/open/apidoc)

sdk使用示例：

```go
package main

import (
	"fmt"
	"github.com/ITITO881/go-zero-service/wdt_client/types"
	"github.com/ITITO881/go-zero-service/wdt_client/wdt_qm/qm_trade"
)

func main() {
	client := qm_trade.NewTradeQueryFunc(map[string]string{
		"Sid":          "sid",
		"AppKey":       "appKey",
		"TargetAppKey": "targetAppKey",
		"AppSecret":    "appSecret",
		"BaseUrl":      "http://hu3cgwt0tc.api.taobao.com/router/qm",
	})
	res := client.TradeQuery(&types.TradeQueryReq{
		SrcTid: "4028630709320530926",
	})
	if res != nil {
		fmt.Println(res.TotalCount)
	}
}
```

## 4. 旺店通wms接口

待更新