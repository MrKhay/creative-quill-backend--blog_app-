package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrkhay/creative-quill-backend/database"
	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

const Success = "Success"

type Database struct {
	db database.Storage
}

func SetUpHandlers(s database.Storage) *Database {

	return &Database{
		db: s,
	}
}

// AccountResponce represents the response for the Signup endpoint.
// swagger:response AccountResponce
type AccountResonce struct {
	Account *t.NewAccount `json:"account"`
	Token   *string       `json:"token"`
}

type LoginResonce struct {
	Account *t.Account `json:"account"`
	Token   *string    `json:"token"`
}

// Signup returns account details with token.
//
// swagger:route POST /account/register account Signup
//
// Returns account details with token..
//
// Responses:
//
//	200: AccountResponce

func (d *Database) Register(w http.ResponseWriter, r *http.Request) *u.ApiError {

	acc := new(t.NewAccount)

	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	acc, error := t.NewAccountFunc(acc.Firstname, acc.Lastname, acc.Email, acc.Password)

	if error != nil {
		return error
	}
	error = d.db.CreateAccount(acc)

	if error != nil {
		return error
	}

	token, error := u.CreateJWT(&acc.Email)
	if error != nil {
		return error
	}
	acc.Password = ""
	res := AccountResonce{
		Account: acc,
		Token:   &token,
	}

	u.WriteJson(w, http.StatusCreated, res)
	return nil
}

func (d *Database) AltRegister(w http.ResponseWriter, r *http.Request) *u.ApiError {

	acc := new(t.NewAccount)

	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	acc, error := t.NewAccountFunc(acc.Firstname, acc.Lastname, acc.Email, acc.Password)

	if error != nil {
		return error
	}
	error = d.db.AltCreateAccount(acc)

	if error != nil {
		return error
	}

	token, error := u.CreateJWT(&acc.Email)
	if error != nil {
		return error
	}
	acc.Password = ""
	res := AccountResonce{
		Account: acc,
		Token:   &token,
	}

	u.WriteJson(w, http.StatusCreated, res)
	return nil
}

// Signin returns account details with token.
//
// swagger:route POST /account/login account Signin
//
// Returns account details with token..
//
// Responses:
//
//	200: AccountResponce

// @Summary Login user.
// @Description to sign in user.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /signin [post]
func (d *Database) Login(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.SigninRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	acc, error := d.db.GetAccount(req)

	if error != nil {
		return error
	}

	// creats new auth token
	token, error := u.CreateJWT(&req.Email)

	if err != nil {
		return error
	}

	res := LoginResonce{
		Account: acc,
		Token:   &token,
	}

	u.WriteJson(w, http.StatusAccepted, res)
	return nil
}

// @Summary AltLogin user.
// @Description to sign in user using google,apple etc.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "OK"
// @Router /altsignin [post]
func (d *Database) AltLogin(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.AltSigninRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	acc, error := d.db.AltGetAccount(req)

	if error != nil {
		return error
	}

	// creats new auth token
	token, error := u.CreateJWT(&req.Email)

	if err != nil {
		return error
	}

	res := LoginResonce{
		Account: acc,
		Token:   &token,
	}

	u.WriteJson(w, http.StatusAccepted, res)
	return nil
}
