package singleitemsynchronize

import (
	"fmt"
	"testing"
)

func TestTaoBaoQimEnSingleItemSynchronize_ToXML(t *testing.T) {
	request:=&TaoBaoQimEnSingleItemSynchronize{}
	item:=Item{
		ItemCode:        "",
		ItemId:          "",
		GoodsCode:       "",
		ItemName:        "",
		ShortName:       "",
		EnglishName:     "",
		BarCode:         "",
		SkuProperty:     "",
		StockUnit:       "",
		Length:          "",
		Width:           "",
		Height:          "",
		Volume:          "",
		GrossWeight:     "",
		NetWeight:       "",
		Color:           "",
		Size:            "",
		Title:           "",
		CategoryId:      "",
		CategoryName:    "",
		PricingCategory: "",
		SafetyStock:     0,
		ItemType:        "",
		TagPrice:        "",
		RetailPrice:     "",
		CostPrice:       "",
		PurchasePrice:   "",
		SeasonCode:      "",
		SeasonName:      "",
		BrandCode:       "",
		BrandName:       "",
		IsSNMgmt:        "",
		ProductDate:     "",
		ExpireDate:      "",
		IsShelfLifeMgmt: "",
		ShelfLife:       0,
		RejectLifecycle: 0,
		LockupLifecycle: 0,
		AdventLifecycle: 0,
		IsBatchMgmt:     "",
		BatchCode:       "",
		BatchRemark:     "",
		PackCode:        "",
		Pcs:             "",
		OriginAddress:   "",
		ApprovalNumber:  "",
		IsFragile:       "",
		IsHazardous:     "",
		Remark:          "",
		CreateTime:      "",
		UpdateTime:      "",
		IsValid:         "",
		IsSku:           "",
		PackageMaterial: "",
		SupplierCode:    "",
		LogisticsType:   "",
		IsLiquid:        "",
	}
	request.SetItem(item)
	fmt.Printf("%s\n",request.ToXML())
}
