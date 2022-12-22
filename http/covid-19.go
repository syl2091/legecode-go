package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	get, err := http.Get("https://coronavirus.m.pipedream.net/")
	if err != nil {
		println(err)
	}
	defer get.Body.Close()
	if get.StatusCode == 200 {
		all, _ := ioutil.ReadAll(get.Body)
		fmt.Println(string(all))
	}
}
