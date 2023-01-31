package code

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func pyq(apiKey string) {
	get, err := http.Get(fmt.Sprintf("https://apis.tianapi.com/pyqwenan/index?key=%s", apiKey))
	if err != nil {
		println(err)
	}
	defer get.Body.Close()
	if get.StatusCode == 200 {
		all, _ := ioutil.ReadAll(get.Body)
		fmt.Println(string(all))
	}
}
