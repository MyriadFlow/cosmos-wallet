package dbconinit

import (
	"fmt"

	"github.com/MyriadFlow/cosmos-wallet/pkg/env"
	"github.com/MyriadFlow/cosmos-wallet/pkg/logo"
	"github.com/MyriadFlow/cosmos-wallet/pkg/store"
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

	store.DB = db

	sqlDb, err := db.DB()
	if err != nil {
		logo.Fatal("failed to ping database", err)
	}
	if err = sqlDb.Ping(); err != nil {
		logo.Fatal("failed to ping database", err)
	}
}
