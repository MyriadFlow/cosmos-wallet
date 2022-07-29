// Package dbconinit provides method to Init database
package dbconinit

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/env"
	"github.com/MyriadFlow/cosmos-wallet/custodial/pkg/store"
	"github.com/MyriadFlow/cosmos-wallet/helpers/logo"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() {
	var (
		host     = env.MustGetEnv("DB_HOST")
		username = env.MustGetEnv("DB_USERNAME")
		password = env.MustGetEnv("DB_PASSWORD")
		dbname   = env.MustGetEnv("DB_NAME")
		port     = env.MustGetEnv("DB_PORT")
	)

	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable port=%s",
		host, username, password, dbname, port)

	var err error
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dns,
	}))
	if err != nil {
		logo.Fatal("failed to connect database", err)
	}

	// Store database in global store
	store.DB = db

	// Get underlying sql database to ping it
	sqlDb, err := db.DB()
	if err != nil {
		logo.Fatal("failed to ping database", err)
	}

	// If ping fails then log error and exit
	if err = sqlDb.Ping(); err != nil {
		logo.Fatal("failed to ping database", err)
	}
}
