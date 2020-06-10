package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Balance stores the balance to debit or credit
type Balance struct {
	Balance int `json:"balance" binding:"required"`
}

func main() {
	balances := make(map[string]int)
	balances["123"] = 0
	balances["456"] = 12

	r := gin.Default()

	// Gets the wallet balance
	r.GET("/api/v1/wallets/:walletid/balance", func(c *gin.Context) {
		id := c.Param("walletid")
		balance := fmt.Sprintf("%d", balances[id])
		c.JSON(http.StatusOK, gin.H{
			"message": balance,
		})
	})

	// Credit
	r.POST("/api/v1/wallets/:walletid/credit", func(c *gin.Context) {
		var json Balance
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		id := c.Param("walletid")
		balances[id] += json.Balance
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// Debit
	r.POST("/api/v1/wallets/:walletid/debit", func(c *gin.Context) {
		var json Balance
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		id := c.Param("walletid")
		balances[id] -= json.Balance
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
