package main

import (
	"fmt"
	"time"
)

//超时

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "result 1"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(2 * time.Second):
		fmt.Println("timeout 1")
	}
}
