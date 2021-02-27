package alphavantage

import "encoding/json"

type BalanceSheet struct {
	AnnualReports    []Report `json:"annualReports"`
	QuarterlyReports []Report `json:"quarterlyReports"`
}

func (c *Client) GetBalanceSheet(symbol string) (BalanceSheet, error) {
	var data BalanceSheet

	series, err := c.get(data, "BALANCE_SHEET", symbol, nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(series, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
