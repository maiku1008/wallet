package main

import (
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/micuffaro/wallet/internal/views"
	"net/http"
	"fmt"
)

const (
	user     = "root"
	password = "pippo123"
	dbname   = "wallet"
)

func main() {
	// Create a DB connection string and then use it to
	// create our model services.
	mysqlInfo := fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		user, password, dbname,
	)
	service, err := models.NewService(mysqlInfo)
	if err != nil {
		panic(err)
	}
	defer service.Close()
	service.AutoMigrate() // Initialize with wallet table

	r := gin.Default()

	// Gets the wallet balance
	r.GET(views.EndpointGETBalance, func(c *gin.Context) {
		wid := c.Param("walletid")
		w, err := service.Wallet.Get(wid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"balance": w.Balance,
		})
	})

	// Credit
	r.POST(views.EndpointPOSTCredit, func(c *gin.Context) {
		var json views.Balance
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		wid := c.Param("walletid")
		w, err := service.Wallet.Get(wid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ww, _ := wallet.New(w.WID, w.Balance.String())
		if err := ww.Credit(json.Balance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		w.Balance = ww.Balance
		err = service.Wallet.Update(w)
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
		if err := c.ShouldBindJSON(&json); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		wid := c.Param("walletid")
		w, err := service.Wallet.Get(wid)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ww, _ := wallet.New(w.WID, w.Balance.String())
		if err := ww.Debit(json.Balance); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		w.Balance = ww.Balance
		err = service.Wallet.Update(w)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	r.Run()
}
