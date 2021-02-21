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

func (c *Client) GetIntra(symbol string, size string) (Daily, error) {
	return c.get(Daily{},"TIME_SERIES_INTRADAY", symbol, map[string]string{"outputsize": size}).(Daily), nil
}

func (c *Client) GetDaily(symbol string, size string) (Daily, error) {
	return c.get(Daily{},"TIME_SERIES_DAILY", symbol, map[string]string{"outputsize": size}).(Daily), nil
}

func (c *Client) get(iface interface{}, function string, symbol string, opts map[string]string) interface{} {
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

	err = json.NewDecoder(res.Body).Decode(&iface)
	if err != nil {
		log.Fatal(err)
	}

	return iface
}
