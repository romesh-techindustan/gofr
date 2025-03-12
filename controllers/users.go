package controllers

import (
	"example/requests"
	"net/http"

	"gofr.dev/pkg/gofr"
)

func (uc *UserController) SignUp() func(ctx *gofr.Context) (any, error) {
	return func(ctx *gofr.Context) (any, error) {
		var request requests.UserRequest

		if err := ctx.Bind(&request); err != nil {
			return nil, err
		}

		user, err := uc.userDAO.CreateUser(request)
		if err != nil {
			http.Error(http.ResponseWriter, err.Error(), 400)
			return nil, gofr.NewHTTPError(400, err.Error())
		}
		return user, nil
	}
}
