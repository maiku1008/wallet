package models

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

var (
	ctx  = context.TODO()
	none = decimal.NewFromInt(0)
)

// NewServices return an object with all services we want
func NewServices(dbConnInfo string) (*Services, error) {
	dbsv, err := NewDBService(dbConnInfo)
	if err != nil {
		return nil, err
	}

	csv := NewCacheService()
	return &Services{
		dbsv,
		csv,
	}, nil
}

// Services has the services we need
type Services struct {
	*DBService
	*CacheService
}

// NewDBService handles the database connection
func NewDBService(dbConnInfo string) (*DBService, error) {
	db, err := gorm.Open("mysql", dbConnInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &DBService{
		Wallet: &walletGorm{db},
		db:     db,
	}, nil
}

// DBService represents the database connection service
type DBService struct {
	Wallet *walletGorm
	db     *gorm.DB
}

// Close closes the database connection
func (s *DBService) Close() error {
	return s.db.Close()
}

// AutoMigrate will attempt to automatically migrate all tables
func (s *DBService) AutoMigrate() error {
	return s.db.AutoMigrate(&Wallet{}).Error
}

// DestructiveReset drops all tables and rebuilds them
func (s *DBService) DestructiveReset() error {
	err := s.db.DropTableIfExists(&Wallet{}).Error
	if err != nil {
		return err
	}
	return s.AutoMigrate()
}

// NewCacheService returns a new cache service object
func NewCacheService() *CacheService {
	rdb := redis.NewClient(&redis.Options{
		// Move these
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	return &CacheService{
		Cache: rdb,
	}
}

// CacheService represents the cache connection service
type CacheService struct {
	Cache *redis.Client
}

// Store is an interface that implements functions
// to be used for cache and storage alike
type Store interface {
	GetBalance(wid string) (decimal.Decimal, error)
}

// DatabaseStore is a wrapper around DBService
// Most probably not needed
type DatabaseStore struct {
	DBService
}

// GetBalance is a wrapper around the DBServices's Get()
// A copy-paste of walletcontroller's method
// Most probably not needed
// func (dbs *DatabaseStore) GetBalance(wid string) (decimal.Decimal, error) {
// 	mw, err := dbs.Wallet.Get(wid)
// 	if err != nil {
// 		return decimal.NewFromInt(0), err
// 	}
//
// 	return mw.Balance, nil
// }

// Cachestore embeds
type CacheStore struct {
	CacheService
	DatabaseStore
	Store
}

// GetBalance attempts to fetch a wid's balance from cache, gets it from storage otherwise
func (chs *CacheStore) GetBalance(wid string) (decimal.Decimal, error) {
	fmt.Println("checking cache")
	// Check cache
	b, err := chs.Cache.Get(ctx, wid).Result()
	if err != nil {
		return none, err
	}

	// Not found in cache
	if err == redis.Nil {
		var mw *Wallet
		// Get from storage
		mw, err = chs.Wallet.Get(wid)
		if err != nil {
			return none, err
		}

		// Set wid and balance in cache
		err = chs.Cache.Set(ctx, wid, mw.Balance.StringFixed(2), 0).Err()
		if err != nil {
			return none, err
		}
		return mw.Balance, nil
	}

	// Found in cache
	var balance decimal.Decimal
	balance, err = decimal.NewFromString(b)

	return balance, nil
}
