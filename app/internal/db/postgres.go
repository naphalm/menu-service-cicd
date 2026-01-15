package db

import (
	"database/sql"
	"errors"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func Init() error {
	conn := os.Getenv("DATABASE_URL")
	if conn == "" {
		return errors.New("DATABASE_URL not set")
	}

	var err error
	DB, err = sql.Open("postgres", conn)
	if err != nil {
		return err
	}

	// Connection pool tuning (VERY IMPORTANT in k8s)
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(30 * time.Minute)

	if err = DB.Ping(); err != nil {
		return err
	}

	return nil
}

func Close() error {
	if DB != nil {
		return DB.Close()
	}
	return nil
}
