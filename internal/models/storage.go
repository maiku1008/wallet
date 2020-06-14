// Package models exposes a set of models
package models

import (
	"errors"
	"github.com/jinzhu/gorm"
	// Mysql dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shopspring/decimal"
)

var (
	// errNotFound is returned when a resource cannot be found
	// in the database.
	errNotFound = errors.New("models: resource not found")
)

// Wallet represents a wallet object in our database.
type Wallet struct {
	gorm.Model
	WID     string          `gorm:"not null;unique_index"`
	Balance decimal.Decimal `sql:"type:decimal(20,8);"`
}

// walletGorm represents our database interaction layer
type walletGorm struct {
	db *gorm.DB
}

// Get query the db for an object with the given wid
func (wg *walletGorm) Get(wid string) (*Wallet, error) {
	var wallet Wallet
	db := wg.db.Where("w_id = ?", wid)
	if err := first(db, &wallet); err != nil {
		return nil, err
	}
	return &wallet, nil
}

// Create creates a wallet
// func (wg *walletGorm) Create(wallet *Wallet) error {
// 	return wg.db.Create(wallet).Error
// }

// Update updates a wallet
func (wg *walletGorm) Update(wallet *Wallet) error {
	return wg.db.Save(wallet).Error
}

// Delete deletes a wallet
// func (wg *walletGorm) Delete(wallet *Wallet) error {
// 	return wg.db.Delete(wallet).Error
// }

// first is a helper function that will query using the provided gorm.DB.
// it will get the first item returned and place it into a destination dst.
// If nothing is found in the query, it will return errNotFound
func first(db *gorm.DB, dst interface{}) error {
	err := db.First(dst).Error
	if err == gorm.ErrRecordNotFound {
		return errNotFound
	}
	return err
}
