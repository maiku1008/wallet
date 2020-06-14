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
	err     error
	service *models.DBService
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	mysqlInfo := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, dbname,
	)
	service, err = models.NewDBService(mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer service.Close()
	_ = service.AutoMigrate() // Initialize with wallet table

	// Wallet controller that directly checks underlying storage
	// walletC := controllers.NewWalletController(service)

	// Wallet controller that uses a cache service on top of underlying storage
	walletC := &models.CacheStore{
		models.NewCacheService(),
		controllers.NewWalletController(service),
	}

	r := gin.Default()
	// Gets the wallet balance
	r.GET(views.EndpointGETBalance, func(c *gin.Context) {
		var balance decimal.Decimal
		wid := c.Param("walletid")
		balance, err = walletC.GetBalance(wid)
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
		_, err = walletC.Credit(wid, json.Balance)
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
		_, err = walletC.Debit(wid, json.Balance)
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
