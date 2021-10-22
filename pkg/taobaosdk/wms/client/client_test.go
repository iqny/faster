package client

import (
	"encoding/xml"
	"fmt"
	"orp/pkg/taobaosdk/wms/request/deliveryordercreate"
	"orp/pkg/taobaosdk/wms/request/entryordercreate"
	"orp/pkg/taobaosdk/wms/request/ordercancel"
	"orp/pkg/taobaosdk/wms/response"
	"testing"
)

func withAppKey(c *Config) {
	c.appKey = ""
}
func withAppSecret(c *Config) {
	c.appSecret = ""
}
func withCustomerId(c *Config) {
	c.customerId = ""
}
func withGatewayUrl(c *Config) {
	c.gatewayUrl = "https://qimen.api.taobao.com/router/qmtest"
}
func TestNew(t *testing.T) {
	cli := New(withAppKey, withAppSecret, withCustomerId, withGatewayUrl)
	request := &deliveryordercreate.TaoBaoQiMenDeliveryOrderCreate{Version: "1"}
	invoice := deliveryordercreate.Invoice{}
	deliveryOrder := deliveryordercreate.DeliveryOrder{}
	deliveryOrder.SetExtendProps(deliveryordercreate.ExtendProps{})
	item := deliveryordercreate.Item{Sku: "1"}
	item2 := deliveryordercreate.Item{Sku: "2"}
	detail := deliveryordercreate.Detail{}
	detail.Append(item)
	detail.Append(item2)
	invoice.Detail = detail
	invoices := deliveryordercreate.Invoices{}
	invoices.Append(invoice)
	invoices.Append(invoice)
	deliveryOrder.SetInvoices(invoices)
	deliveryOrder.SetInvoices(invoices)
	orderLines := deliveryordercreate.OrderLines{}
	orderLines.Append(deliveryordercreate.OrderLine{OrderLineNo: "1"})
	orderLines.Append(deliveryordercreate.OrderLine{OrderLineNo: "2"})
	request.SetOrderLines(orderLines)
	request.SetDeliveryOrder(deliveryOrder)
	res, err := cli.Execute(request)
	if err == nil {
		switch res.(type) {
		case response.ErrResponse:
			res := res.(response.ErrResponse)
			fmt.Println(res.Message, err)
		case response.SuccessResponse:
			res := res.(response.SuccessResponse)
			fmt.Println(res.Req, err)
		}
	} else {
		r := res.(response.ErrResponse)
		fmt.Println(r.Message, err)
	}
}
func TestEntryOrder(t *testing.T) {
	cli := New(withAppKey, withAppSecret, withCustomerId)
	request := &entryordercreate.TaoBaoQiMenEntryOrderCreate{Version: "1"}
	entryOrder := entryordercreate.EntryOrder{}
	relatedOrders := entryordercreate.RelatedOrders{}
	relatedOrders.Append(entryordercreate.RelatedOrder{})
	relatedOrders.Append(entryordercreate.RelatedOrder{})
	entryOrder.RelatedOrders = relatedOrders
	request.EntryOrder = entryOrder
	orderLines := entryordercreate.OrderLines{}
	snList := entryordercreate.SnList{Sn: []string{"aa", "bb"}}
	orderLines.Append(entryordercreate.OrderLine{OrderLineNo: "1", SnList: snList})
	orderLines.Append(entryordercreate.OrderLine{OrderLineNo: "2", SnList: snList})

	request.SetOrderLines(orderLines)
	res, err := cli.Execute(request)
	if err == nil {
		switch res.(type) {
		case response.ErrResponse:
			res := res.(response.ErrResponse)
			fmt.Println(res.Message, err)
		case response.SuccessResponse:
			res := res.(response.SuccessResponse)
			fmt.Println(res.Req, err)
		}
	}
}
func TestOrderCancel(t *testing.T) {
	cli := New(withAppKey, withAppSecret, withCustomerId)
	request := &ordercancel.TaoBaoQiMenOrderCancel{
		XMLName:       xml.Name{},
		WarehouseCode: "TT",
		OwnerCode:     "STcs1",
		OrderCode:     "B2BR7",
		OrderId:       "0000000001",
		OrderType:     "B2BR",
		CancelReason:  "",
	}
	res, err := cli.Execute(request)
	if err == nil {
		switch res.(type) {
		case response.ErrResponse:
			res := res.(response.ErrResponse)
			fmt.Println(res.Message, err)
		case response.SuccessResponse:
			res := res.(response.SuccessResponse)
			fmt.Println(res.Req, err)
		}
	}
}
func BenchmarkName(b *testing.B) {
	b.ResetTimer()
	cli := New(withAppKey, withAppSecret, withCustomerId)

	for i := 0; i < b.N; i++ {
		request := &deliveryordercreate.TaoBaoQiMenDeliveryOrderCreate{}
		res, err := cli.Execute(request)
		if err == nil {
			switch res.(type) {
			case response.ErrResponse:
				//res := res.(response.ErrResponse)
				//fmt.Println(res.Message, err)
			case response.SuccessResponse:
				//res := res.(response.SuccessResponse)
				//fmt.Println(res.Req, err)
			}
		}
	}
}
