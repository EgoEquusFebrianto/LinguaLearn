package mysql

import (
	"database/sql"
	"time"

	"github.com/EgoEquusFebrianto/LinguaLearn/internal/config"
	// _ "github.com/go-sql-driver/mysql"
)

func NewMySQL(cfg *config.MySQLConfig) (*sql.DB, error) {
	db, err := sql.Open("mysql", cfg.DSN())

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}