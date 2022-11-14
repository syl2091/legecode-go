package excel2

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
	"time"
)

type excel2 struct {
	xuhao     string
	zhandian  string
	people    string
	idCard    string
	xingzhi   string
	daiyu     string
	startDate time.Time
	fuzheren  string
	name      string
	phone     string
	huliyuan  string
	phone2    string
}

func GetStockList2(file *excelize.File, sheetNames []string) ([]excel2, error) {

	var list []excel2

	for _, sheetName := range file.GetSheetList() {

		rows, err := file.GetRows(sheetName)

		if err != nil {
			log.Printf("excel get rows error, err:%+v\n", err)
			return nil, err
		}

		for index, row := range rows {
			// 跳过第一行标题
			if index < 4 {
				continue
			}
			local, _ := time.LoadLocation("Asia/Shanghai")
			startDate, err := time.ParseInLocation("2006-01-02", strings.TrimSpace(row[6]), local)
			if err != nil {
				return nil, err
			}

			list = append(list, excel2{
				xuhao:     row[0],
				zhandian:  row[1],
				people:    row[2],
				idCard:    row[3],
				xingzhi:   row[4],
				daiyu:     row[5],
				startDate: startDate,
				fuzheren:  row[7],
				name:      row[8],
				phone:     row[9],
				huliyuan:  row[10],
				phone2:    row[11],
			})
		}
	}

	return list, nil
}
