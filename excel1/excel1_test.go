package excel1

import (
	"fmt"
	"github.com/xuri/excelize/v2"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	file, _ := excelize.OpenFile("../10月打卡汇总.xlsx")
	list1, _ := GetStockList1(file)
	h := new(huliyuan)
	h.huliyuans = make(map[string]huliyuan2, 0)
	m := make(map[string]excel3, 0)
	for index, _ := range list1 {
		if index == 0 {
			h = new(huliyuan)
			h.huliyuans = make(map[string]huliyuan2, 0)
		} else if list1[index].name1 != list1[index-1].name1 {
			h = new(huliyuan)
			h.huliyuans = make(map[string]huliyuan2, 0)
		} else if list1[index].name2 != list1[index-1].name2 {
			h = new(huliyuan)
			h.huliyuans = make(map[string]huliyuan2, 0)
		}
		e := excel3{
			canbaoren: strings.TrimSpace(list1[index].name2),
			idCard:    list1[index].idCard,
			formTo:    list1[index].formTo,
			level:     list1[index].level,
			huliyuan:  count(h, list1[index]),
		}
		m[list1[index].name2] = e
	}
	for k, v := range m {
		k = k
		for k1, v1 := range v.huliyuan.huliyuans {
			v.huliyuan.name += "\n" + k1
			v.huliyuan.tianshu += v1.num1
		}
		m[k] = v
	}
	/*fmt.Printf("%v \n", m["伍东成"])
		for k, _ := range m {
		if NumberOf(m[k].huliyuan.huliyuans) != 1 {
			fmt.Println(k)
		}
	}*/
	file2, _ := excelize.OpenFile("../10月汇总 - 1号调整(1)(1).xlsx")
	i := 59
	for k, _ := range m {
		file2.SetCellValue("长护记账", "B"+strconv.Itoa(i), k)
		file2.SetCellValue("长护记账", "C"+strconv.Itoa(i), m[k].idCard)
		file2.SetCellValue("长护记账", "D"+strconv.Itoa(i), m[k].formTo)
		file2.SetCellValue("长护记账", "E"+strconv.Itoa(i), m[k].level)
		file2.SetCellValue("长护记账", "K"+strconv.Itoa(i), m[k].huliyuan.name)
		file2.SetCellValue("长护记账", "L"+strconv.Itoa(i), m[k].huliyuan.tianshu)
		for k, v := range m[k].huliyuan.huliyuans {
			file2.SetCellValue("长护记账", "M"+strconv.Itoa(i), k)
			file2.SetCellValue("长护记账", "N"+strconv.Itoa(i), v.num1)
			file2.SetCellValue("长护记账", "O"+strconv.Itoa(i), v.day1)
			file2.SetCellValue("长护记账", "P"+strconv.Itoa(i), v.num2)
			file2.SetCellValue("长护记账", "Q"+strconv.Itoa(i), v.day2)
			file2.SetCellValue("长护记账", "R"+strconv.Itoa(i), v.num3)
			file2.SetCellValue("长护记账", "S"+strconv.Itoa(i), v.day3)
			file2.SetCellValue("长护记账", "T"+strconv.Itoa(i), v.num4)
			file2.SetCellValue("长护记账", "U"+strconv.Itoa(i), v.day4)
			file2.SetCellValue("长护记账", "V"+strconv.Itoa(i), v.num5)
			file2.SetCellValue("长护记账", "W"+strconv.Itoa(i), v.day5)
		}
		i++
	}
	defer file2.Close()
	file2.SaveAs("../10月汇总 - 1号调整(1)(1).xlsx")

}

func NumberOf(huliyuans map[string]huliyuan2) int {
	i := 0
	for k, _ := range huliyuans {
		k = k
		i++
	}
	return i
}

