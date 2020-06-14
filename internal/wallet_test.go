package wallet

import (
	"github.com/shopspring/decimal"
	"testing"
)

func TestNew(t *testing.T) {
	for _, test := range testNewCases {
		balance, _ := decimal.NewFromString(test.balance)
		_, err := New(balance)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("New(%s): %s\n\t Expected: %t\n\t Got: %t", test.balance, test.description, test.err, err)
		}
	}
}

func TestCredit(t *testing.T) {
	for _, test := range testCreditCases {
		balance, _ := decimal.NewFromString(test.balance)
		w, _ := New(balance)
		amount, _ := decimal.NewFromString(test.amount)
		err := w.Credit(amount)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("Credit(%s): %s\n\t Expected: %t\n\t Got: %t", test.amount, test.description, test.err, err)
		}

		if w.printBalance() != test.newBalance {
			t.Fatalf("Credit(%s): %s\n\t Expected: %s\n\t Got: %s", test.amount, test.description, test.newBalance, w.printBalance())
		}
	}
}

func TestDebit(t *testing.T) {
	for _, test := range testDebitCases {
		balance, _ := decimal.NewFromString(test.balance)
		w, _ := New(balance)
		amount, _ := decimal.NewFromString(test.amount)
		err := w.Debit(amount)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("Debit(%s): %s\n\t Expected: %t\n\t Got: %t", test.amount, test.description, test.err, err)
		}
		if w.printBalance() != test.newBalance {
			t.Fatalf("Debit(%s): %s\n\t Expected: %st\n\t Got: %s", test.amount, test.description, test.newBalance, w.printBalance())
		}
	}
}
