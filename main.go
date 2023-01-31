package main

import (
	"code/excel1"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	file1, _ := excelize.OpenFile("10月打卡汇总.xlsx")
	list, _ := excel1.GetStockList1(file1)
	fmt.Printf("%v", list[0])
	m := excel1.Long(list[0])
	fmt.Printf("%v", m)
	/*for index, value := range list {
		fmt.Printf("%d %v ln", index, value)
	}*/

}
