package tiingo

import (
	"fmt"
	"net/http"
	"net/url"
)

type Client struct {
	ApiKey  string
	BaseUrl url.URL
}

func New(apiKey string) Client {
	u := url.URL{
		Scheme: "https",
		Host:   "api.tiingo.com",
		Path:   "tiingo",
	}
	return Client{
		ApiKey:  apiKey,
		BaseUrl: u,
	}
}

func (c *Client) GetPrice(symbol string) error {

	parameters := url.Values{}
	parameters.Add("token", c.ApiKey)
	c.BaseUrl.RawQuery = parameters.Encode()
	c.BaseUrl.Path += fmt.Sprintf("/daily/%s/prices", symbol)

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		return err
	}

	_, err = http.DefaultClient.Do(req)
	if err != nil {
		return err
	}

	//json.NewDecoder(res.Body).Decode(&quote)
	return nil
}
