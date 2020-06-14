// Package controllers exposes a controller that connects business rules and models
package controllers

import (
	wallet "github.com/micuffaro/wallet/internal"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/shopspring/decimal"
)

var (
	err  error
	mw   *models.Wallet
	ww   *wallet.Wallet
	amnt decimal.Decimal
	none = decimal.NewFromInt(0)
)

// NewWalletController initializes a new wallet controller
func NewWalletController(service *models.DBService) *WalletController {
	return &WalletController{
		service,
	}
}

// WalletController is a controller for wallet operations
type WalletController struct {
	*models.DBService
}

var _ Store = &WalletController{}

// GetBalance fetches from storage the balance of the object identified by wid
// and returns it
func (wc *WalletController) GetBalance(wid string) (decimal.Decimal, error) {
	mw, err = wc.Wallet.Get(wid)
	if err != nil {
		return none, err
	}

	return mw.Balance, nil
}

// Credit fetches from storage the object identified by wid and credits an amount
func (wc *WalletController) Credit(wid, amount string) (decimal.Decimal, error) {
	mw, err = wc.Wallet.Get(wid)
	if err != nil {
		return none, err
	}

	ww, err = wallet.New(mw.Balance)
	if err != nil {
		return none, err
	}

	amnt, err = decimal.NewFromString(amount)
	if err != nil {
		return none, err
	}

	if err = ww.Credit(amnt); err != nil {
		return none, err
	}

	mw.Balance = ww.Balance
	err = wc.Wallet.Update(mw)
	if err != nil {
		return none, err
	}

	return mw.Balance, nil
}

// Debit fetches from storage the object identified by wid and debits an amount
func (wc *WalletController) Debit(wid, amount string) (decimal.Decimal, error) {
	mw, err = wc.Wallet.Get(wid)
	if err != nil {
		return none, err
	}

	ww, err = wallet.New(mw.Balance)
	if err != nil {
		return none, err
	}

	amnt, err = decimal.NewFromString(amount)
	if err != nil {
		return none, err
	}

	if err = ww.Debit(amnt); err != nil {
		return none, err
	}

	mw.Balance = ww.Balance
	err = wc.Wallet.Update(mw)
	if err != nil {
		return none, err
	}

	return mw.Balance, nil
}
