package main

type ConvertibleCurrency interface {
	Name() string
}

type Currency struct {
	name string
}

func (currency Currency) Name() string {
	return currency.name
}