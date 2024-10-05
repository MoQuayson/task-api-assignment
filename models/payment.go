package models

import (
	"context"
	"database/sql"
	"errors"
	"github.com/moquayson/task-api-assignment/requests"
	"github.com/moquayson/task-api-assignment/utils"
	"log"
	"time"
)

type Payment struct {
	ID                  int    `json:"id,omitempty"`
	UserID              int    `json:"user_id"`
	SenderMobileNo      string `json:"sender_mobile_no"`
	BeneficiaryMobileNo string `json:"beneficiary_mobile_no"`
	Amount              int    `json:"amount"`
	TransactionID       string `json:"transaction_id"`
	Status              string `json:"status"`
	//CreatedAt           time.Time `json:"created_at"`
	//UpdatedAt           time.Time `json:"updated_at"`
}

func NewPaymentWithMakePaymentRequest(req *requests.MakePaymentRequest, transactionId *string) *Payment {
	return &Payment{
		UserID:              req.UserID,
		SenderMobileNo:      req.SenderMobileNo,
		BeneficiaryMobileNo: req.BeneficiaryMobileNo,
		Amount:              req.Amount,
		TransactionID:       *transactionId,
		Status:              string(utils.PaymentStatus_Pending),
	}
}

func InsertIntoPayment(req *requests.MakePaymentRequest, transactionId *string, db *sql.DB) (*Payment, error) {
	dataChan, errChan := make(chan *Payment, 1), make(chan error, 1)

	go func(*requests.MakePaymentRequest, *sql.DB, chan *Payment, chan error) {
		ctx, cancel := context.WithTimeout(context.Background(), utils.DatabaseTimeout)

		defer cancel()
		query := "insert into payments (user_id,sender_mobile_no,beneficiary_mobile_no,amount,transaction_id,status) values (?,?,?,?,?,?)"

		//get email and password
		result, err := db.ExecContext(ctx, query, req.UserID, req.SenderMobileNo, req.BeneficiaryMobileNo, req.Amount, transactionId, utils.PaymentStatus_Pending)
		if err != nil {
			log.Printf("InsertIntoPayment err: %v", err)
			dataChan <- nil
			errChan <- err
			return
		}

		if rows, _ := result.RowsAffected(); rows != 1 {
			log.Printf("InsertIntoPayment err: %v", "no rows affected")
			dataChan <- nil
			errChan <- errors.New("no rows affected")
			return
		}

		//check password
		dataChan <- NewPaymentWithMakePaymentRequest(req, transactionId)
		errChan <- nil
		return
	}(req, db, dataChan, errChan)

	return <-dataChan, <-errChan
}

func UpdatePaymentStatusByTransactionId(transactionId *string, status utils.PaymentStatus, db *sql.DB) error {
	errChan := make(chan error, 1)

	go func(*string, utils.PaymentStatus, *sql.DB, chan error) {
		log.Printf("random status update: %s", string(status))
		ctx, cancel := context.WithTimeout(context.Background(), utils.DatabaseTimeout)

		defer cancel()
		query := "update payments set status = ?, updated_at = ? where transaction_id = ?"
		//query := fmt.Sprintf("update payments set status = '%s' and updated_at = '%s' where transaction_id = '%s'",string(status), time.Now().Format(time.DateTime), transactionId)

		//get email and password
		result, err := db.ExecContext(ctx, query, string(status), time.Now().Format(time.DateTime), transactionId)
		if err != nil {
			log.Printf("UpdatePaymentStatusByTransactionId err: %v", err)
			errChan <- err
			return
		}

		if rows, _ := result.RowsAffected(); rows != 1 {
			log.Printf("UpdatePaymentStatusByTransactionId err: %v", "no rows affected")
			errChan <- errors.New("no rows affected")
			return
		}

		errChan <- nil
		return
	}(transactionId, status, db, errChan)

	return <-errChan
}
func GetPaymentByTransactionId(transactionId *string, db *sql.DB) (*Payment, error) {
	dataChan, errChan := make(chan *Payment, 1), make(chan error, 1)

	go func(*string, *sql.DB, chan *Payment, chan error) {
		ctx, cancel := context.WithTimeout(context.Background(), utils.DatabaseTimeout)

		defer cancel()
		query := "select user_id,sender_mobile_no,beneficiary_mobile_no,amount,transaction_id,status from payments where transaction_id = ?"

		payment := &Payment{}
		//get email and password
		err := db.QueryRowContext(ctx, query, transactionId).Scan(&payment.UserID, &payment.SenderMobileNo, &payment.BeneficiaryMobileNo, &payment.Amount, &payment.TransactionID, &payment.Status)
		if err != nil {
			log.Printf("GetPaymentByTransactionId err: %v", err)
			dataChan <- nil
			errChan <- err
			return
		}

		if len(payment.TransactionID) == 0 {
			log.Printf("GetPaymentByTransactionId err: %v", "payment not found")
			dataChan <- nil
			errChan <- errors.New("payment not found")
			return
		}

		//check password
		dataChan <- payment
		errChan <- nil
		return
	}(transactionId, db, dataChan, errChan)

	return <-dataChan, <-errChan
}
