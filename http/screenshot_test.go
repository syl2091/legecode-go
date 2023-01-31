package code

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func TestScreenshotWebSite(t *testing.T) {
	apiKey := "54fcbc57d3254dd39c31c253a26898f1"
	url := "https://tophub.today/c/developer"
	site := ScreenshotWebSite(url, apiKey)
	fmt.Println(site)
}

func TestScreenshotWebSiteServer(t *testing.T) {
	server := http.HandlerFunc(ScreenshotWebSiteServer)
	log.Fatalln(http.ListenAndServe(":5000", server))
}
