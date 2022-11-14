package main

import (
	"code/excel1"
	"fmt"
	"github.com/xuri/excelize/v2"
)

func main() {
	file1, _ := excelize.OpenFile("10月打卡汇总.xlsx")

	list, _ := excel1.GetStockList1(file1)
	fmt.Printf("%v", list[2])

}
