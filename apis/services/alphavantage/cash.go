package alphavantage

import "encoding/json"

type CashFlow struct {
	AnnualReports    []Report `json:"annualReports"`
	QuarterlyReports []Report `json:"quarterlyReports"`
}

func (c *Client) CashFlow(symbol string) (CashFlow, error) {
	var data CashFlow

	series, err := c.get(data, "CASH_FLOW", symbol, nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(series, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
