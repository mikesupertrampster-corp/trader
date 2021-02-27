package alphavantage

import (
	"encoding/json"
)

type MetaData struct {
	Symbol string `json:"2. Symbol"`
	Last   string `json:"3. Last Refreshed"`
}

type Entry struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume float64 `json:"5. volume,string"`
}

type Daily struct {
	MetaData   MetaData          `json:"Meta AVData"`
	TimeSeries map[string]*Entry `json:"Time Series (Daily)"`
}

type Intra struct {
	MetaData   MetaData          `json:"Meta AVData"`
	TimeSeries map[string]*Entry `json:"Time Series (5min)"`
}

func (c *Client) GetIntra(symbol string, interval string, size string) (Intra, error) {
	var data Intra
	opts := map[string]string{
		"interval":   interval,
		"outputsize": size,
	}

	series, err := c.get(data, "TIME_SERIES_INTRADAY", symbol, opts)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(series, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func (c *Client) GetDaily(symbol string, size string) (Daily, error) {
	var data Daily

	series, err := c.get(data, "TIME_SERIES_DAILY", symbol, map[string]string{"outputsize": size})
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(series, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
