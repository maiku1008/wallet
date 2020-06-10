# wallet
Proof of concept of a wallet api for an online casino


JSON API in Go to get the balance and manage credit/debit operations on players wallets. 
For example, you might receive calls on your API to get the balance of the wallet with id 123, or to credit the wallet with id 456 by 10.00 €. 
The storage mechanism to use will be MySQL.
Below are the 3 endpoints to implement, as well as the business rules.
 
## Endpoints

- balance : retrieves the balance of a given wallet id  
GET `/api/v1/wallets/{wallet_id}/balance` 
- credit : credits money on a given wallet id  
POST `/ api/v1/wallets/{wallet_id}/credit`
- debit : debits money from a given wallet id  
POST `/ api/v1/wallets/{wallet_id}/debit`

## Business rules
- A wallet balance cannot go below 0.
- Amounts sent in the credit and debit operations cannot be negative.

## Bonus
- Cache the wallet balances in Redis, so that they can be fetched from cache
- Add auth endpoint and authentication verification
- Add unit tests for the business rules/logic
- Log the incoming requests

## Libraries to use
- HTTP : https://github.com/gin-gonic/gin
- MySQL : https://github.com/go-gorm/gorm
- Redis : https://github.com/go-redis/redis
- Numbers : https://github.com/shopspring/decimal
- Logger : https://github.com/sirupsen/logrus

## Notes
- No need to care about the currencies.
- No need to create wallets, they can be pre-populated in storage.
- Make sure to return some meaningful errors if an operation is not possible.
- A particular attention will be put on how the application is constructed, more
specifically how the web layer, repositories for storage, entities and business logic are structured.  
Ideally, this architecture should make the application testable, and not too dependent on implementation details (for example, which repository/storage mechanism we use etc...)
