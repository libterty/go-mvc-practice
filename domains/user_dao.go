package domains

import (
	"fmt"
	"net/http"

	"../utils"
)

var (
	users = map[int64]*User{
		123: &User{
			Id:        1,
			FirstName: "lib",
			LastName:  "11",
			Email:     "lib@test.com",
		},
	}
)

func GetUser(userId int64) (*User, *utils.ApplicationError) {
	user := users[userId]
	if user == nil {
		return nil, &utils.ApplicationError{
			Message:    fmt.Sprintf("User was not found: %v", userId),
			StatusCode: http.StatusNotFound,
			Code:       "not_found",
		}
	}
	return user, nil
}
