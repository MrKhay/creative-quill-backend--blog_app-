package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
	"golang.org/x/crypto/bcrypt"
)

type SigninRequest struct {
	Email   string `json:"email"`
	Pasword string `json:"password"`
}

type AltSigninRequest struct {
	Email string `json:"email"`
}
type NewAccount struct {
	ID        uuid.UUID `json:"id,omitempty"`
	Firstname string    `json:"first_name,omitempty"`
	Lastname  string    `json:"last_name,omitempty"`
	Email     string    `json:"email,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	Password  string    `json:"password,omitempty"`
}

type HeaderPicUrl string
type ProfilePicUrl string
type WebsiteUrl string

// Signin represents a account object.
// swagger:model
// Account represents a account object.
// swagger:model
type Account struct {
	ID                  uuid.UUID     `json:"id,omitempty"`
	Firstname           string        `json:"first_name"`
	Lastname            string        `json:"last_name"`
	Displayname         string        `json:"display_name"`
	Username            string        `json:"username"`
	Email               string        `json:"email"`
	FollowersCount      int           `json:"followers_no"`
	AccFolloweringCount int           `json:"acc_following_no"`
	WebsiteUrl          WebsiteUrl    `json:"website_url"`
	Description         string        `json:"description"`
	ProfilePicUrl       ProfilePicUrl `json:"profilepic_url"`
	HeaderPicUrl        HeaderPicUrl  `json:"headerpic_url"`
	CreatedAt           time.Time     `json:"createdAt"`
	Password            string        `json:"password,omitempty"`
}

func NewAccountFunc(firstname, lastname, emailadr, password string) (*NewAccount, *u.ApiError) {
	encow, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}
	return &NewAccount{
		ID:        id,
		Firstname: firstname,
		Lastname:  lastname,
		Email:     emailadr,
		CreatedAt: time.Now().UTC(),
		Password:  string(encow),
	}, nil

}

// Handles signin using google apple
// and any other third party signin option
func NewAltAccount(firstname, lastname, emailadr string) (*Account, *u.ApiError) {

	return &Account{
		Firstname: firstname,
		Lastname:  lastname,
		Email:     emailadr,
		CreatedAt: time.Now().UTC(),
	}, nil

}
