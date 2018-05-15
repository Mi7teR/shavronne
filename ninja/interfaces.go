package ninja

type CurrencyInterface interface {
	Price() string
	Name() string
	Icon() string
}
