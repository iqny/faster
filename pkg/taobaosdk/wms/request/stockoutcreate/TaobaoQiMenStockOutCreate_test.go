package stockoutcreate

import (
	"fmt"
	"testing"
)

func TestToXml(t *testing.T) {
	request := &TaoBaoQiMenStockOutCreate{Version: "1"}
	orderLines := OrderLines{}
	orderLines.Append(OrderLine{
		OutBizCode:         "",
		OrderLineNo:        "1",
		SourceOrderCode:    "",
		SubSourceOrderCode: "",
		PayNo:              "",
		OwnerCode:          "",
		ItemCode:           "",
		ItemId:             "",
		InventoryType:      "",
		ItemName:           "",
		ExtCode:            "",
		PlanQty:            0,
		RetailPrice:        "",
		ActualPrice:        "",
		DiscountAmount:     "",
		BatchCode:          "",
		ProductDate:        "",
		ExpireDate:         "",
		ProduceCode:        "",
	})
	orderLines.Append(OrderLine{
		OutBizCode:         "",
		OrderLineNo:        "2",
		SourceOrderCode:    "",
		SubSourceOrderCode: "",
		PayNo:              "",
		OwnerCode:          "",
		ItemCode:           "",
		ItemId:             "",
		InventoryType:      "",
		ItemName:           "",
		ExtCode:            "",
		PlanQty:            0,
		RetailPrice:        "",
		ActualPrice:        "",
		DiscountAmount:     "",
		BatchCode:          "",
		ProductDate:        "",
		ExpireDate:         "",
		ProduceCode:        "",
	})
	deliveryOrder := DeliveryOrder{
		TotalOrderLines:      len(orderLines.OrderLine),
		DeliveryOrderCode:    "",
		OrderType:            "",
		RelatedOrders:        "",
		WarehouseCode:        "",
		CreateTime:           "",
		ScheduleDate:         "",
		LogisticsCode:        "",
		LogisticsName:        "",
		SupplierCode:         "",
		SupplierName:         "",
		TransportMode:        "",
		PickerInfo:           PickerInfo{},
		SenderInfo:           SenderInfo{},
		ReceiverInfo:         ReceiverInfo{},
		Remark:               "",
		OrderSourceType:      "",
		ReceivingTime:        "",
		ShippingTime:         "",
		TargetWarehouseName:  "",
		TargetWarehouseCode:  "",
		TargetEntryOrderCode: "",
		WarehouseName:        "",
		ExtendProps:          ExtendProps{},
	}
	deliveryOrder.SetExtendProps(ExtendProps{})

	request.SetOrderLines(orderLines)
	request.SetDeliveryOrder(deliveryOrder)
	fmt.Println(request.ToXML())
}
