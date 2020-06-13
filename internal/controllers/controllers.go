package controllers

import (
	wallet "github.com/micuffaro/wallet/internal"
	"github.com/micuffaro/wallet/internal/models"
	"github.com/shopspring/decimal"
)

var (
	err error
	mw   *models.Wallet
	ww  *wallet.Wallet
	amnt   decimal.Decimal
	z   = decimal.NewFromInt(0)
)

// GetBalance fetches from storage the balance of the object identified by wid
// and returns it
func GetBalance(wid string, sv *models.Service) (decimal.Decimal, error) {
	mw, err = sv.Wallet.Get(wid)
	if err != nil {
		return z, err
	}

	return mw.Balance, nil
}

// Credit fetches from storage the object identified by wid and credits an amount
func Credit(wid, amount string, sv *models.Service) error {
	mw, err = sv.Wallet.Get(wid)
	if err != nil {
		return err
	}

	ww, err = wallet.New(mw.Balance)
	if err != nil {
		return err
	}

	amnt, err = decimal.NewFromString(amount)
	if err != nil {
		return err
	}

	if err = ww.Credit(amnt); err != nil {
		return err
	}

	mw.Balance = ww.Balance
	err = sv.Wallet.Update(mw)
	if err != nil {
		return err
	}

	return nil
}

// Debit fetches from storage the object identified by wid and debits an amount
func Debit(wid, amount string, sv *models.Service) error {
	mw, err = sv.Wallet.Get(wid)
	if err != nil {
		return err
	}

	ww, err = wallet.New(mw.Balance)
	if err != nil {
		return err
	}

	amnt, err = decimal.NewFromString(amount)
	if err != nil {
		return err
	}

	if err = ww.Debit(amnt); err != nil {
		return err
	}

	mw.Balance = ww.Balance
	err = sv.Wallet.Update(mw)
	if err != nil {
		return err
	}

	return nil
}
