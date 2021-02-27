package alphavantage

import "encoding/json"

type IncomeStatement struct {
	AnnualReports    []Report `json:"annualReports"`
	QuarterlyReports []Report `json:"quarterlyReports"`
}

func (c *Client) GetIncomeStatement(symbol string) (IncomeStatement, error) {
	var data IncomeStatement

	series, err := c.get(data, "INCOME_STATEMENT", symbol, nil)
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(series, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