func count(i *huliyuan, e excel1) huliyuan {
	long := Long(e)
	for k, _ := range long {
		time1 := e.endDate.Unix() - e.startDate.Unix()
		if float64(time1) >= k*3600 && !TheyRe(e) {
			i.huliyuans[e.name1] = huliyuan2{
				num1: i.huliyuans[e.name1].num1 + 1,
				day1: i.huliyuans[e.name1].day1 + "\n" + e.startDate.String()[:10],
				num2: i.huliyuans[e.name1].num2,
				day2: i.huliyuans[e.name1].day2,
				num3: i.huliyuans[e.name1].num3,
				day3: i.huliyuans[e.name1].day3,
				num4: i.huliyuans[e.name1].num4,
				day4: i.huliyuans[e.name1].day4,
				num5: i.huliyuans[e.name1].num5,
				day5: i.huliyuans[e.name1].day5,
				num6: i.huliyuans[e.name1].num6,
				day6: i.huliyuans[e.name1].day6,
			}
		}
		if float64(time1) < k*3600 {
			i.huliyuans[e.name1] = huliyuan2{
				num1: i.huliyuans[e.name1].num1,
				day1: i.huliyuans[e.name1].day1,
				num2: i.huliyuans[e.name1].num2,
				day2: i.huliyuans[e.name1].day2,
				num3: i.huliyuans[e.name1].num3,
				day3: i.huliyuans[e.name1].day3,
				num4: i.huliyuans[e.name1].num4,
				day4: i.huliyuans[e.name1].day4,
				num5: i.huliyuans[e.name1].num5 + 1,
				day5: i.huliyuans[e.name1].day5 + "\n" + e.startDate.String()[:10],
				num6: i.huliyuans[e.name1].num6,
				day6: i.huliyuans[e.name1].day6,
			}
		}
		if float64(time1) > 12*3600 {
			i.huliyuans[e.name1] = huliyuan2{
				num1: i.huliyuans[e.name1].num1,
				day1: i.huliyuans[e.name1].day1,
				num2: i.huliyuans[e.name1].num2 + 1,
				day2: i.huliyuans[e.name1].day2 + "\n" + e.startDate.String()[:10],
				num3: i.huliyuans[e.name1].num3,
				day3: i.huliyuans[e.name1].day3,
				num4: i.huliyuans[e.name1].num4,
				day4: i.huliyuans[e.name1].day4,
				num5: i.huliyuans[e.name1].num5,
				day5: i.huliyuans[e.name1].day5,
				num6: i.huliyuans[e.name1].num6,
				day6: i.huliyuans[e.name1].day6,
			}
		}
		if TheyRe(e) {
			i.huliyuans[e.name1] = huliyuan2{
				num1: i.huliyuans[e.name1].num1,
				day1: i.huliyuans[e.name1].day1,
				num2: i.huliyuans[e.name1].num2,
				day2: i.huliyuans[e.name1].day2,
				num3: i.huliyuans[e.name1].num3 + 1,
				day3: i.huliyuans[e.name1].day3 + "\n" + e.startDate.String()[:10],
				num4: i.huliyuans[e.name1].num4,
				day4: i.huliyuans[e.name1].day4,
				num5: i.huliyuans[e.name1].num5,
				day5: i.huliyuans[e.name1].day5,
				num6: i.huliyuans[e.name1].num6,
				day6: i.huliyuans[e.name1].day6,
			}
		}

	}
	return *i
}

func TestName4(t *testing.T) {
	file, _ := excelize.OpenFile("../10月打卡汇总.xlsx")
	list1, _ := GetStockList1(file)
	for index, _ := range list1 {
		fmt.Printf("%v", list1[index])
	}
}

/*func TestName1(t *testing.T) {
	m := make(map[string]excel3, 1)
	m[]
}
*/

func TestName2(t *testing.T) {
	h := new(huliyuan2)
	fmt.Printf("%v", h)
}

func TestName3(t *testing.T) {
	local, _ := time.LoadLocation("Asia/Shanghai")
	startDate, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-11-15 11:04:55", local)
	endDate, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-11-15 12:04:55", local)
	i := endDate.Unix() - startDate.Unix()
	fmt.Println(i)

}

func TestName5(t *testing.T) {

}
