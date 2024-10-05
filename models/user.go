package models

import (
	"context"
	"database/sql"
	"github.com/moquayson/task-api-assignment/utils"
	"log"
)

type User struct {
	Id       int    `json:"id" db:"id"`
	FullName string `json:"full_name" db:"full_name"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password,omitempty" db:"password"`
	Token    string `json:"token,omitempty"`
}

func GetUserByEmail(email *string, db *sql.DB) (*User, error) {
	dataChan, errChan := make(chan *User, 1), make(chan error, 1)

	go func(*string, *sql.DB, chan *User, chan error) {
		ctx, cancel := context.WithTimeout(context.Background(), utils.DatabaseTimeout)

		defer cancel()
		query := "select id,full_name,email from users where email = ?"
		user := &User{}

		//get email and password
		err := db.QueryRowContext(ctx, query, email).Scan(&user.Id, &user.FullName, &user.Email)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("GetUserByEmail err: %v", err)
			dataChan <- nil
			errChan <- err
			return
		}

		if user == nil || user.Id == 0 {
			dataChan <- nil
			errChan <- nil
			return
		}

		//check password
		dataChan <- user
		errChan <- nil
		return
	}(email, db, dataChan, errChan)

	return <-dataChan, <-errChan
}
