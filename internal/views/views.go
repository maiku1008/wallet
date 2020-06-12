package views

const (
	// Balance endpoint
	EndpointGETBalance = "/api/v1/wallets/:walletid/balance"
	// Credit endpoint
	EndpointPOSTCredit = "/api/v1/wallets/:walletid/credit"
	// Debit endpoint
	EndpointPOSTDebit = "/api/v1/wallets/:walletid/debit"
)

// Balance stores the balance to debit or credit
type Balance struct {
	Balance string `json:"balance" binding:"required"`
}
