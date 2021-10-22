package deliveryordercreate

//发货创建
import (
	"encoding/xml"
	"errors"
	"fmt"
	"orp/pkg/taobaosdk/wms/response"
	"strings"
)

const method = "taobao.qimen.deliveryorder.create"

type TaoBaoQiMenDeliveryOrderCreate struct {
	XMLName       xml.Name      `xml:"request"`
	Version       string        `xml:"version,attr"`
	DeliveryOrder DeliveryOrder `xml:"deliveryOrder"`
	OrderLines    OrderLines    `xml:"orderLines"`
}

func (t *TaoBaoQiMenDeliveryOrderCreate) ToXML() string {

	return ""
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQiMenDeliveryOrderCreate) GetMethod() string {
	return method
}
func (t *TaoBaoQiMenDeliveryOrderCreate) Check() (response.Response, error) {
	/*if res, err := t.DeliveryOrder.SenderInfo.check(); err != nil {
		return res, err
	}*/
	return nil, nil
}
func (t *TaoBaoQiMenDeliveryOrderCreate) SetDeliveryOrder(deliveryOrder DeliveryOrder) {
	t.DeliveryOrder = deliveryOrder
}
func (t *TaoBaoQiMenDeliveryOrderCreate) SetOrderLines(orderLines OrderLines) {
	t.OrderLines = orderLines
}

type DeliveryOrder struct {
	DeliveryOrderCode    string       `xml:"deliveryOrderCode"`    //	TB1234	出库单号
	PreDeliveryOrderCode string       `xml:"preDeliveryOrderCode"` //	Old1234	原出库单号(ERP分配)
	PreDeliveryOrderId   string       `xml:"preDeliveryOrderId"`   //	Oragin1234	原出库单号(WMS分配)
	OrderType            string       `xml:"orderType"`            //	JYCK	出库单类型(JYCK=一般交易出库单;HHCK=换货出库单;BFCK=补发出库单;QTCK=其他出 库 单)
	WarehouseCode        string       `xml:"warehouseCode"`        //	OTHER	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OrderFlag            string       `xml:"orderFlag"`            //	COD	订单标记(用字符串格式来表示订单标记列表:例如COD=货到付款;LIMIT=限时配 送;PRESELL=预 售;COMPLAIN=已投诉;SPLIT=拆单;EXCHANGE=换货;VISIT=上门;MODIFYTRANSPORT=是否 可改配送方式;CONSIGN = 物流宝代理发货;SELLER_AFFORD=是否卖家承担运费;FENXIAO=分销订 单)
	SourcePlatformCode   string       `xml:"sourcePlatformCode"`   //	TB	订单来源平台编码(TB=淘宝、TM=天猫、JD=京东、DD=当当、PP=拍拍、YX= 易讯、 EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂 、MGJ=蘑菇街、 JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿 里巴巴、POS=POS门店、 MIA=蜜芽、OTHER=其他(只传英文编码))
	SourcePlatformName   string       `xml:"sourcePlatformName"`   //	淘宝	订单来源平台名称
	CreateTime           string       `xml:"createTime"`           //	2016-07-06 12:00:00	发货单创建时间(YYYY-MM-DD HH:MM:SS)
	PlaceOrderTime       string       `xml:"placeOrderTime"`       //	2016-07-06 12:00:00	前台订单/店铺订单的创建时间/下单时间
	PayTime              string       `xml:"payTime"`              //	2016-07-06 12:00:00	订单支付时间(YYYY-MM-DD HH:MM:SS)
	PayNo                string       `xml:"payNo"`                //	P1234	支付平台交易号
	OperatorCode         string       `xml:"operatorCode"`         //	0123	操作员(审核员)编码
	OperatorName         string       `xml:"operatorName"`         //	老王	操作员(审核员)名称
	OperateTime          string       `xml:"operateTime"`          //	2016-07-06 12:00:00	操作(审核)时间(YYYY-MM-DD HH:MM:SS)
	ShopNick             string       `xml:"shopNick"`             //	淘宝店	店铺名称
	SellerNick           string       `xml:"sellerNick"`           //	淘宝店主	卖家名称
	BuyerNick            string       `xml:"buyerNick"`            //	淘公仔	买家昵称
	TotalAmount          string       `xml:"totalAmount"`          //	123	订单总金额(订单总金额=应收金额+已收金额=商品总金额-订单折扣金额+快递费用 ;单位 元)
	ItemAmount           string       `xml:"itemAmount"`           //	234	商品总金额(元)
	DiscountAmount       string       `xml:"discountAmount"`       //	111	订单折扣金额(元)
	Freight              string       `xml:"freight"`              //	111	快递费用(元)
	ArAmount             string       `xml:"arAmount"`             //	111	应收金额(消费者还需要支付多少--货到付款时消费者还需要支付多少约定使用这个字 段;单位元 )
	GotAmount            string       `xml:"gotAmount"`            //	111	已收金额(消费者已经支付多少;单位元)
	ServiceFee           string       `xml:"serviceFee"`           //	111	COD服务费
	LogisticsCode        string       `xml:"logisticsCode"`        //	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通 、ZTO=中 通(ZTO)、HTKY=百世汇通、UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷 、POSTB=邮政小包、 GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、AMAZON=亚马逊物流、 OTHER=其他)
	LogisticsName        string       `xml:"logisticsName"`        //	顺丰	物流公司名称
	ExpressCode          string       `xml:"expressCode"`          //	Y12345	运单号
	LogisticsAreaCode    string       `xml:"logisticsAreaCode"`    //	043300	快递区域编码
	SenderInfo           SenderInfo   `xml:"senderInfo"`           // 	发件人信息
	ReceiverInfo         ReceiverInfo `xml:"receiverInfo"`         //	收件人信息
	Invoices             Invoices     `xml:"invoices"`             //发票信息
	Insurance            Insurance    `xml:"insurance"`            //保险信息
	ExtendProps          ExtendProps  `xml:"extendProps"`          //扩展字段
}

