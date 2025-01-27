package models

type StockDataRequest struct {
	RequestType        string `json:"requestType"`
	Symbol             string `json:"symbol"`
	DailyOpenCloseDate string `json:"dailyOpenCloseDate"`
}

type StockDataResponse struct {
	Symbol             string  `json:"symbol"`
	DailyOpenCloseDate string  `json:"dailyOpenCloseDate"`
	Open               float64 `json:"open"`
	Close              float64 `json:"close"`
}
