package handlers

import (
	"context"
	"log"
	"os"
	"time"

	dataModels "stock-data-retriever/models"

	polygon "github.com/polygon-io/client-go/rest"
	"github.com/polygon-io/client-go/rest/models"
)

var polygonIoApiKey = os.Getenv("POLYGON_IO_API_KEY")

func GetDailyOpenCloseData(symbol string, dailyOpenCloseDate string) (*dataModels.StockDataResponse, error) {
	newClient := polygon.New(polygonIoApiKey)

	parsedDate, parsedDateErr := time.Parse("2006-01-02", dailyOpenCloseDate)

	if parsedDateErr != nil {
		return nil, parsedDateErr
	}

	params := &models.GetDailyOpenCloseAggParams{
		Ticker: symbol,
		Date:   models.Date(parsedDate),
	}

	res, err := newClient.GetDailyOpenCloseAgg(context.Background(), params)

	if err != nil {
		log.Printf("Error getting daily open close data: %v", err)
		return nil, err
	}

	stockDataResponse := dataModels.StockDataResponse{
		Symbol:             res.Symbol,
		DailyOpenCloseDate: res.From,
		Open:               res.Open,
		Close:              res.Close,
	}

	return &stockDataResponse, nil
}
