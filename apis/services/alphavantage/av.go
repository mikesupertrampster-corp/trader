package alphavantage

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
)

var log *logrus.Logger

type Client struct {
	ApiKey  string
	BaseUrl url.URL
}

func New(logger *logrus.Logger, apiKey string) Client {
	log = logger

	u := url.URL{
		Scheme: "https",
		Host:   "www.alphavantage.co",
		Path:   "query",
	}

	return Client{
		ApiKey:  apiKey,
		BaseUrl: u,
	}
}

func (c *Client) get(i interface{}, function string, symbol string, opts map[string]string) ([]byte, error) {
	parameters := url.Values{}
	parameters.Add("function", function)
	parameters.Add("symbol", symbol)
	parameters.Add("apikey", c.ApiKey)
	for k, v := range opts {
		parameters.Add(k, v)
	}

	c.BaseUrl.RawQuery = parameters.Encode()

	req, err := http.NewRequest(http.MethodGet, c.BaseUrl.String(), nil)
	if err != nil {
		log.Error("Could not create new http request: ", err)
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Error("Could not make http request: ", err)
		return nil, err
	}

	b, err := io.ReadAll(res.Body)
	if err != nil {
		log.Error("Could not read response body: ", err)
		return nil, err
	}

	data := reflect.Indirect(reflect.ValueOf(i)).Interface()
	err = json.NewDecoder(strings.NewReader(strings.ReplaceAll(string(b), "None", "0"))).Decode(&data)
	if err != nil {
		log.Error("Could not decode json string: ", err)
		return nil, err
	}

	result, err := json.Marshal(data)
	if err != nil {
		log.Error("Could not marshal data: ", err)
		return nil, err
	}

	return result, nil
}

type Report []ReportFields

