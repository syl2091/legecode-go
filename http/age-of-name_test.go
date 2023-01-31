package code

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

func TestAgeOfName(t *testing.T) {
	get, err := http.Get("https://api.agify.io?name=沈永乐&country_id=CN")
	if err != nil {
		println(err)
	}
	if get.StatusCode == 200 {
		all, _ := ioutil.ReadAll(get.Body)
		fmt.Println(string(all))
	}
}
