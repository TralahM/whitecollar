package currency

type Currency interface {
	Code() string
	Country() string
	Symbol() string
	ExchangeRate(Currency) float64
	Name() string
}
