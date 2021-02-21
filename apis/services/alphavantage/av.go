package alphavantage

import (
	"net/url"
)

type Client struct {
	ApiKey  string
	BaseUrl url.URL
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
