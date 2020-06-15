// Package views exposes resources used for presenting data to the end user
package views

import (
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal/controllers"
	"github.com/shopspring/decimal"
	"github.com/sirupsen/logrus"
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

// NewHandlers returns a new Handlers object
func NewHandlers(logger *logrus.Logger, cs controllers.Store) *Handlers {
	return &Handlers{
		logger,
		cs,
	}
}

// Handlers contains resources to pass to Handlers
type Handlers struct {
	logger *logrus.Logger
	controllers.Store
}

// NewGetBalanceHandler returns a handler for the GetBalance endpoint
func (h *Handlers) NewGetBalanceHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		wid := c.Param(walletParam)

		// Log the request
		h.logger.WithFields(logrus.Fields{
			"method": "GET",
			"wid":    wid,
		}).Info("GET Balance request")

		balance, err = h.GetBalance(wid)
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
func (h *Handlers) NewPostCreditHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		wid := c.Param(walletParam)
		var json Balance
		if err = c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the request
		h.logger.WithFields(logrus.Fields{
			"method": "POST",
			"wid":    wid,
			"amount": json.Balance,
		}).Info("POST Credit request")

		_, err = h.Credit(wid, json.Balance)
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
func (h *Handlers) NewPostDebitHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		wid := c.Param(walletParam)
		var json Balance
		if err = c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Log the request
		h.logger.WithFields(logrus.Fields{
			"method": "POST",
			"wid":    wid,
			"amount": json.Balance,
		}).Info("POST Debit request")

		_, err = h.Debit(wid, json.Balance)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	}
}
