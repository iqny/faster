package entryordercreate

//入库创建
import (
	"encoding/xml"
	"fmt"
	"orp/pkg/wms/interface_factory"
	"strings"
)

const method = "taobao.qimen.entryorder.create"

type TaoBaoQiMenEntryOrderCreate struct {
	XMLName     xml.Name    `xml:"request"`
	Version     string      `xml:"version,attr"`
	EntryOrder  EntryOrder  `xml:"entryOrder"`
	OrderLines  OrderLines  `xml:"orderLines"`
	ExtendProps ExtendProps `xml:"extendProps"` //扩展字段
}

func (t *TaoBaoQiMenEntryOrderCreate) Check() (interface_factory.Response, error) {
	return nil, nil
}

func (t *TaoBaoQiMenEntryOrderCreate) ToXML() string {
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQiMenEntryOrderCreate) GetMethod() string {
	return method
}

func (t *TaoBaoQiMenEntryOrderCreate) SetDeliveryOrder(entryOrder EntryOrder) {
	t.EntryOrder = entryOrder
}
func (t *TaoBaoQiMenEntryOrderCreate) SetOrderLines(orderLines OrderLines) {
	t.OrderLines = orderLines
}
func (t *TaoBaoQiMenEntryOrderCreate) SetExtendProps(extendProps ExtendProps) {
	t.ExtendProps = extendProps
}

type EntryOrder struct {
	EntryOrder          string        `xml:"entryOrder"`          //		入库单信息
	EntryOrderCode      string        `xml:"entryOrderCode"`      //	E1234	入库单号
	OwnerCode           string        `xml:"ownerCode"`           //	O1234	货主编码
	PurchaseOrderCode   string        `xml:"purchaseOrderCode"`   //	C123455	采购单号(当orderType=CGRK时使用)
	WarehouseCode       string        `xml:"warehouseCode"`       //	W1234	入库仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OrderCreateTime     string        `xml:"orderCreateTime"`     //	2016-09-09 12:00:00	订单创建时间(YYYY-MM-DD HH:MM:SS)
	OrderType           string        `xml:"orderType"`           //	SCRK	业务类型(SCRK=生产入库;LYRK=领用入库;CCRK=残次品入库;CGRK=采购入库;DBRK=调拨入库;QTRK=其他入库;B2BRK=B2B入 库;XNRK=虚拟入库;只传英文编码)
	RelatedOrders       RelatedOrders `xml:"relatedOrders"`       //关联订单信息
	ExpectStartTime     string        `xml:"expectStartTime"`     //	2015-09-09 12:00:00	预期到货时间(YYYY-MM-DD HH:MM:SS)
	ExpectEndTime       string        `xml:"expectEndTime"`       //	2015-09-09 12:00:00	最迟预期到货时间(YYYY-MM-DD HH:MM:SS)
	LogisticsCode       string        `xml:"logisticsCode"`       //	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中通(ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他(只传英文编码))
	LogisticsName       string        `xml:"logisticsName"`       //	顺丰	物流公司名称
	ExpressCode         string        `xml:"expressCode"`         //	Y1234	运单号
	SupplierCode        string        `xml:"supplierCode"`        //	G1234	供应商编码
	SupplierName        string        `xml:"supplierName"`        //	GN1234	供应商名称
	OperatorCode        string        `xml:"operatorCode"`        //	ON1234	操作员编码
	OperatorName        string        `xml:"operatorName"`        //	老王	操作员名称
	OperateTime         string        `xml:"operateTime"`         //	2017-09-09 12:00:00	操作时间(YYYY-MM-DD HH:MM:SS)
	SenderInfo          SenderInfo    `xml:"senderInfo"`          // 	发件人信息
	ReceiverInfo        ReceiverInfo  `xml:"receiverInfo"`        //	收件人信息
	Remark              string        `xml:"remark"`              //	备注信息	备注
	TotalOrderLines     string        `xml:"totalOrderLines"`     //	12	totalOrderLines
	WarehouseName       string        `xml:"warehouseName"`       //	E1234	入库仓库名称
	SourceWarehouseCode string        `xml:"sourceWarehouseCode"` //	E1234	出库仓库编码
	SourceWarehouseName string        `xml:"sourceWarehouseName"` //	E1234	出库仓库名称
}

func (d *EntryOrder) SetSenderInfo(senderInfo SenderInfo) {
	d.SenderInfo = senderInfo
}
func (d *EntryOrder) SetReceiverInfo(receiverInfo ReceiverInfo) {
	d.ReceiverInfo = receiverInfo
}

type OrderLines struct {
	OrderLine []OrderLine `xml:"orderLine"`
}

