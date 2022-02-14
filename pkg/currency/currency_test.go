package currency

import (
	"testing"
)

func TestCurrency(t *testing.T) {
	t.Run("Test Prints.", func(t *testing.T) {
		var b Bitcoin
		b = 289
		t.Error(b.Code())
		t.Error(b.Country())
		t.Error(b.Symbol())
		t.Error(b.Name())
		t.Error(b.ExchangeRate(b))
		t.Error(b)
	})
}
