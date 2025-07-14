package util

const (
	USD = "USD"
	EUR = "EUR"
	GBP = "GBP"
)

func IsSuppoortedCurrency(currency string) bool {
	switch currency {
	case USD, EUR, GBP:
		return true
	}
	return false
}
