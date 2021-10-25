package inventoryquery

import (
	"encoding/xml"
	"fmt"
	"orp/pkg/wms/interface_factory"
	"strings"
)

//库存查询接口(多商品)
const method = "taobao.qimen.inventory.query"

type TaoBaoQiMenInventoryQuery struct {
	XMLName      xml.Name     `xml:"request"`
	Version      string       `xml:"version,attr"`
	CriteriaList CriteriaList `xml:"criteriaList"`
	ExtendProps  ExtendProps  `xml:"extendProps"` //扩展字段
}

func (t *TaoBaoQiMenInventoryQuery) Check() (interface_factory.Response, error) {
	return nil, nil
}

func (t *TaoBaoQiMenInventoryQuery) ToXML() string {
	output, err := xml.MarshalIndent(t, "", "")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	return strings.Join([]string{xml.Header, string(output)}, "")
}
func (t *TaoBaoQiMenInventoryQuery) GetMethod() string {
	return method
}
func (t *TaoBaoQiMenInventoryQuery) SetCriteriaList(criteriaList CriteriaList) {
	t.CriteriaList = criteriaList
}
func (t *CriteriaList) SetCriteria(criteria Criteria) {
	t.Criteria = append(t.Criteria, criteria)
}

type CriteriaList struct {
	Criteria []Criteria `xml:"criteria"`
}
type Criteria struct {
	WarehouseCode string `xml:"warehouseCode"` //	W1234	仓库编码
	OwnerCode     string `xml:"ownerCode"`     //	H1234	货主编码
	ItemCode      string `xml:"itemCode"`      //	I1234	商品编码
	ItemId        string `xml:"itemId"`        //	C1234	仓储系统商品ID
	InventoryType string `xml:"inventoryType"` //	ZP	库存类型(ZP=正品;CC=残次;JS=机损;XS=箱损;ZT=在途库存;默认为查所有类型的库存)
	Remark        string `xml:"remark"`        //	备注	备注
}
type ExtendProps struct {
}
