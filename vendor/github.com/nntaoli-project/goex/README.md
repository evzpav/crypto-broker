<div align="center">
<img width="409" heigth="205" src="https://upload-images.jianshu.io/upload_images/6760989-dec7dc747846880e.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240"  alt="goex">
<img src="https://travis-ci.org/nntaoli-project/goex.svg?branch=dev"/>
</div>

### goex目标

goex项目是为了统一并标准化各个数字资产交易平台的接口而设计，同一个策略可以随时切换到任意一个交易平台，而不需要更改任何代码。

[English](https://github.com/nntaoli-project/goex/blob/dev/README_en.md)

### goex已支持交易所 `22+`

| 交易所 | 行情接口 | 交易接口 | 版本号 |   
| ---   | ---     | ---     | ---   |  
| hbg.com | Y | Y | 1 |   
| hbdm.com | Y| Y |  1 |    
| okex.com  | Y | Y | 3 |
| binance.com | Y | Y | 1 |  
| bitstamp.net | Y | Y | 1 |  
| bitfinex.com | Y | Y | 1 |  
| zb.com | Y | Y | 1 |  
| kraken.com | Y | Y | * |  
| poloniex.com | Y | Y | * |  
| aacoin.com | Y | Y | 1 |   
| allcoin.ca | Y | Y | * |  
| big.one | Y | Y | 2\|3 | 
| fcoin.com | Y | Y | 2 |  
| hitbtc.com | Y | Y | * |
| coinex.com | Y | Y | 1 |
| exx.com | Y | Y | 1 |
| bithumb.com | Y | Y | * |
| gate.io | Y | N | 1 |
| btcbox.co.jp | Y | N | * |
| coinbig.com | Y | Y | * |
|coinbene.com|Y|Y|*|

### 安装goex库  

``` go get github.com/nntaoli-project/goex ```

>建议go mod 管理依赖
``` 
require (
          github.com/nntaoli-project/goex v1.0.4
)
```

### 例子

```golang

   package main
   
   import (
   	"github.com/nntaoli-project/goex"
   	"github.com/nntaoli-project/goex/builder"
   	"log"
   	"time"
   )
   
   func main() {
   	apiBuilder := builder.NewAPIBuilder().HttpTimeout(5 * time.Second)
   	//apiBuilder := builder.NewAPIBuilder().HttpTimeout(5 * time.Second).HttpProxy("socks5://127.0.0.1:1080")
   	
   	//build spot api
   	//api := apiBuilder.APIKey("").APISecretkey("").ClientID("123").Build(goex.BITSTAMP)
   	api := apiBuilder.APIKey("").APISecretkey("").Build(goex.HUOBI_PRO)
   	log.Println(api.GetExchangeName())
   	log.Println(api.GetTicker(goex.BTC_USD))
   	log.Println(api.GetDepth(2, goex.BTC_USD))
   	//log.Println(api.GetAccount())
   	//log.Println(api.GetUnfinishOrders(goex.BTC_USD))
   
   	//build future api
   	futureApi := apiBuilder.APIKey("").APISecretkey("").BuildFuture(goex.HBDM)
   	log.Println(futureApi.GetExchangeName())
   	log.Println(futureApi.GetFutureTicker(goex.BTC_USD, goex.QUARTER_CONTRACT))
   	log.Println(futureApi.GetFutureDepth(goex.BTC_USD, goex.QUARTER_CONTRACT, 5))
   	//log.Println(futureApi.GetFutureUserinfo()) // account
   	//log.Println(futureApi.GetFuturePosition(goex.BTC_USD , goex.QUARTER_CONTRACT))//position info
   }

```

### websocket 使用例子

```golang
import (
	"github.com/nntaoli-project/goex"
	"github.com/nntaoli-project/goex/huobi"
	//"github.com/nntaoli-project/goex/okcoin"
	"log"
)

func main() {

	//ws := okcoin.NewOKExFutureWs() //ok期货
	ws := huobi.NewHbdmWs() //huobi期货
	//设置回调函数
	ws.SetCallbacks(func(ticker *goex.FutureTicker) {
		log.Println(ticker)
	}, func(depth *goex.Depth) {
		log.Println(depth)
	}, func(trade *goex.Trade, contract string) {
		log.Println(contract, trade)
	})
	//订阅行情
	ws.SubscribeTrade(goex.BTC_USDT, goex.NEXT_WEEK_CONTRACT)
	ws.SubscribeDepth(goex.BTC_USDT, goex.QUARTER_CONTRACT, 5)
	ws.SubscribeTicker(goex.BTC_USDT, goex.QUARTER_CONTRACT)
}  

```

### 更多文档

[goex.TOP](https://goex.top)

### 注意事项

1. 推荐使用GoLand开发。
2. 推荐关闭自动格式化功能,代码请使用go fmt 格式化.
3. 不建议对现已存在的文件进行重新格式化，这样会导致commit特别糟糕。
4. 请用OrderID2这个字段代替OrderID
5. 请不要使用deprecated关键字标注的方法和字段，后面版本可能随时删除的
6. 交流QQ群：574829125
-----------------

donate
-----------------
BTC:13cBHLk6B7t3Uj7caJbCwv1UaiuiA6Qx8z

LTC:LVxM7y1K2dnpuNBU42ei3dKzPySf4VAm1H
 
ETH:0x98573ddb33cdddce480c3bc1f9279ccd88ca1e93

### 欢迎为作者付一碗面钱

<img src="https://raw.githubusercontent.com/nntaoli-project/goex/dev/wx_pay.JPG" width="250" alt="一碗面钱">&nbsp;&nbsp;&nbsp;<img src="https://raw.githubusercontent.com/nntaoli-project/goex/dev/IMG_1177.jpg" width="250" alt="一碗面钱">
