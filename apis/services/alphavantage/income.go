package alphavantage

import "encoding/json"

type Reports struct {
	FiscalDateEnding                  string `json:"fiscalDateEnding"`
	ReportedCurrency                  string `json:"reportedCurrency"`
	TotalRevenue                      int64  `json:"totalRevenue,string"`
	TotalOperatingExpense             int64  `json:"totalOperatingExpense,string"`
	CostOfRevenue                     int64  `json:"costOfRevenue,string"`
	GrossProfit                       int64  `json:"grossProfit,string"`
	Ebit                              int64  `json:"ebit,string"`
	NetIncome                         int64  `json:"netIncome,string"`
	ResearchAndDevelopment            int64  `json:"researchAndDevelopment,string"`
	EffectOfAccountingCharges         string `json:"effectOfAccountingCharges"`
	IncomeBeforeTax                   int64  `json:"incomeBeforeTax,string"`
	MinorityInterest                  int64  `json:"minorityInterest,string"`
	SellingGeneralAdministrative      int64  `json:"sellingGeneralAdministrative,string"`
	OtherNonOperatingIncome           string `json:"otherNonOperatingIncome"`
	OperatingIncome                   int64  `json:"operatingIncome,string"`
	OtherOperatingExpense             int64  `json:"otherOperatingExpense,string"`
	InterestExpense                   int64  `json:"interestExpense,string"`
	TaxProvision                      int64  `json:"taxProvision,string"`
	InterestIncome                    string `json:"interestIncome"`
	NetInterestIncome                 int64  `json:"netInterestIncome,string"`
	ExtraordinaryItems                int64  `json:"extraordinaryItems,string"`
	NonRecurring                      string `json:"nonRecurring"`
	OtherItems                        string `json:"otherItems"`
	IncomeTaxExpense                  int64  `json:"incomeTaxExpense,string"`
	TotalOtherIncomeExpense           int64  `json:"totalOtherIncomeExpense,string"`
	DiscontinuedOperations            int64  `json:"discontinuedOperations,string"`
	NetIncomeFromContinuingOperations int64  `json:"netIncomeFromContinuingOperations,string"`
	NetIncomeApplicableToCommonShares int64  `json:"netIncomeApplicableToCommonShares,string"`
	PreferredStockAndOtherAdjustments string `json:"preferredStockAndOtherAdjustments"`
}

type IncomeStatement struct {
	Symbol           string    `json:"symbol"`
	AnnualReports    []Reports `json:"annualReports"`
	QuarterlyReports []Reports `json:"quarterlyReports"`
}

func (c *Client) GetIncomeStatement(symbol string) (IncomeStatement, error) {
	var data IncomeStatement
	err := json.Unmarshal(c.get(data, "INCOME_STATEMENT", symbol, nil), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
