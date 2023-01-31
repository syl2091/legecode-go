package concurrency

import (
	"fmt"
	"net/http"
	"testing"
)

func TestName(t *testing.T) {
	response, err := http.Head("https://blog.csdn.net/")
	if err != nil {
		fmt.Println(response.Body)
	}
	fmt.Println(response.StatusCode)
}
