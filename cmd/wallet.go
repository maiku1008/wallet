package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/micuffaro/wallet/internal/controllers"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/micuffaro/wallet/internal/views"
	"github.com/sirupsen/logrus"
	"os"
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
	// Initialize logger
	log := logrus.New()
	log.SetOutput(os.Stdout)

	// Initialize DB connection
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

	// Initialize handlers
	handlers := views.NewHandlers(log, walletC)

	r := gin.Default()
	// GET the wallet balance
	r.GET(views.EndpointGETBalance, handlers.NewGetBalanceHandler())
	// POST credit to the wallet balance
	r.POST(views.EndpointPOSTCredit, handlers.NewPostCreditHandler())
	// POST debit to the wallet balance
	r.POST(views.EndpointPOSTDebit, handlers.NewPostDebitHandler())
	r.Run()
}
