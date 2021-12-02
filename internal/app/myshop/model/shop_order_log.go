package model

import "orp/pkg/time"

type TShopOrderLog struct {
	Id             int64         `json:"id" form:"id"`
	ShopId         int64         `json:"shop_id" from:"shop_id"`
	Tid            string        `json:"tid" form:"tid"`
	RefundId       int64         `json:"refund_id"`
	Type           int8          `json:"type" form:"type"`
	CurrentVersion uint8         `json:"current_version"` //255个版本够了
	Content        string        `json:"content" form:"content"`
	CreatedAt      time.JsonTime `json:"created_at" form:"created_at" xorm:"created"`
}
