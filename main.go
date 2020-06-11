package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal"
	"net/http"
)

// Balance stores the balance to debit or credit
type Balance struct {
	Balance string `json:"balance" binding:"required"`
}

func main() {
	// Temporary storage
	wallets := make(map[string]*wallet.Wallet)
	wallets["123"], _ = wallet.New("123", "0.0")
	wallets["456"], _ = wallet.New("456", "12.0")

	r := gin.Default()

	// Gets the wallet balance
	r.GET("/api/v1/wallets/:walletid/balance", func(c *gin.Context) {
		id := c.Param("walletid")
		c.JSON(http.StatusOK, gin.H{
			"balance": wallets[id].PrintBalance(),
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
		if err := wallets[id].Credit(json.Balance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

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
		if err := wallets[id].Debit(json.Balance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
