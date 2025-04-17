package main

//Country_or_area,Year,comm_code,commodity,Flow,Trade_usd,

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"fmt"
	"strconv"
)

type trade struct {
	ID             int    `json:"id"`
	Country        string `json:"Country"`
	Year           int    `json:"Year"`
	Commodity_code string `json:"Commodity_code"`
	Commodity_desc string `json:"Commodity_desc"`
	Flow           string `json:"Flow"`
	Trade_usd      int    `json:"Trade_usd"`
}

var trades = []trade{
	{ID: 1, Country: "Afghanistan", Year: 2016, Commodity_code: "010410", Commodity_desc: "Sheep, live", Flow: "Export", Trade_usd: 6088},
	{ID: 2, Country: "Argentina", Year: 1996, Commodity_code: "010410", Commodity_desc: "Sheep, live", Flow: "Import", Trade_usd: 798372},
	{ID: 3, Country: "Australia", Year: 2007, Commodity_code: "010410", Commodity_desc: "Sheep, live", Flow: "Export", Trade_usd: 219992664},
}

func getTrades(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, trades)
}

func getTradesByYear(c *gin.Context) {
	// conversion is required from string in JSON to int
	year, err := strconv.Atoi(c.Param("year"))

	if err != nil {
		fmt.Println(err)
	}

	var matchingTrades = []trade{}

	// Loop over list of trades, looking for
	// a trade whose Year matches the parameter.
	for _, a := range trades {
		if a.Year == year {
			matchingTrades = append(matchingTrades, a)
		}
	}

	// return all matching trades
	c.IndentedJSON(http.StatusOK, matchingTrades)

}

func postTrades(c *gin.Context) {
	var newTrade trade

	// Call BindJSON to bind the received JSON to
	// newTrade.
	if err := c.BindJSON(&newTrade); err != nil {
		return
	}

	// Add the new trade to the slice.
	trades = append(trades, newTrade)
	c.IndentedJSON(http.StatusCreated, newTrade)
}

// getTradeByCountry

func main() {
	//fmt.Println(trades)

	router := gin.Default()
	router.GET("/trades", getTrades)
	router.GET("/trades/year=:year", getTradesByYear)
	router.POST("/trades", postTrades)

	router.Run("localhost:8080")
}
