package rates

import (
	"kubertest/logic/application/rates/cache"
	"kubertest/logic/application/rates/openexchange"
	"log"
	"math"
	"strconv"
)

type Service struct {
	Base string
	CacheDao cache.IDao
	OpenExchangeService openexchange.IService
	Logger *log.Logger
}

func (s Service) GetExchangeRate(symbol string) (float64, error) {
	rateStr,err := s.CacheDao.GetValue(symbol)
	if err != nil {
		s.Logger.Printf("there was an error getting value for symbol %s: %s\n", symbol, err.Error())
		return 0, ErrFetchingRate
	}
	rate,err := strconv.ParseFloat(rateStr, 64)
	if err != nil {
		return 0, ErrParsingRate
	}
	return rate,nil
}

func (s Service) UpdateCache() error {
	info, err := s.OpenExchangeService.GetLatestExchangeRates()
	if err != nil {
		s.Logger.Printf("there was an error getting exchange rates: %s\n", err.Error())
		return ErrUpdatingCache
	}

	s.Base = info.Base
	for k,v := range info.Rates {
		err := s.CacheDao.SetValue(k,v)
		if err != nil {
			s.Logger.Printf("there was an error updating currency %s: %s\n", k, err.Error())
		}
	}

	return nil
}

func (s Service) Convert(from string, to string, amount float64) (float64, error) {
	fromRate, err := s.GetExchangeRate(from)
	if err != nil {
		s.Logger.Printf("there was an error fetching rate for symbol %s: %s\n", from, err.Error())
		return 0,err
	}

	if fromRate == 0 {
		s.Logger.Printf("unexpected zero value on fetched rate for symbol %s\n", from)
		return 0, ErrZeroValueRate
	}

	toRate, err := s.GetExchangeRate(to)
	if err != nil {
		s.Logger.Printf("there was an error fetching rate for symbol %s: %s\n", to, err.Error())
		return 0,err
	}


	return math.Floor(((amount / fromRate) * toRate) * 100) / 100, nil
}
