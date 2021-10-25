package deliveryordercreate

import (
	"fmt"
	"testing"
)

func TestToXml(t *testing.T) {
	request := &TaoBaoQiMenDeliveryOrderCreate{Version: "1"}
	invoice := Invoice{}
	deliveryOrder := DeliveryOrder{}
	deliveryOrder.SetExtendProps(ExtendProps{})
	item := Item{Sku: "1"}
	item2 := Item{Sku: "2"}
	detail := Detail{}
	detail.Append(item)
	detail.Append(item2)
	invoice.Detail = detail
	invoices := Invoices{}
	invoices.Append(invoice)
	invoices.Append(invoice)
	deliveryOrder.SetInvoices(invoices)
	deliveryOrder.SetInvoices(invoices)
	orderLines := OrderLines{}
	orderLines.Append(OrderLine{OrderLineNo: "1"})
	orderLines.Append(OrderLine{OrderLineNo: "2"})
	request.SetOrderLines(orderLines)
	request.SetDeliveryOrder(deliveryOrder)
	fmt.Println(request.ToXML())
}
