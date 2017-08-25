package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
)

const FIXER_URL = "http://api.fixer.io/latest?base="

type FixerResponse struct {
	Name string `json:"base"`
	Date string `json:"date"`
	Rates map[string] float64 `json:"rates"`
}

func Convert(base ConvertibleCurrency, amount float64) (map[Currency]float64, error) {
	ratesUrl := FIXER_URL + base.Name()
	fixerResponse, err := getRates(ratesUrl)

	if err != nil{
		return nil, err
	}

	result := make(map[Currency]float64)

	for currencyName, rate := range fixerResponse.Rates {
		currency := Currency{name: currencyName}
		result[currency] = rate * amount
	}

	return result, nil
}

func getRates(ratesUrl string) (FixerResponse, error) {
	response, err := http.Get(ratesUrl)
	defer response.Body.Close()

	if err != nil {
		return FixerResponse{}, err
	}

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return FixerResponse{}, err
	}

	var fixerResponse FixerResponse

	json.Unmarshal(body, &fixerResponse)
	if err != nil {
		return FixerResponse{}, err
	}

	return fixerResponse, nil
}
