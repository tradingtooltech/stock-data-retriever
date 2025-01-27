package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"stock-data-retriever/handlers"
	dataModels "stock-data-retriever/models"

	"github.com/aws/aws-lambda-go/lambda"
)

var requestMap = map[string]func(symbol string, dailyOpenCloseDate string) (*dataModels.StockDataResponse, error){
	"dailyOpenCloseData": handlers.GetDailyOpenCloseData,
}

func handleRequest(ctx context.Context, event json.RawMessage) (*dataModels.StockDataResponse, error) {
	var stockDataRequest dataModels.StockDataRequest

	if err := json.Unmarshal(event, &stockDataRequest); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return nil, err
	}

	if getDailyOpenCloseData, ok := requestMap[stockDataRequest.RequestType]; ok {
		stockDataResponse, stockDataErr := getDailyOpenCloseData(
			stockDataRequest.Symbol, stockDataRequest.DailyOpenCloseDate)
		if stockDataErr != nil {
			return nil, stockDataErr
		}
		return stockDataResponse, nil
	} else {
		return nil, errors.New("invalid request")
	}
}

func main() {
	lambda.Start(handleRequest)
}
