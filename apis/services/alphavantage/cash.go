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
	Symbol           string    `json:"symbol"`
	AnnualReports    []Reports `json:"annualReports"`
	QuarterlyReports []Reports `json:"quarterlyReports"`
}

func (c *Client) CashFlow(symbol string) (CashFlow, error) {
	var data CashFlow
	err := json.Unmarshal(c.get(data, "CASH_FLOW", symbol, nil), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
