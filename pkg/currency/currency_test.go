package currency

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	t.Run("Test Prints.", func(t *testing.T) {
		var b Bitcoin
		b = 289
		t.Log(b.Code())
		t.Log(b.Country())
		t.Log(b.Symbol())
		t.Log(b.Name())
		t.Log(b.ExchangeRate(b))
		t.Log(b)
	})
}
