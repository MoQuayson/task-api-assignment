package models

import (
	"context"
	"database/sql"
	"github.com/moquayson/task-api-assignment/requests"
	"github.com/moquayson/task-api-assignment/utils"
	"log"
)

type AccessToken struct {
	Token string `json:"token"`
}

// AuthenticateUser checks credentials of the user
func AuthenticateUser(req *requests.LoginRequest, db *sql.DB) (bool, error) {
	dataChan, errChan := make(chan bool, 1), make(chan error, 1)

	go func(*requests.LoginRequest, *sql.DB, chan bool, chan error) {
		ctx, cancel := context.WithTimeout(context.Background(), utils.DatabaseTimeout)

		defer cancel()
		query := "select id,email,password from users where email = ?"
		user := &User{}

		//get email and password
		err := db.QueryRowContext(ctx, query, req.Email).Scan(&user.Id, &user.Email, &user.Password)
		if err != nil && err != sql.ErrNoRows {
			log.Printf("AuthenticateUser err:= %v", err)
			dataChan <- false
			errChan <- err
			return
		}

		if user == nil || user.Id == 0 {
			dataChan <- false
			errChan <- nil
			return
		}

		//check password
		dataChan <- isValidCredentials(req, user)
		errChan <- nil
		return
	}(req, db, dataChan, errChan)

	return <-dataChan, <-errChan
}

func isValidCredentials(req *requests.LoginRequest, user *User) bool {
	return req.Email == user.Email && req.Password == user.Password
}
