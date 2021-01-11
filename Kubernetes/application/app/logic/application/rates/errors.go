package rates

import "errors"

var ErrUpdatingCache = errors.New("there was an error updating exchange rate cache")
var ErrFetchingRate = errors.New("there was an error fetching exchange rate from cache dao")
var ErrParsingRate = errors.New("there was an error parsing retrieved rate as float")
var ErrZeroValueRate = errors.New("there was an zero value returned as an exchange rate")