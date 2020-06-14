package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal/controllers"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/micuffaro/wallet/internal/views"
)

const (
	dbuser        = "root"
	dbpassword    = ""
	dbname        = "wallet"
	cacheserver   = "localhost:6379"
	cachepassword = ""
	cachedb       = 0
)

var (
	err     error
	service *models.DBService
	dbInfo  = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8&parseTime=True&loc=Local",
		dbuser, dbpassword, dbname,
	)
)

func main() {
	// Initiate DB connection
	service, err = models.NewDBService(dbInfo)
	if err != nil {
		panic(err)
	}
	defer service.Close()
	service.AutoMigrate() // Attempt to initialize with wallet table

	// Wallet controller that directly checks underlying storage
	// walletC := controllers.NewWalletController(service)

	// Wallet controller that uses a cache service on top of underlying storage
	walletC := &controllers.CacheStore{
		controllers.NewCacheService(cacheserver, cachepassword, cachedb),
		controllers.NewWalletController(service),
	}

	r := gin.Default()
	// GET the wallet balance
	r.GET(views.EndpointGETBalance, views.NewGetBalanceHandler(walletC))
	// POST credit to the wallet balance
	r.POST(views.EndpointPOSTCredit, views.NewPostCreditHandler(walletC))
	// POST debit to the wallet balance
	r.POST(views.EndpointPOSTDebit, views.NewPostDebitHandler(walletC))
	r.Run()
}
