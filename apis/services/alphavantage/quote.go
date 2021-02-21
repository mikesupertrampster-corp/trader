package alphavantage

type Quote struct {
	GlobalQuote struct {
		Symbol           string  `json:"01. symbol"`
		Open             float32 `json:"02. open"`
		High             float32 `json:"03. high"`
		Low              float32 `json:"04. low"`
		Price            string  `json:"05. price"`
		Volume           int32   `json:"06. volume"`
		LatestTradingDay string  `json:"07. latest trading day"`
		PreviousClose    float32 `json:"08. previous close"`
		Change           float32 `json:"09. change"`
		ChangePercent    string  `json:"10. change percent"`
	} `json:"Global Quote"`
}
