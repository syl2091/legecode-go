package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

//网站截图

func ScreenshotWebSite(url string, apiKey string) []byte {
	post, err := http.Post(fmt.Sprintf("https://api.apiflash.com/v1/urltoimage?"+
		"access_key=%s&wait_until=page_loaded&url=%s", apiKey, url), "image/jpeg", nil)
	if err != nil {
		println(err)
	}
	if post.StatusCode == 200 {
		all, _ := ioutil.ReadAll(post.Body)
		return all
	}
	return nil
}

func ScreenshotWebSiteServer(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	url := query.Get("url")
	apiKey := query.Get("apiKey")
	png := ScreenshotWebSite(url, apiKey)
	w.Header().Set("Content-Type", "image/png")
	w.Write(png)
}
