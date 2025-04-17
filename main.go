package main

//Country_or_area,Year,comm_code,commodity,Flow,Trade_usd,

import (
	"net/http"

	"github.com/gin-gonic/gin"
	//"fmt"
)

type trade struct {
	ID		int64 `json:"id"`
	Country string `json:"Country"`
	Year	int64 `json:"Year"`
	Commodity_code string `json:"Commodity_code"`
	Commodity_desc string `json:"Commodity_desc"`
	Flow string `json:"Flow"`
	Trade_usd int64 `json:"Trade_usd"`
}

var trades = []trade{
	{ID: 1, Country: "Afghanistan",Year: 2016,Commodity_code: "010410", Commodity_desc: "Sheep, live",Flow: "Export",Trade_usd: 6088},
	{ID: 2, Country: "Argentina",Year: 1996,Commodity_code: "010410", Commodity_desc: "Sheep, live",Flow: "Import",Trade_usd: 798372},
	{ID: 3, Country: "Australia",Year: 2007,Commodity_code: "010410",Commodity_desc: "Sheep, live",Flow: "Export",Trade_usd: 219992664},
}

func getTrades(c *gin.Context){
	c.IndentedJSON(http.StatusOK, trades)
}

func main() {
	//fmt.Println(trades)

	router := gin.Default()
	router.GET("/trades", getTrades)

	router.Run("localhost:8080")
}