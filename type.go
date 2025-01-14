package main

import "fmt"

type Money struct {
	Amount int
	Currency string
}


func (m *Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}
