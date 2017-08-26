package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const FIXER_URL = "http://api.fixer.io/latest?base="

type FixerResponse struct {
	Name  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func Convert(base ConvertibleCurrency, amount float64) (map[Currency]float64, error) {
	ratesUrl := FIXER_URL + base.Name()
	fixerResponse, err := getRates(ratesUrl)

	if err != nil {
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

	if err != nil {
		return FixerResponse{}, err
	}

	defer response.Body.Close()

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
