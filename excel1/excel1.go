package excel1

import (
	"github.com/xuri/excelize/v2"
	"log"
	"strings"
	"time"
)

type excel1 struct {
	name1     string
	sex       string
	age       string
	phone     string
	startDate time.Time
	endDate   time.Time
	content   string
	timeLong  string
	state     string
	name2     string
	idCard    string
	formTo    string
	level     string
}

type excel3 struct {
	xuhao     string
	canbaoren string
	idCard    string
	formTo    string
	level     string
	huliyuan  huliyuan
}

type huliyuan struct {
	name      string
	tianshu   int
	huliyuans map[string]huliyuan2
}

type huliyuan2 struct {
	num1 int
	day1 string
	num2 int
	day2 string
	num3 int
	day3 string
	num4 int
	day4 string
	num5 int
	day5 string
	num6 int
	day6 string
}

func GetStockList1(file *excelize.File) ([]excel1, error) {

	var list []excel1

	for _, sheetName := range file.GetSheetList() {

		rows, err := file.GetRows(sheetName)

		if err != nil {
			log.Printf("excel get rows error, err:%+v\n", err)
			return nil, err
		}

		for index, row := range rows {
			// 跳过第一行标题
			if index == 0 {
				continue
			}
			local, _ := time.LoadLocation("Asia/Shanghai")
			startDate, err := time.ParseInLocation("2006-01-02 15:04:05", strings.TrimSpace(row[4]), local)
			if err != nil {
				return nil, err
			}
			endDate, err := time.ParseInLocation("2006-01-02 15:04:05", strings.TrimSpace(row[5]), local)
			if err != nil {
				return nil, err
			}

			list = append(list, excel1{
				name1:     row[0],
				sex:       row[1],
				age:       row[2],
				phone:     row[3],
				startDate: startDate,
				endDate:   endDate,
				content:   row[6],
				timeLong:  row[7],
				state:     row[8],
				name2:     row[9],
				idCard:    row[10],
				formTo:    row[11],
				level:     row[12],
			})
		}
	}

	return list, nil
}

func Summary(list *[]excel1) []excel3 {
	return nil
}

func TheyRe(e excel1) bool {
	start := e.startDate.Day()
	end := e.endDate.Day()
	if end > start {
		return true
	}
	return false
}

// 计算时长
func Long(e excel1) map[float64]int {
	m := make(map[float64]int, 1)
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && e.level == "长护3级" {
		m[3] = 30
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && e.level == "长护2级" {
		m[2.5] = 12
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && e.level == "长护1级" {
		m[2] = 4
		return m
	}
	if e.formTo == "居民医保" && e.level == "长护3级" {
		m[1.5] = 30
		return m
	}
	if e.formTo == "居民医保" && e.level == "长护2级" {
		m[2.5] = 12
		return m
	}
	if e.formTo == "居民医保" && e.level == "长护1级" {
		m[2.5] = 2
		return m
	}
	if e.formTo == "居民医保" && e.level == "长护1级" {
		m[2.5] = 2
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && (e.level == "老年人能力等级5级") {
		m[3] = 30
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && (e.level == "老年人能力等级4级") {
		m[3] = 30
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && (e.level == "老年人能力等级3级") {
		m[2.5] = 12
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && (e.level == "老年人能力等级2级") {
		m[2] = 4
		return m
	}
	if (e.formTo == "市直医保" || e.formTo == "职工医保" || e.formTo == "省直医保") && (e.level == "老年人能力等级1级") {
		m[2] = 4
		return m
	}
	if e.formTo == "居民医保" && e.level == "老年人能力等级5级" {
		m[1.5] = 30
		return m
	}
	if e.formTo == "居民医保" && e.level == "老年人能力等级4级" {
		m[1.5] = 30
		return m
	}
	if e.formTo == "居民医保" && e.level == "老年人能力等级3级" {
		m[2.5] = 12
		return m
	}
	if e.formTo == "居民医保" && e.level == "老年人能力等级2级" {
		m[2.5] = 2
		return m
	}
	if e.formTo == "居民医保" && e.level == "老年人能力等级1级" {
		m[2.5] = 2
		return m
	}
	return nil
}
