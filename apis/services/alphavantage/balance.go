package alphavantage

import "encoding/json"

type Sheet struct {
	FiscalDateEnding                string `json:"fiscalDateEnding"`
	ReportedCurrency                string `json:"reportedCurrency"`
	TotalAssets                     int64  `json:"totalAssets,string"`
	IntangibleAssets                int64  `json:"intangibleAssets,string"`
	EarningAssets                   string `json:"earningAssets"`
	OtherCurrentAssets              int64  `json:"otherCurrentAssets,string"`
	TotalLiabilities                int64  `json:"totalLiabilities,string"`
	TotalShareholderEquity          int64  `json:"totalShareholderEquity,string"`
	DeferredLongTermLiabilities     int64  `json:"deferredLongTermLiabilities,string"`
	OtherCurrentLiabilities         int64  `json:"otherCurrentLiabilities,string"`
	CommonStock                     int64  `json:"commonStock,string"`
	RetainedEarnings                int64  `json:"retainedEarnings,string"`
	OtherLiabilities                int64  `json:"otherLiabilities,string"`
	Goodwill                        int64  `json:"goodwill,string"`
	OtherAssets                     int64  `json:"otherAssets,string"`
	Cash                            int64  `json:"cash,string"`
	TotalCurrentLiabilities         int64  `json:"totalCurrentLiabilities,string"`
	ShortTermDebt                   int64  `json:"shortTermDebt,string"`
	CurrentLongTermDebt             int64  `json:"currentLongTermDebt,string"`
	OtherShareholderEquity          int64  `json:"otherShareholderEquity,string"`
	PropertyPlantEquipment          int64  `json:"propertyPlantEquipment,string"`
	TotalCurrentAssets              int64  `json:"totalCurrentAssets,string"`
	LongTermInvestments             int64  `json:"longTermInvestments,string"`
	NetTangibleAssets               int64  `json:"netTangibleAssets,string"`
	ShortTermInvestments            int64  `json:"shortTermInvestments,string"`
	NetReceivables                  int64  `json:"netReceivables,string"`
	LongTermDebt                    int64  `json:"longTermDebt,string"`
	Inventory                       int64  `json:"inventory,string"`
	AccountsPayable                 int64  `json:"accountsPayable,string"`
	TotalPermanentEquity            int64  `json:"totalPermanentEquity,string"`
	AdditionalPaidInCapital         int64  `json:"additionalPaidInCapital,string"`
	CommonStockTotalEquity          int64  `json:"commonStockTotalEquity,string"`
	PreferredStockTotalEquity       int64  `json:"preferredStockTotalEquity,string"`
	RetainedEarningsTotalEquity     int64  `json:"retainedEarningsTotalEquity,string"`
	TreasuryStock                   int64  `json:"treasuryStock,string"`
	AccumulatedAmortization         string `json:"accumulatedAmortization"`
	OtherNonCurrrentAssets          int64  `json:"otherNonCurrrentAssets,string"`
	DeferredLongTermAssetCharges    int64  `json:"deferredLongTermAssetCharges,string"`
	TotalNonCurrentAssets           int64  `json:"totalNonCurrentAssets,string"`
	CapitalLeaseObligations         string `json:"capitalLeaseObligations"`
	TotalLongTermDebt               int64  `json:"totalLongTermDebt,string"`
	OtherNonCurrentLiabilities      int64  `json:"otherNonCurrentLiabilities,string"`
	TotalNonCurrentLiabilities      int64  `json:"totalNonCurrentLiabilities,string"`
	NegativeGoodwill                string `json:"negativeGoodwill"`
	Warrants                        string `json:"warrants"`
	PreferredStockRedeemable        string `json:"preferredStockRedeemable"`
	CapitalSurplus                  string `json:"capitalSurplus"`
	LiabilitiesAndShareholderEquity int64  `json:"liabilitiesAndShareholderEquity,string"`
	CashAndShortTermInvestments     int64  `json:"cashAndShortTermInvestments,string"`
	AccumulatedDepreciation         int64  `json:"accumulatedDepreciation,string"`
	CommonStockSharesOutstanding    int64  `json:"commonStockSharesOutstanding,string"`
}

type BalanceSheet struct {
	Symbol           string    `json:"symbol"`
	AnnualReports    []Reports `json:"annualReports"`
	QuarterlyReports []Reports `json:"quarterlyReports"`
}

func (c *Client) GetBalanceSheet(symbol string) (BalanceSheet, error) {
	var data BalanceSheet
	err := json.Unmarshal(c.get(data, "BALANCE_SHEET", symbol, nil), &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