func (o *OrderLines) Append(orderLine OrderLine) {
	o.OrderLine = append(o.OrderLine, orderLine)
}

type OrderLine struct {
	OrderLineNo        string `xml:"orderLineNo"`        //	11	单据行号
	SourceOrderCode    string `xml:"sourceOrderCode"`    //	S12345	交易平台订单
	SubSourceOrderCode string `xml:"subSourceOrderCode"` //	S1234	交易平台子订单编码
	PayNo              string `xml:"payNo"`              //	J1234	支付平台交易号(淘系订单传支付宝交易号)
	OwnerCode          string `xml:"ownerCode"`          //	H1234	货主编码
	ItemCode           string `xml:"itemCode"`           //	I1234	商品编码
	ItemId             string `xml:"itemId"`             //	W1234	仓储系统商品编码
	InventoryType      string `xml:"inventoryType"`      //	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;ZT=在途库存;默认为查所有类型的库存)
	ItemName           string `xml:"itemName"`           //	淘公仔	商品名称
	ExtCode            string `xml:"extCode"`            //	J1234	交易平台商品编码
	PlanQty            int    `xml:"planQty"`            //	11	应发商品数量
	RetailPrice        string `xml:"retailPrice"`        //	12.0	零售价(零售价=实际成交价+单件商品折扣金额)
	ActualPrice        string `xml:"actualPrice"`        //	12.0	实际成交价
	DiscountAmount     string `xml:"discountAmount"`     //	0	单件商品折扣金额
	BatchCode          string `xml:"batchCode"`          //	123	批次编码
	ProductDate        string `xml:"productDate"`        //	2016-07-06	生产日期(YYYY-MM-DD)
	ExpireDate         string `xml:"expireDate"`         //	2016-07-06	过期日期(YYYY-MM-DD)
	ProduceCode        string `xml:"produceCode"`        //	123	生产批号
	SnList             SnList `xml:"snList"`             //		sn编码列表
}
type SenderInfo struct {
	Company       string `xml:"company"`       //	淘宝	公司名称
	Name          string `xml:"name"`          //    老王    姓名
	ZipCode       string `xml:"zipCode"`       //	043300	邮编
	Tel           string `xml:"tel"`           //	81020340	固定电话
	Mobile        string `xml:"mobile"`        //   13214567869    移动电话
	Email         string `xml:"email"`         //	345@gmail.com	电子邮箱
	CountryCode   string `xml:"countryCode"`   //	051532	国家二字码
	Province      string `xml:"province"`      //   浙江省    省份
	City          string `xml:"city"`          // 	杭州    城市
	Area          string `xml:"area"`          //	余杭	区域
	Town          string `xml:"town"`          //	横加桥	村镇
	DetailAddress string `xml:"detailAddress"` //    杭州市余杭区989号    详细地址
}
type ReceiverInfo struct {
	Company       string `xml:"company"`       //	淘宝	公司名称
	Name          string `xml:"name"`          //	老王	姓名
	ZipCode       string `xml:"zipCode"`       //	043300	邮编
	Tel           string `xml:"tel"`           //	808786543	固定电话
	Mobile        string `xml:"mobile"`        //	13423456785	移动电话
	IdType        string `xml:"idType"`        //	1 收件人证件类型(1-身份证、2-军官证、3-护照、4-其他)
	IdNumber      string `xml:"idNumber"`      //	12345678	收件人证件号码
	Email         string `xml:"email"`         //	878987654@qq.com	电子邮箱
	CountryCode   string `xml:"countryCode"`   //	045565	国家二字码
	Province      string `xml:"province"`      //	浙江省	省份
	City          string `xml:"city"`          //	杭州	城市
	Area          string `xml:"area"`          //	余杭	区域
	Town          string `xml:"town"`          //	横加桥	村镇
	DetailAddress string `xml:"detailAddress"` //	杭州市余杭区989号	详细地址
	Oaid          string `xml:"oaid"`          //	oaid 收件人地址ID
}
type RelatedOrders struct {
	RelatedOrder []RelatedOrder `xml:"relatedOrder"`
}

func (r *RelatedOrders) Append(relatedOrder RelatedOrder) {
	r.RelatedOrder = append(r.RelatedOrder, relatedOrder)
}

type RelatedOrder struct {
	OrderType string `xml:"orderType"` //	CG	关联的订单类型(CG=采购单;DB=调拨单;CK=出库单;RK=入库单;只传英文编码)
	OrderCode string `xml:"orderCode"` //
}
type SnList struct {
	Sn []string `xml:"sn"`
}
type ExtendProps struct {
}
