package database

import (
	"database/sql"
)

func MigrateTables(db *sql.DB) error {
	query := `

CREATE TABLE IF NOT EXISTS users (
    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    full_name varchar(255) NOT NULL,
    email varchar(100) NOT NULL UNIQUE,
    password varchar(255) not null,
    created_at datetime NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NULL
);

CREATE TABLE IF NOT EXISTS payments (
    id integer NOT NULL PRIMARY KEY AUTOINCREMENT,
    user_id int NOT NULL,
    sender_mobile_no varchar(100) NOT NULL,
    beneficiary_mobile_no varchar(100) NOT NULL,
    amount decimal NOT NULL,
    transaction_id varchar(100) NOT NULL,
    status varchar(100) NOT NULL,
    created_at datetime NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at datetime NULL,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);
`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}
