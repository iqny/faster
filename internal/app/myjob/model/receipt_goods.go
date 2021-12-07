package model

type ReceiptGoods struct {
	ReceiptId     int64  `json:"receipt_id"`
	InventoryType int    `json:"inventory_type"`
	GoodsId       int64  `json:"goods_id"`
	SkuNo         string `json:"sku_no"`
	Number        int64  `json:"number"`
	Batch         int64  `json:"batch"`
}
