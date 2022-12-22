package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func AgeOfName() {
	get, err := http.Get("https://api.agify.io?name=沈永乐")
	if err != nil {
		println(err)
	}
	if get.StatusCode == 200 {
		all, _ := ioutil.ReadAll(get.Body)
		fmt.Println(string(all))
	}
}
