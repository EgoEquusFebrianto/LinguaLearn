package seed

import (
	"database/sql"
	"fmt"
	"os"
)

func SeedMySQL(db *sql.DB) error {
	data, err := os.ReadFile("database/seeds/mysql/roles.sql")
	if err != nil {
		return fmt.Errorf("Read roles seed failed: %w", err)
	}

	if _, err := db.Exec(string(data)); err != nil {
		return fmt.Errorf("Execute roles seed: %w", err)
	}

	return nil
}