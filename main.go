package main

//country_or_area,year,comm_code,commodity,flow,trade_usd,

type tradeRecord struct {
	ID		int64 `json:"id"`
	country string `json:"country"`
	year	string `json:"year"`
	commodity_code string `json:"commodity_code"`
	commodity_desc string `json:"commodity_desc"`
	flow string `json:"flow"`
	trade_usd int64 `json:"trade_usd"`
}

var 
1, Afghanistan,2016,010410,"Sheep, live",Export,6088
2, Argentina,1996,010410,"Sheep, live",Import,798372
3, Australia,2007,010410,"Sheep, live",Export,219992664