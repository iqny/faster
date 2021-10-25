package returnordercreate

import (
	"encoding/xml"
	"fmt"
	"orp/pkg/wms/interface_factory"
	"strings"
)

const method = "taobao.qimen.returnorder.create"

type TaoBaoQiMenReturnOrderCreate struct {
	XMLName     xml.Name    `xml:"request"`
	Version     string      `xml:"version,attr"`
	ReturnOrder ReturnOrder `xml:"entryOrder"`
	OrderLines  OrderLines  `xml:"orderLines"`
	ExtendProps ExtendProps `xml:"extendProps"` //扩展字段
}

func (t *TaoBaoQiMenReturnOrderCreate) Check() (interface_factory.Response, error) {
	return nil, nil
}

func (t *TaoBaoQiMenReturnOrderCreate) ToXML() string {
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQiMenReturnOrderCreate) GetMethod() string {
	return method
}
func (t *TaoBaoQiMenReturnOrderCreate) SetReturnOrder(returnOrder ReturnOrder) {
	t.ReturnOrder = returnOrder
}
func (t *TaoBaoQiMenReturnOrderCreate) SetOrderLines(orderLines OrderLines) {
	t.OrderLines = orderLines
}
func (t *TaoBaoQiMenReturnOrderCreate) SetExtendProps(extendProps ExtendProps) {
	t.ExtendProps = extendProps
}

type OrderLines struct {
	OrderLine []OrderLine `xml:"orderLine"`
}

func (o *OrderLines) Append(orderLine OrderLine) {
	o.OrderLine = append(o.OrderLine, orderLine)
}

type ReturnOrder struct {
	ReturnOrderCode      string     `xml:"returnOrderCode"`      //	R1234	ERP的退货入库单编码
	WarehouseCode        string     `xml:"warehouseCode"`        //	W1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OrderType            string     `xml:"orderType"`            //	THRK	单据类型(THRK=退货入库;HHRK=换货入库;只传英文编码)
	OrderFlag            string     `xml:"orderFlag"`            //	VISIT	用字符串格式来表示订单标记列表(比如VISIT^ SELLER_AFFORD^SYNC_RETURN_BILL等;中间用“^”来隔开订单标记list (所 有字母全部大写) VISIT=上门；SELLER_AFFORD=是否卖家承担运费(默认是)SYNC_RETURN_BILL=同时退回发票)
	PreDeliveryOrderCode string     `xml:"preDeliveryOrderCode"` //	PD1234	原出库单号(ERP分配)
	PreDeliveryOrderId   string     `xml:"preDeliveryOrderId"`   //	PDI1234	原出库单号(WMS分配)
	LogisticsCode        string     `xml:"logisticsCode"`        //	SF	物流公司编码(SF=顺丰、EMS=标准快递、EYB=经济快件、ZJS=宅急送、YTO=圆通、ZTO=中通(ZTO)、HTKY=百世汇通、 UC=优速、STO=申通、TTKDEX=天天快递、QFKD=全峰、FAST=快捷、POSTB=邮政小包、GTO=国通、YUNDA=韵达、JD=京东配送、DD=当当宅配、 AMAZON=亚马逊物流、OTHER=其他;只传英文编码)
	LogisticsName        string     `xml:"logisticsName"`        //	顺丰	物流公司名称
	ExpressCode          string     `xml:"expressCode"`          //	YD1234	运单号
	ReturnReason         string     `xml:"returnReason"`         //	破损退货	退货原因
	BuyerNick            string     `xml:"buyerNick"`            //	淘宝	买家昵称
	SenderInfo           SenderInfo `xml:"senderInfo"`           //		发件人信息
	Remark               string     `xml:"remark"`               //	备注信息	备注
	SourcePlatformCode   string     `xml:"sourcePlatformCode"`   //	TB	订单来源平台编码, string (50),TB= 淘宝 、TM=天猫 、JD=京东、DD=当当、PP=拍拍、YX=易讯、EBAY=ebay、QQ=QQ网购、AMAZON=亚马逊、SN=苏宁、GM=国美、WPH=唯品会、JM=聚美、LF=乐蜂、MGJ=蘑菇街、JS=聚尚、PX=拍鞋、YT=银泰、YHD=1号店、VANCL=凡客、YL=邮乐、YG=优购、1688=阿里巴巴、POS=POS门店、MIA=蜜芽、GW=商家官网、CT=村淘、YJWD=云集微店、PDD=拼多多、OTHERS=其他,
	SourcePlatformName   string     `xml:"sourcePlatformName"`   //	淘宝	订单来源平台名称
	ShopNick             string     `xml:"shopNick"`             //	店铺名称	店铺名称
	SellerNick           string     `xml:"sellerNick"`           //	卖家名称	卖家名称
}

type OrderLine struct {
	OrderLineNo        string `xml:"orderLineNo"`        //	11	单据行号
	SourceOrderCode    string `xml:"sourceOrderCode"`    //	S12345	交易平台订单
	SubSourceOrderCode string `xml:"subSourceOrderCode"` //	S1234	交易平台子订单编码
	OwnerCode          string `xml:"ownerCode"`          //	H1234	货主编码
	ItemCode           string `xml:"itemCode"`           //	I1234	商品编码
	ItemId             string `xml:"itemId"`             //	W1234	仓储系统商品编码
	InventoryType      string `xml:"inventoryType"`      //	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;ZT=在途库存;默认为查所有类型的库存)
	PlanQty            int    `xml:"planQty"`            //	11	应发商品数量
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
type SnList struct {
	Sn []string `xml:"sn"`
}
type ExtendProps struct {
}
