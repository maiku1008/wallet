package wallet

import (
	"errors"
	"testing"
)

var testCreditCases = []struct {
	description string
	balance     float32
	amount      float32
	newBalance	float32
	err         error
}{
	{
		"Valid credit operation",
		0.0,
		100.0,
		100.0,
		nil,
	},
	{
		"Amount to credit is negative",
		0.0,
		-10.0,
		0.0,
		errors.New("amount to credit cannot be negative: -10.00"),
	},
}

func TestCredit(t *testing.T) {
	for _, test := range testCreditCases {
		w := &Wallet{Balance: test.balance}
		err := w.Credit(test.amount)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("Credit(%.2f): %s\n\t Expected: %t\n\t Got: %t", test.amount, test.description, test.err, err)
		}
		if w.Balance != test.newBalance {
			t.Fatalf("Credit(%.2f): %s\n\t Expected: %.2ft\n\t Got: %.2f", test.amount, test.description, test.newBalance, w.Balance)
		}
	}
}

var testDebitCases = []struct {
	description string
	balance     float32
	amount      float32
	newBalance	float32
	err         error
}{
	{
		"Valid debit operation",
		100.0,
		100.0,
		0.0,
		nil,
	},
	{
		"Balance doesn't go negative",
		50.0,
		60.0,
		0.0,
		nil,
	},
	{
		"Amount to debit is negative",
		0.0,
		-10.0,
		0.0,
		errors.New("amount to debit cannot be negative: -10.00"),
	},
}

func TestDebit(t *testing.T) {
	for _, test := range testDebitCases {
		w := &Wallet{Balance: test.balance}
		err := w.Debit(test.amount)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("Debit(%.2f): %s\n\t Expected: %t\n\t Got: %t", test.amount, test.description, test.err, err)
		}
		if w.Balance != test.newBalance {
			t.Fatalf("Debit(%.2f): %s\n\t Expected: %.2ft\n\t Got: %.2f", test.amount, test.description, test.newBalance, w.Balance)
		}
	}
}
