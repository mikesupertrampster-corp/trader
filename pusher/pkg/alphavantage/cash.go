package alphavantage

import "encoding/json"

type Flow struct {
	FiscalDateEnding               string `json:"fiscalDateEnding"`
	ReportedCurrency               string `json:"reportedCurrency"`
	Investments                    int64  `json:"investments,string"`
	ChangeInLiabilities            int64  `json:"changeInLiabilities,string"`
	CashflowFromInvestment         int64  `json:"cashflowFromInvestment,string"`
	OtherCashflowFromInvestment    int64  `json:"otherCashflowFromInvestment,string"`
	NetBorrowings                  int64  `json:"netBorrowings,string"`
	CashflowFromFinancing          int64  `json:"cashflowFromFinancing,string"`
	OtherCashflowFromFinancing     string `json:"otherCashflowFromFinancing"`
	ChangeInOperatingActivities    int64  `json:"changeInOperatingActivities,string"`
	NetIncome                      int64  `json:"netIncome,string"`
	ChangeInCash                   int64  `json:"changeInCash,string"`
	OperatingCashflow              int64  `json:"operatingCashflow,string"`
	OtherOperatingCashflow         string `json:"otherOperatingCashflow"`
	Depreciation                   int64  `json:"depreciation,string"`
	DividendPayout                 int64  `json:"dividendPayout,string"`
	StockSaleAndPurchase           int64  `json:"stockSaleAndPurchase,string"`
	ChangeInInventory              string `json:"changeInInventory"`
	ChangeInAccountReceivables     int64  `json:"changeInAccountReceivables,string"`
	ChangeInNetIncome              int64  `json:"changeInNetIncome,string"`
	CapitalExpenditures            int64  `json:"capitalExpenditures,string"`
	ChangeInReceivables            string `json:"changeInReceivables"`
	ChangeInExchangeRate           string `json:"changeInExchangeRate"`
	ChangeInCashAndCashEquivalents string `json:"changeInCashAndCashEquivalents"`
}

type CashFlow struct {
	Annual    []Flow `json:"annualReports"`
	Quarterly []Flow `json:"quarterlyReports"`
}

func (c *Client) GetCashFlow(symbol string) (Series, error) {
	var data CashFlow
	var series Series

	response, err := c.get(data, "CASH_FLOW", symbol, nil)
	if err != nil {
		log.Fatal("Failed to get cash flows: ", err)
		return series, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Fatal("Failed to get unmarshall cash flow: ", err)
		return series, err
	}

	return c.extract(data, "2006-01-02")
}