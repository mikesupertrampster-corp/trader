package alphavantage

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
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

func (c *Client) get(i interface{}, function string, symbol string, opts map[string]string) ([]byte, error) {
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
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	data := reflect.Indirect(reflect.ValueOf(i)).Interface()
	err = json.NewDecoder(strings.NewReader(strings.ReplaceAll(string(b), "None", "0"))).Decode(&data)
	if err != nil {
		return nil, err
	}

	result, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}