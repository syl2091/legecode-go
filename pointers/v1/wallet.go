package main

import "fmt"

type Bitcoin int

func (b Bitcoin) String() string {
	//TODO implement me
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(bitcoin Bitcoin) {
	w.balance += bitcoin
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
