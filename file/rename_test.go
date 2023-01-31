package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type rows struct {
	Fdeleted        int           `json:"fdeleted"`
	Faddperson      string        `json:"faddperson"`
	FbindColumnTime string        `json:"fbindColumnTime"`
	Sort            string        `json:"sort"`
	Faddtime        string        `json:"faddtime"`
	VideoSubmitType string        `json:"videoSubmitType"`
	BatchId         string        `json:"batchId"`
	UserAreas       []interface{} `json:"userAreas"`
	Columnid        int           `json:"columnid"`
	EndTime         string        `json:"endTime"`
	SideCode        string        `json:"sideCode"`
	AreaId          string        `json:"areaId"`
	UserType        int           `json:"userType"`
	StartTime       string        `json:"startTime"`
	Id              int           `json:"id"`
	AudioUrl        string        `json:"audioUrl"`
	Order           string        `json:"order"`
	Fhits           int           `json:"fhits"`
	Ftypeid         int           `json:"ftypeid"`
	ColumnName      string        `json:"columnName"`
	FendTime        string        `json:"fendTime"`
	Min             int           `json:"min"`
	Limit           int           `json:"limit"`
	FsubmitTime     string        `json:"fsubmitTime"`
	Fsubmit         int           `json:"fsubmit"`
	Holder          interface{}   `json:"holder"`
	ImageSrc        string        `json:"imageSrc"`
	Max             int           `json:"max"`
	FcheckedTime    string        `json:"fcheckedTime"`
	Fupdatetime     string        `json:"fupdatetime"`
	VideoUrl        string        `json:"videoUrl"`
	FstartTime      string        `json:"fstartTime"`
	Areaid          int64         `json:"areaid"`
	OrderByClause   string        `json:"orderByClause"`
	Fchecked        int           `json:"fchecked"`
	Fimage          int           `json:"fimage"`
	Ftitle          string        `json:"ftitle"`
	Fsummary        string        `json:"fsummary"`
	Py              string        `json:"py"`
	Start           int           `json:"start"`
	Fcontent        string        `json:"fcontent"`
	Audioid         int           `json:"audioid"`
}

func TestName(t *testing.T) {
	var rows []rows
	data, err := ioutil.ReadFile("./json2.json")
	if err != nil {
		println(err)
	}
	err1 := json.Unmarshal(data[:], &rows)
	fmt.Println(err1)
	if err != nil {
		panic(err1)
	}
	for _, v := range rows {
		fmt.Println(v.Ftitle)
		file, err := os.ReadFile("./videos")
		for
		if err != nil {
			panic(err)
		}
	}

}

func TestName2(t *testing.T) {
	err := filepath.Walk("./videos", func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, ".mp4") {

		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
