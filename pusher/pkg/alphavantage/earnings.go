package alphavantage

import "encoding/json"

type Earnings struct {
	Annual []struct {
		FiscalDateEnding string  `json:"fiscalDateEnding"`
		ReportedEPS      float64 `json:"reportedEPS,string"`
	} `json:"annualEarnings"`
	Quarterly []struct {
		FiscalDateEnding   string  `json:"fiscalDateEnding"`
		ReportedDate       string  `json:"reportedDate"`
		ReportedEPS        float64 `json:"reportedEPS,string"`
		EstimatedEPS       float64 `json:"estimatedEPS,string"`
		Surprise           float64 `json:"surprise,string"`
		SurprisePercentage float64 `json:"surprisePercentage,string"`
	} `json:"quarterlyEarnings"`
}

func (c *Client) GetEarnings(symbol string) (Series, error) {
	var data Earnings
	var series Series

	response, err := c.get(data, "EARNINGS", symbol, nil)
	if err != nil {
		return series, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		return series, err
	}

	return c.extract(data, "2006-01-02")
}
