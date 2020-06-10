package wallet

import "fmt"

// New initializes a new wallet
// func New(id string, balance float32) (*Wallet, error) {
// 	if balance < 0 {
// 		return nil, fmt.Errorf("initial wallet balance cannot be negative: %f", balance)
// 	}
// 	return &Wallet{
// 		ID:      id,
// 		Balance: balance,
// 	}, nil
// }

// Wallet represents a wallet with an ID and a current Balance
type Wallet struct {
	// Wallet ID
	ID string
	// Wallet Balance
	Balance float32
}

// Debit subtracts an amount from a wallet's balance
func (w *Wallet) Debit(amount float32) error {
	if amount < 0 {
		return fmt.Errorf("amount to debit cannot be negative: %.2f", amount)
	}

	total := w.Balance - amount
	if total < 0 {
		w.Balance = 0
		return nil
	}

	w.Balance = total
	return nil
}

// Credit adds an amount to the wallet's balance
func (w *Wallet) Credit(amount float32) error {
	if amount < 0 {
		return fmt.Errorf("amount to credit cannot be negative: %.2f", amount)
	}
	w.Balance += amount
	return nil
}

func (w *Wallet) String() string {
	return fmt.Sprintf("%.2f", w.Balance)
}
