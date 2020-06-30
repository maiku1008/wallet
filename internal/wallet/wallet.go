// Package wallet defines the business rules of our application
package wallet

import (
	"fmt"
	"github.com/shopspring/decimal"
)

// Fixed is the number of decimal points we want to display
const Fixed = 2

// New initializes a new wallet
func New(balance decimal.Decimal) (*Wallet, error) {
	if balance.IsNegative() {
		return nil, fmt.Errorf("initial wallet balance cannot be negative: %s", balance.StringFixed(Fixed))
	}
	return &Wallet{
		Balance: balance,
	}, nil
}

// Wallet represents a wallet with a current Balance
type Wallet struct {
	// Wallet Balance
	Balance decimal.Decimal
}

// Debit subtracts an amount from a wallet's balance
func (w *Wallet) Debit(amount decimal.Decimal) error {
	if amount.IsNegative() {
		return fmt.Errorf("amount to debit cannot be negative: %s", amount.StringFixed(Fixed))
	}
	total := w.Balance.Sub(amount)
	if total.IsNegative() {
		return fmt.Errorf("debit operation with amount %s returns negative balance", amount.StringFixed(Fixed))
	}
	w.Balance = total
	return nil
}

// Credit adds an amount to the wallet's balance
func (w *Wallet) Credit(amount decimal.Decimal) error {
	if amount.IsNegative() {
		return fmt.Errorf("amount to credit cannot be negative: %s", amount.StringFixed(Fixed))
	}
	w.Balance = w.Balance.Add(amount)
	return nil
}

// PrintBalance prints the objects balance in string form
func (w *Wallet) printBalance() string {
	return w.Balance.StringFixed(Fixed)
}
