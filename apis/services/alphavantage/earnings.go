package alphavantage

import "encoding/json"

type Earnings struct {
	Symbol         string `json:"symbol"`
	AnnualEarnings []struct {
		FiscalDateEnding string  `json:"fiscalDateEnding"`
		ReportedEPS      float64 `json:"reportedEPS,string"`
	} `json:"annualEarnings"`
	QuarterlyEarnings []struct {
		FiscalDateEnding   string  `json:"fiscalDateEnding"`
		ReportedDate       string  `json:"reportedDate"`
		ReportedEPS        float64 `json:"reportedEPS,string"`
		EstimatedEPS       float64 `json:"estimatedEPS,string"`
		Surprise           float64 `json:"surprise,string"`
		SurprisePercentage float64 `json:"surprisePercentage,string"`
	} `json:"quarterlyEarnings"`
}

func (c *Client) GetEarnings(symbol string) (Earnings, error) {
	var data Earnings
	err := json.Unmarshal(c.get(data, "EARNINGS", symbol, nil), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
