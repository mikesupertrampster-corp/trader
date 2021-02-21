package alphavantage

type Quote struct {
	GlobalQuote struct {
		Symbol           string  `json:"01. symbol"`
		Open             float64 `json:"02. open,string"`
		High             float64 `json:"03. high,string"`
		Low              float64 `json:"04. low,string"`
		Price            string  `json:"05. price,string"`
		Volume           int64   `json:"06. volume,string"`
		LatestTradingDay string  `json:"07. latest trading day"`
		PreviousClose    float64 `json:"08. previous close,string"`
		Change           float64 `json:"09. change,string"`
		ChangePercent    string  `json:"10. change percent"`
	} `json:"Global Quote"`
}
