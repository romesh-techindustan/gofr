package requests

import "fmt"

type UserRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
}

func (request UserRequest) ValidateUserRequest() error {
	if request.FirstName == "" || request.LastName == "" {
		return fmt.Errorf("%s", "Please enter the firstName and lastName field.")
	}
	if request.Phone == "" {
		return fmt.Errorf("%s", "Phone number is a required field.")
	}
	return nil
}
