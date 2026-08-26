package mysql

import (
	"database/sql"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGorm(db *sql.DB) (*gorm.DB, error) {
	return gorm.Open(
		mysql.New(mysql.Config{
			Conn: db,
		}),
		&gorm.Config{},
	)
	
	// return gorm.Open(mysql.Open(dsn), &gorm.Config{})
}