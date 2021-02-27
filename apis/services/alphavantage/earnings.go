package alphavantage

import "encoding/json"

type Earnings struct {
	AnnualReports    Report `json:"annualEarnings"`
	QuarterlyReports Report `json:"quarterlyEarnings"`
}

func (c *Client) GetEarnings(symbol string) (Earnings, error) {
	var data Earnings

	series, err := c.get(data, "EARNINGS", symbol, nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(series, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
