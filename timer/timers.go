package main

import (
	"fmt"
	"time"
)

//定时器

func main() {
	timer1 := time.NewTimer(2 * time.Second)
	<-timer1.C
	fmt.Println("timer1")

	//定时器出发之前取消
	timer2 := time.NewTimer(time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2")
	}()
	stop := timer2.Stop()
	if stop {
		fmt.Println("timer2 stoped")
	}

	// 给 `timer2` 足够的时间来触发它，以证明它实际上已经停止了。
	time.Sleep(2 * time.Second)
}
