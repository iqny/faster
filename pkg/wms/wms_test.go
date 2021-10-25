package wms

import (
	"fmt"
	"orp/pkg/wms/interface_factory"
	"orp/pkg/wms/taotao/request/deliveryordercreate"
	"testing"
)

func TestCreateWms(t *testing.T) {
	request := &deliveryordercreate.TaoBaoQiMenDeliveryOrderCreate{Version: "1"}
	wms:=New("taobao")
	res, err := wms.Execute(request)
	if err == nil {
		switch res.(type) {
		case interface_factory.ErrResponse:
			res := res.(interface_factory.ErrResponse)
			fmt.Println(res.Message, err)
		case interface_factory.SuccessResponse:
			res := res.(interface_factory.SuccessResponse)
			fmt.Println(res.Res, err)
		}
	} else {
		r := res.(interface_factory.ErrResponse)
		fmt.Println(r.Message, err)
	}
}
