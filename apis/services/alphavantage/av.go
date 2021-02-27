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

type Report struct {
	AccountsPayable                   int64   `json:"accountsPayable,string"`
	AccumulatedAmortization           string  `json:"accumulatedAmortization"`
	AccumulatedDepreciation           int64   `json:"accumulatedDepreciation,string"`
	AdditionalPaidInCapital           int64   `json:"additionalPaidInCapital,string"`
	CapitalExpenditures               int64   `json:"capitalExpenditures,string"`
	CapitalLeaseObligations           string  `json:"capitalLeaseObligations"`
	CapitalSurplus                    string  `json:"capitalSurplus"`
	Cash                              int64   `json:"cash,string"`
	CashAndShortTermInvestments       int64   `json:"cashAndShortTermInvestments,string"`
	CashflowFromFinancing             int64   `json:"cashflowFromFinancing,string"`
	CashflowFromInvestment            int64   `json:"cashflowFromInvestment,string"`
	ChangeInAccountReceivables        int64   `json:"changeInAccountReceivables,string"`
	ChangeInCash                      int64   `json:"changeInCash,string"`
	ChangeInCashAndCashEquivalents    string  `json:"changeInCashAndCashEquivalents"`
	ChangeInExchangeRate              string  `json:"changeInExchangeRate"`
	ChangeInInventory                 string  `json:"changeInInventory"`
	ChangeInLiabilities               int64   `json:"changeInLiabilities,string"`
	ChangeInNetIncome                 int64   `json:"changeInNetIncome,string"`
	ChangeInOperatingActivities       int64   `json:"changeInOperatingActivities,string"`
	ChangeInReceivables               string  `json:"changeInReceivables"`
	CommonStock                       int64   `json:"commonStock,string"`
	CommonStockSharesOutstanding      int64   `json:"commonStockSharesOutstanding,string"`
	CommonStockTotalEquity            int64   `json:"commonStockTotalEquity,string"`
	CostOfRevenue                     int64   `json:"costOfRevenue,string"`
	CurrentLongTermDebt               int64   `json:"currentLongTermDebt,string"`
	DeferredLongTermAssetCharges      int64   `json:"deferredLongTermAssetCharges,string"`
	DeferredLongTermLiabilities       int64   `json:"deferredLongTermLiabilities,string"`
	Depreciation                      int64   `json:"depreciation,string"`
	DiscontinuedOperations            int64   `json:"discontinuedOperations,string"`
	DividendPayout                    int64   `json:"dividendPayout,string"`
	EarningAssets                     string  `json:"earningAssets"`
	Ebit                              int64   `json:"ebit,string"`
	EffectOfAccountingCharges         string  `json:"effectOfAccountingCharges"`
	EstimatedEPS                      float64 `json:"estimatedEPS,string"`
	ExtraordinaryItems                int64   `json:"extraordinaryItems,string"`
	FiscalDateEnding                  string  `json:"fiscalDateEnding"`
	Goodwill                          int64   `json:"goodwill,string"`
	GrossProfit                       int64   `json:"grossProfit,string"`
	IncomeBeforeTax                   int64   `json:"incomeBeforeTax,string"`
	IncomeTaxExpense                  int64   `json:"incomeTaxExpense,string"`
	IntangibleAssets                  int64   `json:"intangibleAssets,string"`
	InterestExpense                   int64   `json:"interestExpense,string"`
	InterestIncome                    string  `json:"interestIncome"`
	Inventory                         int64   `json:"inventory,string"`
	Investments                       int64   `json:"investments,string"`
	LiabilitiesAndShareholderEquity   int64   `json:"liabilitiesAndShareholderEquity,string"`
	LongTermDebt                      int64   `json:"longTermDebt,string"`
	LongTermInvestments               int64   `json:"longTermInvestments,string"`
	MinorityInterest                  int64   `json:"minorityInterest,string"`
	NegativeGoodwill                  string  `json:"negativeGoodwill"`
	NetBorrowings                     int64   `json:"netBorrowings,string"`
	NetIncome                         int64   `json:"netIncome,string"`
	NetIncomeApplicableToCommonShares int64   `json:"netIncomeApplicableToCommonShares,string"`
	NetIncomeFromContinuingOperations int64   `json:"netIncomeFromContinuingOperations,string"`
	NetInterestIncome                 int64   `json:"netInterestIncome,string"`
	NetReceivables                    int64   `json:"netReceivables,string"`
	NetTangibleAssets                 int64   `json:"netTangibleAssets,string"`
	NonRecurring                      string  `json:"nonRecurring"`
	OperatingCashflow                 int64   `json:"operatingCashflow,string"`
	OperatingIncome                   int64   `json:"operatingIncome,string"`
	OtherAssets                       int64   `json:"otherAssets,string"`
	OtherCashflowFromFinancing        string  `json:"otherCashflowFromFinancing"`
	OtherCashflowFromInvestment       int64   `json:"otherCashflowFromInvestment,string"`
	OtherCurrentAssets                int64   `json:"otherCurrentAssets,string"`
	OtherCurrentLiabilities           int64   `json:"otherCurrentLiabilities,string"`
	OtherItems                        string  `json:"otherItems"`
	OtherLiabilities                  int64   `json:"otherLiabilities,string"`
	OtherNonCurrentLiabilities        int64   `json:"otherNonCurrentLiabilities,string"`
	OtherNonCurrrentAssets            int64   `json:"otherNonCurrrentAssets,string"`
	OtherNonOperatingIncome           string  `json:"otherNonOperatingIncome"`
	OtherOperatingCashflow            string  `json:"otherOperatingCashflow"`
	OtherOperatingExpense             int64   `json:"otherOperatingExpense,string"`
	OtherShareholderEquity            int64   `json:"otherShareholderEquity,string"`
	PreferredStockAndOtherAdjustments string  `json:"preferredStockAndOtherAdjustments"`
	PreferredStockRedeemable          string  `json:"preferredStockRedeemable"`
	PreferredStockTotalEquity         int64   `json:"preferredStockTotalEquity,string"`
	PropertyPlantEquipment            int64   `json:"propertyPlantEquipment,string"`
	ReportedCurrency                  string  `json:"reportedCurrency"`
	ReportedDate                      string  `json:"reportedDate"`
	ReportedEPS                       float64 `json:"reportedEPS,string"`
	ResearchAndDevelopment            int64   `json:"researchAndDevelopment,string"`
	RetainedEarnings                  int64   `json:"retainedEarnings,string"`
	RetainedEarningsTotalEquity       int64   `json:"retainedEarningsTotalEquity,string"`
	SellingGeneralAdministrative      int64   `json:"sellingGeneralAdministrative,string"`
	ShortTermDebt                     int64   `json:"shortTermDebt,string"`
	ShortTermInvestments              int64   `json:"shortTermInvestments,string"`
	StockSaleAndPurchase              int64   `json:"stockSaleAndPurchase,string"`
	Surprise                          float64 `json:"surprise,string"`
	SurprisePercentage                float64 `json:"surprisePercentage,string"`
	TaxProvision                      int64   `json:"taxProvision,string"`
	TotalAssets                       int64   `json:"totalAssets,string"`
	TotalCurrentAssets                int64   `json:"totalCurrentAssets,string"`
	TotalCurrentLiabilities           int64   `json:"totalCurrentLiabilities,string"`
	TotalLiabilities                  int64   `json:"totalLiabilities,string"`
	TotalLongTermDebt                 int64   `json:"totalLongTermDebt,string"`
	TotalNonCurrentAssets             int64   `json:"totalNonCurrentAssets,string"`
	TotalNonCurrentLiabilities        int64   `json:"totalNonCurrentLiabilities,string"`
	TotalOperatingExpense             int64   `json:"totalOperatingExpense,string"`
	TotalOtherIncomeExpense           int64   `json:"totalOtherIncomeExpense,string"`
	TotalPermanentEquity              int64   `json:"totalPermanentEquity,string"`
	TotalRevenue                      int64   `json:"totalRevenue,string"`
	TotalShareholderEquity            int64   `json:"totalShareholderEquity,string"`
	TreasuryStock                     int64   `json:"treasuryStock,string"`
	Warrants                          string  `json:"warrants"`
}