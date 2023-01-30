package main

import (
	"strconv"
	"sync"
)

var mu = &sync.Mutex{}
var myBalance = &balance{amount: 50.00, currency: "GBP"}

type balance struct {
	amount   float64
	currency string
}

func (b *balance) Add(i float64) {
	mu.Lock()
	b.amount += i
	mu.Unlock()
}

func (b *balance) Display() string {
	mu.Lock()
	amt := b.amount
	cur := b.currency
	mu.Unlock()
	return strconv.FormatFloat(amt, 'f', 2, 64) + " " + cur
}
