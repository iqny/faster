package ordercancel

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestOrderCancel(t *testing.T) {
	request:= &TaoBaoQiMenOrderCancel{
		XMLName:       xml.Name{},
		WarehouseCode: "TT1",
		OwnerCode:     "SIT",
		OrderCode:     "B2BR7",
		OrderId:       "",
		OrderType:     "B2BR",
		CancelReason:  "",
	}
	fmt.Printf(request.ToXML())
}
