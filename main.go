package main

//Country_or_area,Year,comm_code,commodity,Flow,Trade_usd,

import (
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

func getFilteredTrades(c *gin.Context) {
	yearParam := c.Query("year")
	countryParam := c.Query("country")
	commodityParam := c.Query("commodity_code")

	var filteredTrades []trade

	for _, t := range trades {
		match := true

		if yearParam != "" {
			year, err := strconv.Atoi(yearParam)
			if err != nil || t.Year != year {
				match = false
			}
		}

		if countryParam != "" && t.Country != countryParam {
			match = false
		}

		if commodityParam != "" && t.Commodity_code != commodityParam {
			match = false
		}

		if match {
			filteredTrades = append(filteredTrades, t)
		}

	}

	if len(filteredTrades) == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "no matching trades found"})
		return
	}
	c.IndentedJSON(http.StatusOK, filteredTrades)
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

func patchTradesById(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid trade id"})
		return
	}

	yearParam := c.Query("year")
	countryParam := c.Query("country")
	commodityParam := c.Query("commodity_code")
	flowParam := c.Query("flow")

	var params = []string{yearParam, countryParam, commodityParam, flowParam}

	//var filteredTrades []trade

	for _, t := range trades {
		if t.ID == id {
			for _, p := range params {

			}

		}

		year, err := strconv.Atoi(yearParam)
		if err != nil || t.Year != year {
			match = false
		}

	}
}

// getTradeByCountry

func main() {
	//fmt.Println(trades)

	router := gin.Default()
	router.GET("/trades", getTrades)
	router.GET("/trades/filter", getFilteredTrades)
	router.POST("/trades", postTrades)
	router.PATCH("trades", patchTradesById)

	router.Run("localhost:8080")
}
