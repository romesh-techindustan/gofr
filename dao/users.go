package dao

import (
	"context"
	"example/models"
	"example/requests"
)

var ctx context.Context

func (dao *UserDAO) CreateUser(request requests.UserRequest) (*models.User, error) {
	if err := request.ValidateUserRequest(); err != nil {
		return nil, err
	}
	user := models.User{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Phone:     request.Phone,
		Gender:    request.Gender,
	}
	if err := dao.createUser(user); err != nil {
		return nil, err
	}
	return &user, nil
}

func (dao *UserDAO) createUser(user models.User) error {
	collection := dao.db.Collection("users")
	_, err := collection.InsertOne(ctx, user)
	return err
}
