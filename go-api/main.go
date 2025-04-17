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

type tradeUpdate struct {
	Country        *string `json:"Country"`
	Year           *int    `json:"Year"`
	Commodity_code *string `json:"Commodity_code"`
	Commodity_desc *string `json:"Commodity_desc"`
	Flow           *string `json:"Flow"`
	Trade_usd      *int    `json:"Trade_usd"`
}

func updateTradeById(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid trade ID"})
		return
	}

	var update tradeUpdate
	if err := c.BindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid JSON body"})
		return
	}

	for i, t := range trades {
		if t.ID == id {
			if update.Country != nil {
				trades[i].Country = *update.Country
			}
			if update.Year != nil {
				trades[i].Year = *update.Year
			}
			if update.Commodity_code != nil {
				trades[i].Commodity_code = *update.Commodity_code
			}
			if update.Commodity_desc != nil {
				trades[i].Commodity_desc = *update.Commodity_desc
			}
			if update.Flow != nil {
				trades[i].Flow = *update.Flow
			}
			if update.Trade_usd != nil {
				trades[i].Trade_usd = *update.Trade_usd
			}

			c.IndentedJSON(http.StatusOK, trades[i])
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "trade not found"})
}

func main() {
	router := gin.Default()
	router.GET("/trades", getTrades)
	router.GET("/trades/filter", getFilteredTrades)
	router.POST("/trades", postTrades)
	router.PATCH("trades/:id", updateTradeById)

	router.Run("localhost:8080")
}
