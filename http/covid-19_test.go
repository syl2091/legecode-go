package code

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	get, err := http.Get("https://coronavirus.m.pipedream.net/")
	if err != nil {
		println(err)
	}
	defer get.Body.Close()
	if get.StatusCode == 200 {
		all, _ := ioutil.ReadAll(get.Body)
		ioutil.WriteFile("covid.json", all, 0644)
	}
}
