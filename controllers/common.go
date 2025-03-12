package controllers

import (
	"example/dao"
)

type UserController struct {
	userDAO *dao.UserDAO
}

// NewUserRepository creates a new instance of UserRepository
func NewUserDAO(userDAO *dao.UserDAO) *UserController {
	return &UserController{userDAO: userDAO}
}
