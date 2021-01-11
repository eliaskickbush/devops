package openexchange

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Dao struct {
	AppID string
	BaseURL url.URL
	Logger *log.Logger
	HTTPClient http.Client
}

func (d Dao) FetchLatest() (*Response, error) {

	builtURL := fmt.Sprintf("%s%s", d.BaseURL.String(), "/latest.json")
	d.Logger.Printf("%s\n", builtURL)
	u, err := url.Parse(builtURL)
	if err != nil {
		d.Logger.Printf("unexpected error building request url: %s\n", err.Error())
		return nil,ErrMalformedURL
	}

	// Set App ID
	q := u.Query()
	q.Set("app_id", d.AppID)
	u.RawQuery = q.Encode()

	// Craft request
	req,_ := http.NewRequest(http.MethodGet, u.String(), nil)

	resp, err := d.HTTPClient.Do(req)
	if err != nil {
		d.Logger.Printf("unexpected error executing request: %s\n", err.Error())
		return nil,ErrRequestExecution
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		d.Logger.Printf("unexpected error reading body: %s\n", err.Error())
		return nil,ErrReadingBody
	}

	response := &Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		d.Logger.Printf("unexpected error unmarshalling body: %s\n", err.Error())
		return nil,ErrUnmarshallingBody
	}

	return response, nil
}
