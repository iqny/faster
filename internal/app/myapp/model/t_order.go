package model

/*type TOrder struct {
	OrderId int64  `json:"order_id"`
	Name    string `json:"name"`
}*/

type TOrder struct {
	OrderSn   int64  `json:"order_sn"`
	UserId    int64  `json:"user_id"`
	Tid       string `json:"tid"`
	Consignee string `json:"consignee"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	District  string `json:"district"`
	Town      string `json:"town"`
	Address   string `json:"address"`
	Tel       string `json:"tel"`
}

type TOrderItem struct {
	Id               int64   `json:"id"`
	OrderSn          int64   `json:"order_sn"`
	UserId           int64   `json:"user_id"`
	Tid              string  `json:"tid"`
	GoodsName        string  `json:"goods_name"`
	GoodsNumber      string  `json:"goods_number"`
	MarketPrice      float64 `json:"market_price"`
	Discount         float64 `json:"discount"`
	TransactionPrice float64 `json:"transaction_price"`
	Oid              string  `json:"oid"`
}
