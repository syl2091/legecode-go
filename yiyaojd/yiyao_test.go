package main

import (
	"fmt"
	"testing"
)

func TestYiyao(t *testing.T) {
	get := Get("https://item.yiyaojd.com/2943430.html?ad_od=share&utm_term=Wxfriends&utm_campaign=t_335139774&gx=RnExkWRQajzcytRP--tyC_Jy5PXa6jY4ictI&utm_source=iosapp&utm_medium=appshare")
	fmt.Println(get)
}
