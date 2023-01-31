package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
)

func main() {
	// create a new collector
	c := colly.NewCollector()

	// authenticate
	err := c.Post("http://113.214.250.148:8080/submitted_system_new/", map[string]string{"username": "hsdsp", "password": "hsdso@666"})
	if err != nil {
		log.Fatal(err)
	}

	// attach callbacks after login
	c.OnResponse(func(r *colly.Response) {
		fmt.Println(r.Headers)
	})

	// start scraping
	c.Visit("http://113.214.250.148:8080/submitted_system_new/main/home")

}
