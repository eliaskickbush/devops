package rates

type IService interface {
	UpdateCache() error
	GetExchangeRate(symbol string) (float64,error)
	Convert(from string, to string, amount float64) (float64,error)
}