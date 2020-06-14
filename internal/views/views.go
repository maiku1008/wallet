// Package views exports some resources used for setting up views
package views

const (
	// EndpointGETBalance is the balance endpoint
	EndpointGETBalance = "/api/v1/wallets/:walletid/balance"
	// EndpointPOSTCredit is the credit endpoint
	EndpointPOSTCredit = "/api/v1/wallets/:walletid/credit"
	// EndpointPOSTDebit is the debit endpoint
	EndpointPOSTDebit = "/api/v1/wallets/:walletid/debit"
)

// Balance stores the incoming balance to debit or credit
type Balance struct {
	Balance string `json:"balance" binding:"required"`
}
