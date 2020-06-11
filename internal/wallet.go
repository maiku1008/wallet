package wallet

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// New initializes a new wallet
func New(id, amount string) (*Wallet, error) {
	balance, err := decimal.NewFromString(amount)
	if err != nil {
		return nil, err
	}

	if balance.IsNegative() {
		return nil, fmt.Errorf("initial wallet balance cannot be negative: %s", balance.StringFixed(2))
	}

	return &Wallet{
		ID:      id,
		Balance: balance,
	}, nil
}

// Wallet represents a wallet with an ID and a current Balance
type Wallet struct {
	// Wallet ID
	ID string
	// Wallet Balance
	Balance decimal.Decimal
}

// Debit subtracts an amount from a wallet's balance
func (w *Wallet) Debit(amount string) error {
	amountDec, err := decimal.NewFromString(amount)
	if err != nil {
		return err
	}

	if amountDec.IsNegative() {
		return fmt.Errorf("amount to debit cannot be negative: %s", amountDec.StringFixed(2))
	}

	total := w.Balance.Sub(amountDec)
	if total.IsNegative() {
		w.Balance = decimal.NewFromInt(0)
		return nil
	}

	w.Balance = total
	return nil
}

// Credit adds an amount to the wallet's balance
func (w *Wallet) Credit(amount string) error {
	amountDec, err := decimal.NewFromString(amount)
	if err != nil {
		return err
	}

	if amountDec.IsNegative() {
		return fmt.Errorf("amount to credit cannot be negative: %s", amountDec.StringFixed(2))
	}
	w.Balance = w.Balance.Add(amountDec)
	return nil
}

func (w *Wallet) PrintBalance() string {
	return w.Balance.StringFixed(2)
}
