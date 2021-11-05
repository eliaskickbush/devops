package openexchange

import "errors"

// Service errors
var ErrFetchingResponse = errors.New("there was an error fetching data from dao")

// Dao errors
var ErrMalformedURL = errors.New("there was an error building request url")
var ErrRequestExecution = errors.New("there was an error executing request")
var ErrReadingBody = errors.New("there was an error reading returned body")
var ErrUnmarshallingBody = errors.New("there was an error unmarshalling body")