package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"../services"
	"../utils"
)

func GetUser(res http.ResponseWriter, req *http.Request) {
	userIdParam := req.URL.Query().Get("user_id")
	userId, err := strconv.ParseInt(userIdParam, 10, 64)
	if err != nil {
		appErr := &utils.ApplicationError{
			Message:    "User id must be a number",
			StatusCode: http.StatusBadRequest,
			Code:       "bad_request",
		}
		errValue, _ := json.Marshal(appErr)
		res.WriteHeader(appErr.StatusCode)
		res.Write(errValue)
		return
	}

	user, apiErr := services.GetUser(userId)
	if apiErr != nil {
		appErr := &utils.ApplicationError{
			Message:    "User not found",
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
		errValue, _ := json.Marshal(appErr)
		res.WriteHeader(appErr.StatusCode)
		res.Write(errValue)
		return
	}
	// return user
	resValue, _ := json.Marshal(user)
	res.Write(resValue)
}
