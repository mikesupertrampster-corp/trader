package alphavantage

import (
	"encoding/json"
	"time"
)

type CompanyOverview struct {
	Symbol                     string  `json:"Symbol"`
	AssetType                  string  `json:"AssetType"`
	Name                       string  `json:"Name"`
	Description                string  `json:"Description"`
	Exchange                   string  `json:"Exchange"`
	Currency                   string  `json:"Currency"`
	Country                    string  `json:"Country"`
	Sector                     string  `json:"Sector"`
	Industry                   string  `json:"Industry"`
	Address                    string  `json:"Address"`
	FullTimeEmployees          int     `json:"FullTimeEmployees,string"`
	FiscalYearEnd              string  `json:"FiscalYearEnd"`
	LatestQuarter              string  `json:"LatestQuarter"`
	MarketCapitalization       int64   `json:"MarketCapitalization,string"`
	EBITDA                     int64   `json:"EBITDA,string"`
	PERatio                    float64 `json:"PERatio,string"`
	PEGRatio                   float64 `json:"PEGRatio,string"`
	BookValue                  float64 `json:"BookValue,string"`
	DividendPerShare           float64 `json:"DividendPerShare,string"`
	DividendYield              float64 `json:"DividendYield,string"`
	EPS                        float64 `json:"EPS,string"`
	RevenuePerShareTTM         float64 `json:"RevenuePerShareTTM,string"`
	ProfitMargin               float64 `json:"ProfitMargin,string"`
	OperatingMarginTTM         float64 `json:"OperatingMarginTTM,string"`
	ReturnOnAssetsTTM          float64 `json:"ReturnOnAssetsTTM,string"`
	ReturnOnEquityTTM          float64 `json:"ReturnOnEquityTTM,string"`
	RevenueTTM                 int64   `json:"RevenueTTM,string"`
	GrossProfitTTM             int64   `json:"GrossProfitTTM,string"`
	DilutedEPSTTM              float64 `json:"DilutedEPSTTM,string"`
	QuarterlyEarningsGrowthYOY float64 `json:"QuarterlyEarningsGrowthYOY,string"`
	QuarterlyRevenueGrowthYOY  float64 `json:"QuarterlyRevenueGrowthYOY,string"`
	AnalystTargetPrice         int64   `json:"AnalystTargetPrice,string"`
	TrailingPE                 float64 `json:"TrailingPE,string"`
	ForwardPE                  float64 `json:"ForwardPE,string"`
	PriceToSalesRatioTTM       float64 `json:"PriceToSalesRatioTTM,string"`
	PriceToBookRatio           float64 `json:"PriceToBookRatio,string"`
	EVToRevenue                float64 `json:"EVToRevenue,string"`
	EVToEBITDA                 float64 `json:"EVToEBITDA,string"`
	Beta                       float64 `json:"Beta,string"`
	Five2WeekHigh              float64 `json:"52WeekHigh,string"`
	Five2WeekLow               float64 `json:"52WeekLow,string"`
	Five0DayMovingAverage      float64 `json:"50DayMovingAverage,string"`
	Two00DayMovingAverage      float64 `json:"200DayMovingAverage,string"`
	SharesOutstanding          int64   `json:"SharesOutstanding,string"`
	SharesFloat                int64   `json:"SharesFloat,string"`
	SharesShort                int64   `json:"SharesShort,string"`
	SharesShortPriorMonth      int64   `json:"SharesShortPriorMonth,string"`
	ShortRatio                 float64 `json:"ShortRatio,string"`
	ShortPercentOutstanding    float64 `json:"ShortPercentOutstanding,string"`
	ShortPercentFloat          float64 `json:"ShortPercentFloat,string"`
	PercentInsiders            float64 `json:"PercentInsiders,string"`
	PercentInstitutions        float64 `json:"PercentInstitutions,string"`
	ForwardAnnualDividendRate  float64 `json:"ForwardAnnualDividendRate,string"`
	ForwardAnnualDividendYield float64 `json:"ForwardAnnualDividendYield,string"`
	PayoutRatio                float64 `json:"PayoutRatio,string"`
	DividendDate               string  `json:"DividendDate"`
	ExDividendDate             string  `json:"ExDividendDate"`
	LastSplitFactor            string  `json:"LastSplitFactor"`
	LastSplitDate              string  `json:"LastSplitDate"`
}

func (c *Client) GetCompanyOverview(symbol string) (Series, error) {
	var data CompanyOverview
	var series Series

	response, err := c.get(data, "OVERVIEW", symbol, nil)
	if err != nil {
		log.Fatal("Failed to get company overview: ", err)
		return series, err
	}

	err = json.Unmarshal(response, &data)
	if err != nil {
		log.Fatal("Failed to get unmarshall company overview: ", err)
		return series, err
	}

	timestamp, err := time.Parse("2006-01-02", data.LatestQuarter)
	if err != nil {
		log.Fatal("Failed to parse timestamp: ", err)
		return series, err
	}

	series = append(series, DataPoint{
		"overview",
		c.toIf(data),
		timestamp,
	})

	return series, nil
}