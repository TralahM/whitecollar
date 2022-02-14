package currency

import "fmt"

type Bitcoin float64

func (b Bitcoin) Code() string    { return "BTC" }
func (b Bitcoin) Country() string { return "WORLD" }
func (b Bitcoin) Symbol() string  { return "BTC" }
func (b Bitcoin) Name() string    { return "Bitcoin" }
func (b Bitcoin) String() string  { return fmt.Sprintf("%f BTC", b) }

func (b Bitcoin) ExchangeRate(c Currency) float64 {
	// TODO. Lookup exchange rate by code
	return 1.00
}
