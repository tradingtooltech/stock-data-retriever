# Stock Data Retriever

This application supports retrieving daily open/close data for a given stock.

Sample event:

`
{
    "requestType": "dailyOpenCloseData",
    "symbol": "AAPL",
    "dailyOpenCloseDate": "2025-01-02"
}
`

Sample response:

`
{
  "symbol": "AAPL",
  "dailyOpenCloseDate": "2025-01-02",
  "open": 248.93,
  "close": 243.85
}
`
