package excel1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"testing"
)

func TestName(t *testing.T) {
	file, _ := excelize.OpenFile("../10月打卡汇总.xlsx")
	list1, _ := GetStockList1(file)
	m := make(map[string]excel3, 0)
	for index, _ := range list1 {
		e := excel3{
			canbaoren: list1[index].name2,
			idCard:    list1[index].idCard,
			formTo:    list1[index].formTo,
			level:     list1[index].level,
		}
		m[list1[index].name2] = e
	}
	for k, v := range m {
		fmt.Printf("%s,%v \n", k, v)
	}

}

/*func TestName1(t *testing.T) {
	m := make(map[string]excel3, 1)
	m[]
}
*/
