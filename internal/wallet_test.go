package wallet

import (
	"testing"
)

const id = "123"

func TestNew(t *testing.T) {
	for _, test := range testNewCases {
		_, err := New(id, test.balance)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("New(%s, %s): %s\n\t Expected: %t\n\t Got: %t", id, test.balance, test.description, test.err, err)
		}
	}
}

func TestCredit(t *testing.T) {
	for _, test := range testCreditCases {
		w, _ := New(id, test.balance)
		err := w.Credit(test.amount)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("Credit(%s): %s\n\t Expected: %t\n\t Got: %t", test.amount, test.description, test.err, err)
		}

		if w.PrintBalance() != test.newBalance {
			t.Fatalf("Credit(%s): %s\n\t Expected: %s\n\t Got: %s", test.amount, test.description, test.newBalance, w.PrintBalance())
		}
	}
}

func TestDebit(t *testing.T) {
	for _, test := range testDebitCases {
		w, _ := New(id, test.balance)
		err := w.Debit(test.amount)
		if err != nil && (err.Error() != test.err.Error()) {
			t.Fatalf("Debit(%s): %s\n\t Expected: %t\n\t Got: %t", test.amount, test.description, test.err, err)
		}
		if w.PrintBalance() != test.newBalance {
			t.Fatalf("Debit(%s): %s\n\t Expected: %st\n\t Got: %s", test.amount, test.description, test.newBalance, w.PrintBalance())
		}
	}
}