func (d *DeliveryOrder) SetSenderInfo(senderInfo SenderInfo) {
	d.SenderInfo = senderInfo
}
func (d *DeliveryOrder) SetReceiverInfo(receiverInfo ReceiverInfo) {
	d.ReceiverInfo = receiverInfo
}
func (d *DeliveryOrder) SetInvoices(invoices Invoices) {
	d.Invoices = invoices
}
func (d *DeliveryOrder) SetInsurance(insurance Insurance) {
	d.Insurance = insurance
}
func (d *DeliveryOrder) SetExtendProps(extendProps ExtendProps) {
	d.ExtendProps = extendProps
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

// Invoices 发票信息
type Invoices struct {
	Invoice []Invoice `xml:"invoice"`
}

func (i *Invoices) Append(invoice Invoice) {
	i.Invoice = append(i.Invoice, invoice)
}

type Invoice struct {
	Type      string `xml:"type"`      //	INVOICE	发票类型(INVOICE=普通发票;VINVOICE=增值税普通发票;EVINVOICE=电子增票;填写的 条件 是:invoiceFlag为Y)
	Header    string `xml:"header"`    //	天猫科技有限公司	发票抬头(填写的条件是:invoiceFlag为Y)
	Amount    string `xml:"amount"`    //	100	发票总金额(填写的条件是:invoiceFlag为Y)
	Content   string `xml:"content"`   //	增值税100元	发票内容(不推荐使用)
	Detail    Detail `xml:"detail"`    //		当content和detail同时存在时，优先处理detail的信息
	TaxNumber string `xml:"taxNumber"` //	123	税号
}

type Detail struct {
	Item []Item `xml:"item"`
}

func (d *Detail) Append(item Item) {
	d.Item = append(d.Item, item)
}

type Item struct {
	Sku string `xml:"sku"`
}

// Insurance 保险信息
type Insurance struct {
	Type   string `xml:"type"`   //消费险	保险类型
	Amount string `xml:"amount"` //1	保险金额
}
type ExtendProps struct {
}

func (s *SenderInfo) check() (response.Response, error) {
	var err = errors.New("check fail")
	var res response.ErrResponse
	if s.Name == "" {
		res.Code = "340"
		return res, err
	}
	return res, nil
}
