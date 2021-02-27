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
	Price    map[string]*Entry `json:"Time Series (Daily)"`
}

type Intra struct {
	Price    map[string]*Entry `json:"Time Series (5min)"`
}

func (c *Client) GetIntra(symbol string, interval string, size string) (Series, error) {
	var data Intra
	opts := map[string]string{
		"interval":   interval,
		"outputsize": size,
	}
	var series Series

	response, err := c.get(data, "TIME_SERIES_INTRADAY", symbol, opts)
	if err != nil {
		log.Fatal("Failed to get intra day prices: ", err)
		return series, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Fatal("Failed to get unmarshall intra day price: ", err)
		return series, err
	}

	return c.extract(data, "2006-01-02 15:04:05")
}

func (c *Client) GetDaily(symbol string, size string) (Series, error) {
	var data Daily
	var series Series

	response, err := c.get(data, "TIME_SERIES_DAILY", symbol, map[string]string{"outputsize": size})
	if err != nil {
		log.Fatal("Failed to get daily prices: ", err)
		return series, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Fatal("Failed to get unmarshall daily price: ", err)
		return series, err
	}

	return c.extract(data, "2006-01-02")
}
