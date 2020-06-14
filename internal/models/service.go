package models

import (
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

var (
	none = decimal.NewFromInt(0)
)

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
