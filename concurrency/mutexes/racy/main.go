package main

import (
	"strconv"
)

var myBalance = &balance{amount: 50.00, currency: "GBP"}

type balance struct {
	amount   float64
	currency string
}

func (b *balance) Add(i float64) {
	b.amount += i
}

func (b *balance) Display() string {
	return strconv.FormatFloat(b.amount, 'f', 2, 64) + " " + b.currency
}
