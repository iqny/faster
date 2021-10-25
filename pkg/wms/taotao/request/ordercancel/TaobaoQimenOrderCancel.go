package ordercancel

//单据取消
import (
	"encoding/xml"
	"fmt"
	"orp/pkg/wms/interface_factory"
	"strings"
)

const method = "taobao.qimen.order.cancel"

type TaoBaoQiMenOrderCancel struct {
	XMLName       xml.Name `xml:"request"`
	WarehouseCode string   `xml:"warehouseCode"` //	W1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OwnerCode     string   `xml:"ownerCode"`     //	H1234	货主编码
	OrderCode     string   `xml:"orderCode"`     //	O1234	单据编码
	OrderId       string   `xml:"orderId"`       //	WQ1234	仓储系统单据编码
	OrderType     string   `xml:"orderType"`     //	JYCK	单据类型(JYCK=一般交易出库单;HHCK= 换货出库;BFCK=补发出库;PTCK=普通出库单;DBCK=调拨出库;B2BRK=B2B入库;B2BCK=B2B出库;QTCK=其他出库;SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK= 调拨入库;QTRK=其他入库;XTRK= 销退入库;THRK=退货入库;HHRK= 换货入库;CNJG= 仓内加工单;CGTH=采购退货出库单)
	CancelReason  string   `xml:"cancelReason"`  //	已经退货	取消原因
}

func (t *TaoBaoQiMenOrderCancel) Check() (interface_factory.Response, error) {
	return nil, nil
}

func (t *TaoBaoQiMenOrderCancel) ToXML() string {
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQiMenOrderCancel) GetMethod() string {
	return method
}
