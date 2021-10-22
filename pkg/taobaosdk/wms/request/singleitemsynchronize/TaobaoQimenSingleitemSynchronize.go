package singleitemsynchronize

//商品同步
import (
	"encoding/xml"
	"fmt"
	"orp/pkg/taobaosdk/wms/response"
	"strings"
)

const method = "taobao.qimen.singleitem.synchronize"

type TaoBaoQimEnSingleItemSynchronize struct {
	XMLName       xml.Name    `xml:"request"`
	Version       string      `xml:"version,attr"`
	ActionType    string      `xml:"actionType"`    //	add	操作类型(两种类型：add|update)
	WarehouseCode string      `xml:"warehouseCode"` //	CK1234	仓库编码(统仓统配等无需ERP指定仓储编码的情况填OTHER)
	OwnerCode     string      `xml:"ownerCode"`     //	HZ123	货主编码
	SupplierCode  string      `xml:"supplierCode"`  //	GY123	供应商编码
	SupplierName  string      `xml:"supplierName"`  //	淘宝	供应商名称
	Item          Item        `xml:"item"`          //商品信息
	ExtendProps   ExtendProps `xml:"extendProps"`   //扩展字段
}

func (t *TaoBaoQimEnSingleItemSynchronize) Check() (response.Response, error) {
	return nil, nil
}

func (t *TaoBaoQimEnSingleItemSynchronize) ToXML() string {
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQimEnSingleItemSynchronize) GetMethod() string {
	return method
}
func (t *TaoBaoQimEnSingleItemSynchronize) SetItem(item Item) {
	t.Item = item
}
func (t *TaoBaoQimEnSingleItemSynchronize) SetExtendProps(extendProps ExtendProps) {
	t.ExtendProps = extendProps
}

// 商品信息
type Item struct {
	ItemCode        string `xml:"itemCode"`        //	I1234	商品编码
	ItemId          string `xml:"itemId"`          //	WI1234	仓储系统商品编码(该字段是WMS分配的商品编号;WMS如果分配了商品编码;则后续的商品操作都需要传该字段;如果WMS不使用;WMS可 以返回itemId=itemCode的值)
	GoodsCode       string `xml:"goodsCode"`       //	H1234	货号
	ItemName        string `xml:"itemName"`        //	SN123	商品名称
	ShortName       string `xml:"shortName"`       //	JC123	商品简称
	EnglishName     string `xml:"englishName"`     //	EN123	英文名
	BarCode         string `xml:"barCode"`         //	T1;T2	条形码(可多个;用分号;隔开)
	SkuProperty     string `xml:"skuProperty"`     //	红色	商品属性(如红色;XXL)
	StockUnit       string `xml:"stockUnit"`       //	个	商品计量单位
	Length          string `xml:"length"`          //	12.0	长(单位：厘米)
	Width           string `xml:"width"`           //	12.0	宽(单位：厘米)
	Height          string `xml:"height"`          //	12.0	高(单位：厘米)
	Volume          string `xml:"volume"`          //	12.0	体积(单位：升)
	GrossWeight     string `xml:"grossWeight"`     //	12.0	毛重(单位：千克)
	NetWeight       string `xml:"netWeight"`       //	12.0	净重(单位：千克)
	Color           string `xml:"color"`           //	红色	颜色
	Size            string `xml:"size"`            //	5英尺	尺寸
	Title           string `xml:"title"`           //	淘公仔	渠道中的商品标题
	CategoryId      string `xml:"categoryId"`      //	LB123	商品类别ID
	CategoryName    string `xml:"categoryName"`    //	手机	商品类别名称
	PricingCategory string `xml:"pricingCategory"` //	手机类	计价货类
	SafetyStock     int    `xml:"safetyStock"`     //	12	安全库存
	ItemType        string `xml:"itemType"`        //	ZC	商品类型(ZC=正常商品;FX=分销商品;ZH=组合商品;ZP=赠品;BC=包材;HC=耗材;FL=辅料;XN=虚拟品;FS=附属品;CC=残次品; OTHER=其它;只传英文编码)
	TagPrice        string `xml:"tagPrice"`        //	12.0	吊牌价
	RetailPrice     string `xml:"retailPrice"`     //	12.0	零售价
	CostPrice       string `xml:"costPrice"`       //	12.0	成本价
	PurchasePrice   string `xml:"purchasePrice"`   //	12.0	采购价
	SeasonCode      string `xml:"seasonCode"`      //	CHUN	季节编码
	SeasonName      string `xml:"seasonName"`      //	春季	季节名称
	BrandCode       string `xml:"brandCode"`       //	LAL	品牌代码
	BrandName       string `xml:"brandName"`       //	HM	品牌名称
	IsSNMgmt        string `xml:"isSNMgmt"`        //	N	是否需要串号管理(Y/N ;默认为N)
	ProductDate     string `xml:"productDate"`     //	2016-09-09	生产日期(YYYY-MM-DD)
	ExpireDate      string `xml:"expireDate"`      //	2016-09-09	过期日期(YYYY-MM-DD)
	IsShelfLifeMgmt string `xml:"isShelfLifeMgmt"` //	N	是否需要保质期管理(Y/N ;默认为N)
	ShelfLife       int    `xml:"shelfLife"`       //	1	保质期(单位：小时)
	RejectLifecycle int    `xml:"rejectLifecycle"` //	1	保质期禁收天数
	LockupLifecycle int    `xml:"lockupLifecycle"` //	1	保质期禁售天数
	AdventLifecycle int    `xml:"adventLifecycle"` //	1	保质期临期预警天数
	IsBatchMgmt     string `xml:"isBatchMgmt"`     //	N	是否需要批次管理(Y/N ;默认为N)
	BatchCode       string `xml:"batchCode"`       //	P1234	批次代码
	BatchRemark     string `xml:"batchRemark"`     //	备注信息	批次备注
	PackCode        string `xml:"packCode"`        //	B12	包装代码
	Pcs             string `xml:"pcs"`             //	XG123	箱规
	OriginAddress   string `xml:"originAddress"`   //	HK	商品的原产地
	ApprovalNumber  string `xml:"approvalNumber"`  //	PB123	批准文号
	IsFragile       string `xml:"isFragile"`       //	N	是否易碎品(Y/N ;默认为N)
	IsHazardous     string `xml:"isHazardous"`     //	N	是否危险品(Y/N ;默认为N)
	Remark          string `xml:"remark"`          //	备注信息	备注
	CreateTime      string `xml:"createTime"`      //	2017-09-09 12:00:00	创建时间(YYYY-MM-DD HH:MM:SS)
	UpdateTime      string `xml:"updateTime"`      //	2017-09-09 12:00:00	更新时间(YYYY-MM-DD HH:MM:SS)
	IsValid         string `xml:"isValid"`         //	N	是否有效(Y/N ;默认为N)
	IsSku           string `xml:"isSku"`           //	N	是否sku(Y/N ;默认为N)
	PackageMaterial string `xml:"packageMaterial"` //	BX123	商品包装材料类型
	SupplierCode    string `xml:"supplierCode"`    //	temp	temp
	LogisticsType   string `xml:"logisticsType"`   //	0	销售配送方式（0=自配|1=菜鸟）
	IsLiquid        string `xml:"isLiquid"`        //	Y	是否液体, Y/N, (默认为N)
}
type ExtendProps struct {
}
