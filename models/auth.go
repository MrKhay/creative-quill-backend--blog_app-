package models

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type ApiError struct {
	Error  error `json:"error"`
	Status int
}

func NewError(error error, code int) *ApiError {
	return &ApiError{
		Error:  error,
		Status: code,
	}
}

// Account represents a account object.
// swagger:model
type Account struct {
	ID          string    `json:"_id"`
	Firstname   string    `json:"first_name,omitempty"`
	Lastname    string    `json:"last_name,omitempty"`
	Displayname string    `json:"display_name,omitempty"`
	Username    string    `json:"username,omitempty"`
	Email       string    `json:"email,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
	Password    string    `json:"-"`
}

func NewAccount(firstname, lastname, emailadr, password string) (*Account, *ApiError) {

	encow, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, NewError(err, http.StatusBadRequest)
	}

	return &Account{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     emailadr,
		CreatedAt: time.Now().UTC(),
		Password:  string(encow),
	}, nil

}
