package alphavantage

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Daily struct{
	MetaData   MetaData `json:"Meta Data"`
	TimeSeries map[string]*Entry `json:"Time Series (Daily)"`
}

type Intra struct {
	MetaData   MetaData `json:"Meta Data"`
	TimeSeries map[string]*Entry `json:"Time Series (5min)"`
}

func (c *Client) GetIntra(symbol string, size string) (Intra, error) {
	var data Intra
	err := json.Unmarshal(c.get(data,"TIME_SERIES_INTRADAY", "IBM", nil), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *Client) GetDaily(symbol string, size string) (Daily, error) {
	var data Daily
	err := json.Unmarshal(c.get(data,"TIME_SERIES_DAILY", "IBM", nil), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *Client) get(i interface{}, function string, symbol string, opts map[string]string) []byte {
	var data interface{}

	switch i.(type) {
	case Daily:
		data = Daily{}
	case Intra:
		data = Intra{}
	}

	parameters := url.Values{}
	parameters.Add("function", function)
	parameters.Add("symbol", symbol)
	parameters.Add("apikey", c.ApiKey)
	for k, v := range opts {
		parameters.Add(k, v)
	}

	c.BaseUrl.RawQuery = parameters.Encode()

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}

	result, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	return result
}