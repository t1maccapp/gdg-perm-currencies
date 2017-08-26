package main

import "time"

type Rates struct {
	base  ConvertibleCurrency
	date  time.Time
	rates map[ConvertibleCurrency]float64
}
