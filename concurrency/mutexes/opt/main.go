package main

import (
	"strconv"
	"sync"
)

var myBalance = &balance{amount: 50.00, currency: "GBP"}

type balance struct {
	amount   float64
	currency string
	mu       sync.Mutex
}

func (b *balance) Add(i float64) {
	b.mu.Lock()
	b.amount += i
	b.mu.Unlock()
}

func (b *balance) Display() string {
	b.mu.Lock()
	defer b.mu.Unlock()
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}
