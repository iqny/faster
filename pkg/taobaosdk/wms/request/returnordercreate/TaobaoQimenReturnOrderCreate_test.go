package returnordercreate

import (
	"fmt"
	"testing"
)

func TestToXML(t *testing.T) {
	request := &TaoBaoQiMenReturnOrderCreate{}
	orderLists :=OrderLines{}
	orderLists.Append(OrderLine{OrderLineNo: "1",SnList: SnList{Sn: []string{"1","2"}},PlanQty: 1})
	orderLists.Append(OrderLine{OrderLineNo: "2",SnList: SnList{Sn: []string{"3","4"}},PlanQty: 2})
	request.SetOrderLines(orderLists)
	returnOrder:=ReturnOrder{
		ReturnOrderCode:      "",
		WarehouseCode:        "",
		OrderType:            "",
		OrderFlag:            "",
		PreDeliveryOrderCode: "",
		PreDeliveryOrderId:   "",
		LogisticsCode:        "",
		LogisticsName:        "",
		ExpressCode:          "",
		ReturnReason:         "",
		BuyerNick:            "",
		SenderInfo:           SenderInfo{},
		Remark:               "",
		SourcePlatformCode:   "",
		SourcePlatformName:   "",
		ShopNick:             "",
		SellerNick:           "",
	}
	request.SetReturnOrder(returnOrder)
	fmt.Println(request.ToXML())
}
