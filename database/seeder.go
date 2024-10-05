package database

import (
	"database/sql"
	"github.com/moquayson/task-api-assignment/models"
	"log"
)

func SeedAdminUser(db *sql.DB) error {
	email := "admin@example.com"
	user, err := models.GetUserByEmail(&email, db)
	if err != nil {
		log.Fatalf("SeedAdminUser err: %v", err)
		return err
	}

	if user == nil {
		query := `
				INSERT INTO users (full_name, email, password)
				VALUES ('Administrator','admin@example.com','password');
			
			`
		if _, err = db.Exec(query); err != nil {
			log.Fatalf("SeedAdminUser err: %v", err)
			return err
		}
	}

	return nil
}
