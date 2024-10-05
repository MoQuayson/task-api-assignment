package handlers

import (
	"encoding/json"
	"github.com/moquayson/task-api-assignment/configs"
	"github.com/moquayson/task-api-assignment/models"
	"github.com/moquayson/task-api-assignment/requests"
	"github.com/moquayson/task-api-assignment/utils"
	"log"
	"net/http"
	"strings"
	"time"
)

func MakePaymentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == string(utils.ActionVerb_Post) {
		//simulating request
		sender := utils.GenerateMobileNumber()
		beneficiary := utils.GenerateMobileNumber()
		req := requests.NewMakePaymentRequest(1, &sender, &beneficiary, utils.GenerateRandomAmount())

		transactionId := strings.ToUpper(utils.GenerateToken())

		go func(string) {
			//insert payment
			_, err := models.InsertIntoPayment(req, &transactionId, configs.DBContext)
			if err != nil {
				log.Printf("InsertIntoPayment err: %v", err)
				return
			}

			//randomly sleep thread for sometime
			sleepDuration := utils.GenerateTimeSleepDuration() * time.Second
			log.Printf("time.Sleep Duration %v", sleepDuration)
			time.Sleep(sleepDuration)

			//update status of payment to SUCCESS or declined
			if err = models.UpdatePaymentStatusByTransactionId(&transactionId, utils.RandomizePaymentStatus(), configs.DBContext); err != nil {
				log.Printf("UpdatePaymentStatusByTransactionId err: %v", err)
				return
			}

		}(transactionId)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.ApiResponse{
			Code:    201,
			Message: "payment initiated",
			Data: &models.Payment{
				Amount:              req.Amount,
				SenderMobileNo:      sender,
				BeneficiaryMobileNo: beneficiary,
				Status:              string(utils.PaymentStatus_Pending),
				TransactionID:       transactionId,
			},
		})
	} else {
		http.Error(w, "method must be a POST verb", http.StatusMethodNotAllowed)
		return
	}

}

func GetPaymentStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == string(utils.ActionVerb_Get) {
		// Extract transaction_id from URL
		transactionId := r.URL.Path[len("/api/payments/status/"):]

		//get transaction id
		payment, err := models.GetPaymentByTransactionId(&transactionId, configs.DBContext)

		if err != nil {
			http.Error(w, "something went wrong. try again later", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.ApiResponse{
			Code:    200,
			Message: "payment status retrieved",
			Data:    payment,
		})
	} else {
		http.Error(w, "method must be a GET verb", http.StatusMethodNotAllowed)
		return
	}

}

func GetPaymentCompletedStatusHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == string(utils.ActionVerb_Get) {
		// Extract transaction_id from URL
		transactionId := r.URL.Path[len("/api/payments/completed/"):]

		//get transaction id
		payment, err := models.GetPaymentByTransactionId(&transactionId, configs.DBContext)

		if err != nil {
			http.Error(w, "something went wrong. try again later", http.StatusInternalServerError)
			return
		}

		if payment.Status == string(utils.PaymentStatus_Pending) {
			http.Error(w, "transaction still pending and not completed", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.ApiResponse{
			Code:    200,
			Message: "payment completed",
			Data:    payment,
		})
	} else {
		http.Error(w, "method must be a GET verb", http.StatusMethodNotAllowed)
		return
	}

}
