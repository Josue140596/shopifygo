package db

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewConnection returns a connection to the database
func NewConnection(dsnDB string) *gorm.DB {

	db, err := gorm.Open(postgres.Open(dsnDB), &gorm.Config{})

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection DB: %v\n", err)
		os.Exit(1)
	}

	sqlDB, err := db.DB()

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to pool connection DB: %v\n", err)
		os.Exit(1)
	}
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}
