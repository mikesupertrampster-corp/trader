package alphavantage

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"time"
)

var log *logrus.Logger

type Series []DataPoint

type DataPoint struct {
	Name      string
	Data 	  map[string]interface{}
	Timestamp time.Time
}

type Client struct {
	ApiKey  string
	BaseUrl url.URL
}

func New(logger *logrus.Logger, apiKey string) Client {
	log = logger

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
		log.Error("Could not create new http request: ", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Could not make http request: ", err)
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("Could not read response body: ", err)
		return nil, err
	}

	data := reflect.Indirect(reflect.ValueOf(i)).Interface()
	err = json.NewDecoder(strings.NewReader(strings.ReplaceAll(string(b), "None", "0"))).Decode(&data)
	if err != nil {
		log.Error("Could not decode json string: ", err)
		return nil, err
	}

	result, err := json.Marshal(data)
	if err != nil {
		log.Error("Could not marshal data: ", err)
		return nil, err
	}

	return result, nil
}

func (c *Client) extract(data interface{}, timeLayout string) (Series, error) {
	var series Series

	fundamental := reflect.TypeOf(data).Name()
	v := reflect.ValueOf(data)
	for idx := 0; idx < v.NumField(); idx++ {
		field := v.Field(idx)
		name := v.Type().Field(idx).Name

		switch field.Kind() {
		case reflect.Slice:
			for i := 0; i < field.Len(); i++ {
				report := field.Index(i)
				date := reflect.Indirect(report).FieldByName("FiscalDateEnding").String()

				timestamp, err := time.Parse(timeLayout, date)
				if err != nil {
					log.Fatal("Failed to parse timestamp: ", err)
					return series, err
				}

				series = append(series, DataPoint{
					fmt.Sprintf("%s_%s", fundamental, name),
					c.toIf(report.Interface()),
					timestamp,
				})
			}
		case reflect.Map:
			for _, dateTime := range field.MapKeys() {
				report := field.MapIndex(dateTime)

				timestamp, err := time.Parse(timeLayout, dateTime.String())
				if err != nil {
					log.Fatal("Failed to parse timestamp: ", err)
					return series, err
				}

				series = append(series, DataPoint{
					fmt.Sprintf("%s_%s", fundamental, name),
					c.toIf(report.Elem().Interface()),
					timestamp,
				})
			}
		}
	}

	return series, nil
}

func (c *Client) toIf(i interface{}) map[string]interface{} {
	iface := make(map[string]interface{})

	v := reflect.ValueOf(i)
	for i := 0; i < v.NumField(); i++ {
		iface[strings.ToLower(v.Type().Field(i).Name)] = v.Field(i)
	}

	return iface
}