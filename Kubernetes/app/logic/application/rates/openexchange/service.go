package openexchange

import (
	"log"
)

type Service struct {
	Dao IDao
	Logger *log.Logger
}

func (s Service) GetLatestExchangeRates() (*Info, error) {
	response,err := s.Dao.FetchLatest();
	if err != nil {
		s.Logger.Printf("there was an error fetching latest exchange rate info: %s\n", err.Error())
		return nil, ErrFetchingResponse
	}
	info := &Info{
		Base: response.Base,
		Rates: response.Rates,
	}
	return info, nil
}
