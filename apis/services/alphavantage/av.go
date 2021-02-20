package alphavantage

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey  string
	BaseUrl url.URL
}

type Daily struct {
	MetaData struct {
		Symbol string `json:"2. Symbol"`
		Last   string `json:"3. Last Refreshed"`
	} `json:"Meta Data"`

	TimeSeries map[string]*Entry `json:"Time Series (Daily)"`
}

type Entry struct {
	Open   float64 `json:"1. open,string"`
	High   float64 `json:"2. high,string"`
	Low    float64 `json:"3. low,string"`
	Close  float64 `json:"4. close,string"`
	Volume float64 `json:"5. volume,string"`
}

type Quote struct {
	GlobalQuote struct {
		Symbol           string  `json:"01. symbol"`
		Open             float32 `json:"02. open"`
		High             float32 `json:"03. high"`
		Low              float32 `json:"04. low"`
		Price            string  `json:"05. price"`
		Volume           int32   `json:"06. volume"`
		LatestTradingDay string  `json:"07. latest trading day"`
		PreviousClose    float32 `json:"08. previous close"`
		Change           float32 `json:"09. change"`
		ChangePercent    string  `json:"10. change percent"`
	} `json:"Global Quote"`
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

func (c *Client) GetDaily(symbol string) (Daily, error) {
	var daily Daily

	parameters := url.Values{}
	parameters.Add("function", "TIME_SERIES_DAILY")
	parameters.Add("symbol", symbol)
	parameters.Add("apikey", c.ApiKey)
	c.BaseUrl.RawQuery = parameters.Encode()

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		return daily, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return daily, err
	}

	err = json.NewDecoder(res.Body).Decode(&daily)
	if err != nil {
		return daily, err
	}

	return daily, nil
}

func (c *Client) GetQuote(symbol string) (Quote, error) {
	var quote Quote

	parameters := url.Values{}
	parameters.Add("function", "GLOBAL_QUOTE")
	parameters.Add("symbol", symbol)
	parameters.Add("apikey", c.ApiKey)
	c.BaseUrl.RawQuery = parameters.Encode()

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		return quote, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return quote, err
	}

	json.NewDecoder(res.Body).Decode(&quote)
	return quote, nil
}