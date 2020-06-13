package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal/controllers"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/micuffaro/wallet/internal/views"
	"github.com/shopspring/decimal"
	"net/http"
)

const (
	user     = "root"
	password = "pippo123"
	dbname   = "wallet"
)

var (
	err error
	service *models.Service
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	mysqlInfo := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, dbname,
	)
	service, err = models.NewService(mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer service.Close()
	_ = service.AutoMigrate() // Initialize with wallet table

	r := gin.Default()

	// Gets the wallet balance
	r.GET(views.EndpointGETBalance, func(c *gin.Context) {
		var balance decimal.Decimal
		wid := c.Param("walletid")
		balance, err = controllers.GetBalance(wid, service)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"balance": balance,
		})
	})

	// Credit
	r.POST(views.EndpointPOSTCredit, func(c *gin.Context) {
		var json views.Balance
		if err = c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		wid := c.Param("walletid")
		err = controllers.Credit(wid, json.Balance, service)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	// Debit
	r.POST(views.EndpointPOSTDebit, func(c *gin.Context) {
		var json views.Balance
		if err = c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		wid := c.Param("walletid")
		err = controllers.Debit(wid, json.Balance, service)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	_ = r.Run()
}
