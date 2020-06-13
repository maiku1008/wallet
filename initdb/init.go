package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/shopspring/decimal"
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

	b1, _ := decimal.NewFromString("0.01")
	w1 := &models.Wallet{
		gorm.Model{},
		"123",
		b1,
	}

	b2, _ := decimal.NewFromString("12.30")
	w2 := &models.Wallet{
		gorm.Model{},
		"456",
		b2,
	}

	// Create
	err = service.Wallet.Create(w1)
	if err != nil {
		panic(err)
	}

	// Create
	err = service.Wallet.Create(w2)
	if err != nil {
		panic(err)
	}
}