type ReportFields struct {
	AccountsPayable                   int64   `json:"accountsPayable,string,omitempty"`
	AccumulatedAmortization           string  `json:"accumulatedAmortization,omitempty"`
	AccumulatedDepreciation           int64   `json:"accumulatedDepreciation,string,omitempty"`
	AdditionalPaidInCapital           int64   `json:"additionalPaidInCapital,string,omitempty"`
	CapitalExpenditures               int64   `json:"capitalExpenditures,string,omitempty"`
	CapitalLeaseObligations           string  `json:"capitalLeaseObligations,omitempty"`
	CapitalSurplus                    string  `json:"capitalSurplus,omitempty"`
	Cash                              int64   `json:"cash,string,omitempty"`
	CashAndShortTermInvestments       int64   `json:"cashAndShortTermInvestments,string,omitempty"`
	CashflowFromFinancing             int64   `json:"cashflowFromFinancing,string,omitempty"`
	CashflowFromInvestment            int64   `json:"cashflowFromInvestment,string,omitempty"`
	ChangeInAccountReceivables        int64   `json:"changeInAccountReceivables,string,omitempty"`
	ChangeInCash                      int64   `json:"changeInCash,string,omitempty"`
	ChangeInCashAndCashEquivalents    string  `json:"changeInCashAndCashEquivalents,omitempty"`
	ChangeInExchangeRate              string  `json:"changeInExchangeRate,omitempty"`
	ChangeInInventory                 string  `json:"changeInInventory,omitempty"`
	ChangeInLiabilities               int64   `json:"changeInLiabilities,string,omitempty"`
	ChangeInNetIncome                 int64   `json:"changeInNetIncome,string,omitempty"`
	ChangeInOperatingActivities       int64   `json:"changeInOperatingActivities,string,omitempty"`
	ChangeInReceivables               string  `json:"changeInReceivables,omitempty"`
	CommonStock                       int64   `json:"commonStock,string,omitempty"`
	CommonStockSharesOutstanding      int64   `json:"commonStockSharesOutstanding,string,omitempty"`
	CommonStockTotalEquity            int64   `json:"commonStockTotalEquity,string,omitempty"`
	CostOfRevenue                     int64   `json:"costOfRevenue,string,omitempty"`
	CurrentLongTermDebt               int64   `json:"currentLongTermDebt,string,omitempty"`
	DeferredLongTermAssetCharges      int64   `json:"deferredLongTermAssetCharges,string,omitempty"`
	DeferredLongTermLiabilities       int64   `json:"deferredLongTermLiabilities,string,omitempty"`
	Depreciation                      int64   `json:"depreciation,string,omitempty"`
	DiscontinuedOperations            int64   `json:"discontinuedOperations,string,omitempty"`
	DividendPayout                    int64   `json:"dividendPayout,string,omitempty"`
	EarningAssets                     string  `json:"earningAssets,omitempty"`
	Ebit                              int64   `json:"ebit,string,omitempty"`
	EffectOfAccountingCharges         string  `json:"effectOfAccountingCharges,omitempty"`
	EstimatedEPS                      float64 `json:"estimatedEPS,string,omitempty"`
	ExtraordinaryItems                int64   `json:"extraordinaryItems,string,omitempty"`
	FiscalDateEnding                  string  `json:"fiscalDateEnding,omitempty"`
	Goodwill                          int64   `json:"goodwill,string,omitempty"`
	GrossProfit                       int64   `json:"grossProfit,string,omitempty"`
	IncomeBeforeTax                   int64   `json:"incomeBeforeTax,string,omitempty"`
	IncomeTaxExpense                  int64   `json:"incomeTaxExpense,string,omitempty"`
	IntangibleAssets                  int64   `json:"intangibleAssets,string,omitempty"`
	InterestExpense                   int64   `json:"interestExpense,string,omitempty"`
	InterestIncome                    string  `json:"interestIncome,omitempty"`
	Inventory                         int64   `json:"inventory,string,omitempty"`
	Investments                       int64   `json:"investments,string,omitempty"`
	LiabilitiesAndShareholderEquity   int64   `json:"liabilitiesAndShareholderEquity,string,omitempty"`
	LongTermDebt                      int64   `json:"longTermDebt,string,omitempty"`
	LongTermInvestments               int64   `json:"longTermInvestments,string,omitempty"`
	MinorityInterest                  int64   `json:"minorityInterest,string,omitempty"`
	NegativeGoodwill                  string  `json:"negativeGoodwill,omitempty"`
	NetBorrowings                     int64   `json:"netBorrowings,string,omitempty"`
	NetIncome                         int64   `json:"netIncome,string,omitempty"`
	NetIncomeApplicableToCommonShares int64   `json:"netIncomeApplicableToCommonShares,string,omitempty"`
	NetIncomeFromContinuingOperations int64   `json:"netIncomeFromContinuingOperations,string,omitempty"`
	NetInterestIncome                 int64   `json:"netInterestIncome,string,omitempty"`
	NetReceivables                    int64   `json:"netReceivables,string,omitempty"`
	NetTangibleAssets                 int64   `json:"netTangibleAssets,string,omitempty"`
	NonRecurring                      string  `json:"nonRecurring,omitempty"`
	OperatingCashflow                 int64   `json:"operatingCashflow,string,omitempty"`
	OperatingIncome                   int64   `json:"operatingIncome,string,omitempty"`
	OtherAssets                       int64   `json:"otherAssets,string,omitempty"`
	OtherCashflowFromFinancing        string  `json:"otherCashflowFromFinancing,omitempty"`
	OtherCashflowFromInvestment       int64   `json:"otherCashflowFromInvestment,string,omitempty"`
	OtherCurrentAssets                int64   `json:"otherCurrentAssets,string,omitempty"`
	OtherCurrentLiabilities           int64   `json:"otherCurrentLiabilities,string,omitempty"`
	OtherItems                        string  `json:"otherItems,omitempty"`
	OtherLiabilities                  int64   `json:"otherLiabilities,string,omitempty"`
	OtherNonCurrentLiabilities        int64   `json:"otherNonCurrentLiabilities,string,omitempty"`
	OtherNonCurrrentAssets            int64   `json:"otherNonCurrrentAssets,string,omitempty"`
	OtherNonOperatingIncome           string  `json:"otherNonOperatingIncome,omitempty"`
	OtherOperatingCashflow            string  `json:"otherOperatingCashflow,omitempty"`
	OtherOperatingExpense             int64   `json:"otherOperatingExpense,string,omitempty"`
	OtherShareholderEquity            int64   `json:"otherShareholderEquity,string,omitempty"`
	PreferredStockAndOtherAdjustments string  `json:"preferredStockAndOtherAdjustments,omitempty"`
	PreferredStockRedeemable          string  `json:"preferredStockRedeemable,omitempty"`
	PreferredStockTotalEquity         int64   `json:"preferredStockTotalEquity,string,omitempty"`
	PropertyPlantEquipment            int64   `json:"propertyPlantEquipment,string,omitempty"`
	ReportedCurrency                  string  `json:"reportedCurrency,omitempty"`
	ReportedDate                      string  `json:"reportedDate,omitempty"`
	ReportedEPS                       float64 `json:"reportedEPS,string,omitempty"`
	ResearchAndDevelopment            int64   `json:"researchAndDevelopment,string,omitempty"`
	RetainedEarnings                  int64   `json:"retainedEarnings,string,omitempty"`
	RetainedEarningsTotalEquity       int64   `json:"retainedEarningsTotalEquity,string,omitempty"`
	SellingGeneralAdministrative      int64   `json:"sellingGeneralAdministrative,string,omitempty"`
	ShortTermDebt                     int64   `json:"shortTermDebt,string,omitempty"`
	ShortTermInvestments              int64   `json:"shortTermInvestments,string,omitempty"`
	StockSaleAndPurchase              int64   `json:"stockSaleAndPurchase,string,omitempty"`
	Surprise                          float64 `json:"surprise,string,omitempty"`
	SurprisePercentage                float64 `json:"surprisePercentage,string,omitempty"`
	TaxProvision                      int64   `json:"taxProvision,string,omitempty"`
	TotalAssets                       int64   `json:"totalAssets,string,omitempty"`
	TotalCurrentAssets                int64   `json:"totalCurrentAssets,string,omitempty"`
	TotalCurrentLiabilities           int64   `json:"totalCurrentLiabilities,string,omitempty"`
	TotalLiabilities                  int64   `json:"totalLiabilities,string,omitempty"`
	TotalLongTermDebt                 int64   `json:"totalLongTermDebt,string,omitempty"`
	TotalNonCurrentAssets             int64   `json:"totalNonCurrentAssets,string,omitempty"`
	TotalNonCurrentLiabilities        int64   `json:"totalNonCurrentLiabilities,string,omitempty"`
	TotalOperatingExpense             int64   `json:"totalOperatingExpense,string,omitempty"`
	TotalOtherIncomeExpense           int64   `json:"totalOtherIncomeExpense,string,omitempty"`
	TotalPermanentEquity              int64   `json:"totalPermanentEquity,string,omitempty"`
	TotalRevenue                      int64   `json:"totalRevenue,string,omitempty"`
	TotalShareholderEquity            int64   `json:"totalShareholderEquity,string,omitempty"`
	TreasuryStock                     int64   `json:"treasuryStock,string,omitempty"`
	Warrants                          string  `json:"warrants,omitempty"`
}