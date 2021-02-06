package services

import (
	"../domains"
	"../utils"
)

func GetUser(userId int64) (*domains.User, *utils.ApplicationError)  {
	return domains.GetUser(userId)
}