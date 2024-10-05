package handlers

import (
	"encoding/json"
	"github.com/moquayson/task-api-assignment/configs"
	"github.com/moquayson/task-api-assignment/models"
	"github.com/moquayson/task-api-assignment/requests"
	"github.com/moquayson/task-api-assignment/utils"
	"net/http"
)

func LoginUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == string(utils.ActionVerb_Post) {
		request := requests.LoginRequest{}
		// Parse the JSON request body
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if len(request.Email) == 0 {
			http.Error(w, "email cannot be empty", http.StatusBadRequest)
			return
		}

		if len(request.Password) == 0 {
			http.Error(w, "password cannot be empty", http.StatusBadRequest)
			return
		}

		//check credentials
		isValidUser, err := models.AuthenticateUser(&request, configs.DBContext)

		if err != nil {
			http.Error(w, "something went wrong. try again later", http.StatusInternalServerError)
			return
		}

		if !isValidUser {
			http.Error(w, "invalid credentials", http.StatusUnauthorized)
			return
		}

		user, err := models.GetUserByEmail(&request.Email, configs.DBContext)
		if err != nil {
			http.Error(w, "something went wrong. try again later", http.StatusInternalServerError)
			return
		}

		//bearer token
		token := utils.GenerateToken()
		models.BearerTokens[token] = token
		user.Token = token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.ApiResponse{
			Code:    200,
			Message: "valid credentials",
			Data:    user,
		})
		return
	} else {
		http.Error(w, "method must be a POST verb", http.StatusMethodNotAllowed)
		return
	}

}

func GenerateAccessTokenHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == string(utils.ActionVerb_Post) {

		//bearer token
		token := utils.GenerateToken()
		models.BearerTokens[token] = token
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(&models.ApiResponse{
			Code:    200,
			Message: "access token refreshed",
			Data:    &models.AccessToken{Token: token},
		})
		return
	} else {
		http.Error(w, "method must be a POST verb", http.StatusMethodNotAllowed)
		return
	}

}
