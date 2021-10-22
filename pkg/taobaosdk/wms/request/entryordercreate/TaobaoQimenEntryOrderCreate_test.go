package entryordercreate

import "testing"

func TestCreate(t *testing.T) {
	request := &TaoBaoQiMenEntryOrderCreate{Version: "1"}
	request.EntryOrder = EntryOrder{}
	relatedOrders :=RelatedOrders{}
	relatedOrders.Append(RelatedOrder{})
	relatedOrders.Append(RelatedOrder{})
	request.EntryOrder.RelatedOrders = relatedOrders
	orderLines := OrderLines{}
	orderLines.Append(OrderLine{OrderLineNo: "1"})
	orderLines.Append(OrderLine{OrderLineNo: "2"})
	request.SetOrderLines(orderLines)
}

func BenchmarkCreate(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {

	}
}
