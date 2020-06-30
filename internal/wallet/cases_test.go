package wallet

import "errors"

var testNewCases = []struct {
	description string
	balance     string
	err         error
}{
	{
		"Initial negative balance",
		"-1.00",
		errors.New("initial wallet balance cannot be negative: -1.00"),
	},
}

var testCreditCases = []struct {
	description string
	balance     string
	amount      string
	newBalance  string
	err         error
}{
	{
		"Valid credit operation",
		"0.00",
		"100.00",
		"100.00",
		nil,
	},
	{
		"Valid credit operation with cents",
		"0.08",
		"100.12",
		"100.20",
		nil,
	},
	{
		"Amount to credit is negative",
		"0.00",
		"-10.00",
		"0.00",
		errors.New("amount to credit cannot be negative: -10.00"),
	},
}

var testDebitCases = []struct {
	description string
	balance     string
	amount      string
	newBalance  string
	err         error
}{
	{
		"Valid debit operation",
		"100.00",
		"100.00",
		"0.00",
		nil,
	},
	{
		"Amount to debit is negative",
		"0.00",
		"-10.00",
		"0.00",
		errors.New("amount to debit cannot be negative: -10.00"),
	},
	{
		"Debit operation leads to negative balance",
		"5.00",
		"10.00",
		"5.00",
		errors.New("debit operation with amount 10.00 returns negative balance"),
	},
}
