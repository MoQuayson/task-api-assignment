package main

import (
	"encoding/json"
	"fmt"
	"github.com/moquayson/task-api-assignment/handlers"
	"github.com/moquayson/task-api-assignment/middlewares"
	"github.com/moquayson/task-api-assignment/utils"
	"log"
	"net/http"
	"os"
)

func main() {

	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	http.HandleFunc("/", func(w http.ResponseWriter, request *http.Request) {
		json.NewEncoder(w).Encode("welcome world!")
	})
	http.HandleFunc(utils.LoginUrl, handlers.LoginUserHandler)
	http.HandleFunc(utils.AccessTokenUrl, handlers.GenerateAccessTokenHandler)
	http.HandleFunc(utils.MakePaymentUrl, middlewares.RequireAuthentication(handlers.MakePaymentHandler))
	http.HandleFunc(utils.GetPaymentStatusUrl, middlewares.RequireAuthentication(handlers.GetPaymentStatusHandler))
	http.HandleFunc(utils.GetPaymentCompletedStatusUrl, middlewares.RequireAuthentication(handlers.GetPaymentCompletedStatusHandler))

	log.Printf("Server started on :%s\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))

}
