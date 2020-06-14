// Package views exposes resources used for presenting data to the end user
package views

import (
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal/controllers"
	"github.com/shopspring/decimal"
	"net/http"
)

const (
	// EndpointGETBalance is the balance endpoint
	EndpointGETBalance = "/api/v1/wallets/:walletid/balance"
	// EndpointPOSTCredit is the credit endpoint
	EndpointPOSTCredit = "/api/v1/wallets/:walletid/credit"
	// EndpointPOSTDebit is the debit endpoint
	EndpointPOSTDebit = "/api/v1/wallets/:walletid/debit"
	walletParam       = "walletid"
)

var (
	balance decimal.Decimal
	err     error
)

// Balance stores the incoming balance to debit or credit
type Balance struct {
	Balance string `json:"balance" binding:"required"`
}

// NewGetBalanceHandler returns a handler for the GetBalance endpoint
func NewGetBalanceHandler(cs controllers.Store) func(c *gin.Context) {
	return func(c *gin.Context) {
		wid := c.Param(walletParam)
		balance, err = cs.GetBalance(wid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"balance": balance,
		})
	}
}

// NewPostCreditHandler returns a handler for the PostCredit endpoint
func NewPostCreditHandler(cs controllers.Store) func(c *gin.Context) {
	return func(c *gin.Context) {
		var json Balance
		if err = c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		wid := c.Param("walletid")
		_, err = cs.Credit(wid, json.Balance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}

// NewPostDebitHandler returns a handler for the PostDebit endpoint
func NewPostDebitHandler(cs controllers.Store) func(c *gin.Context) {
	return func(c *gin.Context) {
		var json Balance
		if err = c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		wid := c.Param("walletid")
		_, err = cs.Debit(wid, json.Balance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
