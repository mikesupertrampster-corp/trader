package alphavantage

import (
	"net/url"
)

type Client struct {
	ApiKey  string
	BaseUrl url.URL
}

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

func New(apiKey string) Client {
	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "query",
	}
	return Client{
		ApiKey:  apiKey,
		BaseUrl: u,
	}
}
