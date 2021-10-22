package inventoryquery

import (
	"encoding/xml"
	"fmt"
	"testing"
)

func TestTaoBaoQiMenInventoryQuery_ToXML(t *testing.T) {
	request:=&TaoBaoQiMenInventoryQuery{
		XMLName:      xml.Name{},
		Version:      "1.0",
		CriteriaList: CriteriaList{},
		ExtendProps:  ExtendProps{},
	}
	criteria:=Criteria{
		WarehouseCode: "",
		OwnerCode:     "",
		ItemCode:      "",
		ItemId:        "",
		InventoryType: "",
		Remark:        "",
	}
	criteriaList :=CriteriaList{}
	criteriaList.SetCriteria(criteria)
	request.SetCriteriaList(criteriaList)
	fmt.Printf("%s\n",request.ToXML())
}
