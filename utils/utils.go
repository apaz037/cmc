package utils

import (
	"fmt"
	"github.com/leekchan/accounting"
)

func FormatPrice(price float64, symbol string, precision int) string {
	ac := accounting.Accounting{
		Symbol: symbol,
		Precision: precision,
	}

	result := ac.FormatMoney(price)

	return result
}

func BuildStringOutput(ticker string, price string) string {
	output := fmt.Sprintln(ticker, ": ", price)
	return output
}