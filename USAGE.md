## Usage notes

### Getting started
- Create a Mysql database `wallet`
- Change connection strings for db in `main.go:14-16`
- Init database with some wallets. `go run initdb/init.go`
- Start application `go run cmd/wallet.go`
- Example api calls:

```
curl -v http://localhost:8080/api/v1/wallets/456/balance
```

```
curl -v -X POST http://localhost:8080/api/v1/wallets/456/credit -H 'content-type: application/json' -d '{ "balance": "3.3" }'
```

```
curl -v -X POST http://localhost:8080/api/v1/wallets/456/debit -H 'content-type: application/json' -d '{ "balance": "3.0" }'
```

### Cache
- Have redis with default setup
- Enable redis caching by commenting out the controller on `main.go:39`
and uncommenting the cachestore wrapper in `main.go:42-45`

### Business rules
- Business logic is located in `internal/wallet/wallet.go`, along with unit tests
