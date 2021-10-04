package openexchange

type Response struct {
	Disclaimer string `json:"disclaimer"`
	License    string `json:"license"`
	Timestamp  int    `json:"timestamp"`
	Base       string `json:"base"`
	Rates      map[string]float64 `json:"rates"`
}

type Info struct {
	Base string
	Rates map[string]float64
}

type IDao interface {
	FetchLatest() (*Response, error)
}

type IService interface {
	GetLatestExchangeRates() (*Info, error)
}