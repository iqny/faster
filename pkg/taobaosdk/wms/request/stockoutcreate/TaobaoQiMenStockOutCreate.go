package stockoutcreate

//出库创建
import (
	"encoding/xml"
	"fmt"
	"orp/pkg/taobaosdk/wms/response"
	"strings"
)

const method = "taobao.qimen.stockout.create"

type TaoBaoQiMenStockOutCreate struct {
	XMLName       xml.Name      `xml:"request"`
	Version       string        `xml:"version,attr"`
	DeliveryOrder DeliveryOrder `xml:"deliveryOrder"`
	OrderLines    OrderLines    `xml:"orderLines"`
}

func (t *TaoBaoQiMenStockOutCreate) Check() (response.Response, error) {
	return nil, nil
}

func (t *TaoBaoQiMenStockOutCreate) ToXML() string {
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQiMenStockOutCreate) GetMethod() string {
	return method
}
func (t *TaoBaoQiMenStockOutCreate) SetDeliveryOrder(deliveryOrder DeliveryOrder) {
	t.DeliveryOrder = deliveryOrder
}
func (t *TaoBaoQiMenStockOutCreate) SetOrderLines(orderLines OrderLines) {
	t.OrderLines = orderLines
}

type OrderLines struct {
	OrderLine []OrderLine `xml:"orderLine"`
}

func (d *DeliveryOrder) SetSenderInfo(senderInfo SenderInfo) {
	d.SenderInfo = senderInfo
}
func (d *DeliveryOrder) SetReceiverInfo(receiverInfo ReceiverInfo) {
	d.ReceiverInfo = receiverInfo
}
func (d *DeliveryOrder) SetExtendProps(extendProps ExtendProps) {
	d.ExtendProps = extendProps
}
func (o *OrderLines) Append(orderLine OrderLine) {
	o.OrderLine = append(o.OrderLine, orderLine)
}

type DeliveryOrder struct {
	TotalOrderLines      int          `xml:"totalOrderLines"`      //	12	单据总行数(当单据需要分多个请求发送时;发送方需要将totalOrderLines填入;接收方收到后;根据实际接收到的条数和totalOrderLines进行比对;如果小于;则继续等待接收请求。如果等于;则表示该单据的所有请求发送完成.)
	DeliveryOrderCode    string       `xml:"deliveryOrderCode"`    //	TB1234	出库单号
	OrderType            string       `xml:"orderType"`            //	PTCK	出库单类型(PTCK=普通出库单;DBCK=调拨出库;B2BCK=B2B出库;QTCK=其他出库;CGTH=采购退货出库单;XNCK=虚拟出库单, JITCK=唯品出库)
	RelatedOrders        string       `xml:"relatedOrders"`        //		关联单据信息
	WarehouseCode        string       `xml:"warehouseCode"`        //	CK1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	CreateTime           string       `xml:"createTime"`           //	2016-09-09 12:00:00	出库单创建时间(YYYY-MM-DD HH:MM:SS)
	ScheduleDate         string       `xml:"scheduleDate"`         //	2017-09-09	要求出库时间(YYYY-MM-DD)
	LogisticsCode        string       `xml:"logisticsCode"`        //	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中通(ZTO)、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsName        string       `xml:"logisticsName"`        //	顺丰	物流公司名称(包括干线物流公司等)
	SupplierCode         string       `xml:"supplierCode"`         //	TB	供应商编码
	SupplierName         string       `xml:"supplierName"`         //	淘宝	供应商名称
	TransportMode        string       `xml:"transportMode"`        //	自提	提货方式(到仓自提、快递、干线物流)
	PickerInfo           PickerInfo   `xml:"pickerInfo"`           //		提货人信息
	SenderInfo           SenderInfo   `xml:"senderInfo"`           //		发件人信息
	ReceiverInfo         ReceiverInfo `xml:"receiverInfo"`         //		收件人信息
	Remark               string       `xml:"remark"`               //	备注信息	备注
	OrderSourceType      string       `xml:"orderSourceType"`      //	VIP	出库单渠道类型,VIP=唯品会，FX=分销 ，SHOP=门店
	ReceivingTime        string       `xml:"receivingTime"`        //	2016-09-09 12:00:00	到货时间(YYYY-MM-DD HH:MM:SS)
	ShippingTime         string       `xml:"shippingTime"`         //	2016-09-09 12:00:00	送货时间(YYYY-MM-DD HH:MM:SS)
	TargetWarehouseName  string       `xml:"targetWarehouseName"`  //	入库仓库名称, string (50)	入库仓库名称, string (50)
	TargetWarehouseCode  string       `xml:"targetWarehouseCode"`  //	入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER	入库仓库编码, string (50) ，统仓统配等无需ERP指定仓储编码的情况填OTHER
	TargetEntryOrderCode string       `xml:"targetEntryOrderCode"` //	关联的入库单号（ERP分配）	关联的入库单号（ERP分配）
	WarehouseName        string       `xml:"warehouseName"`        //	123	仓库名称
	ExtendProps          ExtendProps  `xml:"extendProps"`          //		扩展属性
}
type OrderLine struct {
	OutBizCode         string `xml:"outBizCode"`         //外部业务编码(消息ID;用于去重;当单据需要分批次发送时使用)
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
	Id            string `xml:"id"`            //	id 收件人地址ID
}
type PickerInfo struct {
	Company string `xml:"company"` //	天猫	公司名称
	Name    string `xml:"name"`    //	老王	姓名
	Tel     string `xml:"tel"`     //	897765	固定电话
	Mobile  string `xml:"mobile"`  //	123421234	移动电话
	Id      string `xml:"id"`      //	1232344322	证件号
	CarNo   string `xml:"carNo"`   //	XA1234	车牌号
}
type ExtendProps struct {
}
