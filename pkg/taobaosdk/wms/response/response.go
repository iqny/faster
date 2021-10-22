package response

import "encoding/xml"

type Response interface {
	GetCode() string
	GetContent() string
}
// SuccessResponse x := "<response><flag>success</flag><code>200</code><message>Success</message><data></data></response>"
type SuccessResponse struct {
	XMLName    xml.Name    `xml:"response"`
	Flag       string      `xml:"flag"`
	Code       string      `xml:"code"`
	Message    string      `xml:"message"`
	SubCode    string      `xml:"sub_code"`
	SubMessage string      `xml:"sub_message"`
	RequestId  string      `xml:"request_id"`
	Data       interface{} `xml:"data"`
	Items      Items       `xml:"items"`
	Res        string      `xml:"-"`
	Req        string      `xml:"-"`
}

func (r SuccessResponse) GetCode() string {
	return r.Code
}
func (r SuccessResponse) GetContent() string {
	return r.Res
}

type Items struct {
	Item []Item `xml:"item"`
}

// 商品的库存信息列表
type Item struct {
	WarehouseCode string `xml:"warehouseCode"` //C1234	仓库编码
	ItemCode      string `xml:"itemCode"`      //I1234	商品编码
	ItemId        string `xml:"itemId"`        //	W1234	仓储系统商品ID
	InventoryType string `xml:"inventoryType"` //	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS= 箱损;ZT=在途库存)
	Quantity      int    `xml:"quantity"`      //	11	未冻结库存数量
	LockQuantity  int    `xml:"lockQuantity"`  //	1	冻结库存数量
	BatchCode     string `xml:"batchCode"`     //	P1234	批次编码
	ProductDate   string `xml:"productDate"`   //	2017-09-09	商品生产日期(YYYY-MM-DD)
	ExpireDate    string `xml:"expireDate"`    //	2017-09-09	商品过期日期(YYYY-MM-DD)
	ProduceCode   string `xml:"produceCode"`   //	P1234	生产批号
}
// ErrResponse xres:="<error_response><code>50</code><msg>Remote service error</msg><sub_code>isv.invalid-parameter</sub_code><sub_message>非法参数</sub_message></error_response>"
type ErrResponse struct {
	XMLName    xml.Name `xml:"error_response"`
	Code       string   `xml:"code"`
	Message    string   `xml:"msg"`
	SubCode    string   `xml:"sub_code"`
	SubMessage string   `xml:"sub_message"`
	RequestId  string   `xml:"request_id"`
	Res        string   `xml:"-"`
	Req        string   `xml:"-"`
}

func (r ErrResponse) GetCode() string {
	return r.Code
}
func (r ErrResponse) GetContent() string {
	return r.Res
}
