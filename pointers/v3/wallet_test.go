package main

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	err := wallet.Withdraw(9)
	if err != nil {
		panic(err)
	}
	fmt.Println(wallet.balance)
}
