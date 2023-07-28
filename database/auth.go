package database

import (
	"fmt"
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
	"golang.org/x/crypto/bcrypt"
)

/*
Stores user details in db
Email must be unique
Username is not assigned
*/
func (s *PostgresStorage) CreateAccount(acc *t.NewAccount) *u.ApiError {

	// begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return u.NewError(err, http.StatusBadRequest)
	}
	ch := make(chan *u.ApiError)

	go func() {

		row, err := tx.Query(`SELECT count(*) FROM accounts WHERE email = $1`, acc.Email)
		if err != nil {
			tx.Rollback()
			ch <- u.NewError(err, http.StatusBadRequest)
		}

		count := 0
		for row.Next() {

			err := row.Scan(
				&count,
			)

			if err != nil {
				tx.Rollback()
				ch <- u.NewError(err, http.StatusBadRequest)
			}

		}

		if count > 0 {
			tx.Rollback()
			ch <- u.NewError(fmt.Errorf("email already in use"), http.StatusConflict)
		}
		ch <- nil
	}()

	info := <-ch

	if info != nil {
		tx.Rollback()
		return info
	}

	query :=
		`insert into accounts
(id,first_name, last_name,email,created_at,password)
values($1,$2,$3,$4,$5,$6)`

	_, err = tx.Exec(
		query,
		acc.ID,
		acc.Firstname,
		acc.Lastname,
		acc.Email,
		acc.CreatedAt,
		acc.Password,
	)

	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusConflict)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return u.NewError(err, http.StatusConflict)
	}

	return nil
}

func (s *PostgresStorage) GetAccount(req *t.SigninRequest) (*t.Account, *u.ApiError) {

	// get email and password
	row, err := s.db.Query(
		`SELECT * FROM accounts_view WHERE email = $1`, req.Email)

	if err != nil {
		return nil, u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return nil, error
		}

	}

	if acc.Password == "" {
		return nil, u.NewError(fmt.Errorf("invalid email or password"), http.StatusUnauthorized)
	}

	// validating password
	if err := bcrypt.CompareHashAndPassword([]byte(string(acc.Password)), []byte(req.Pasword)); err != nil {

		return nil, u.NewError(fmt.Errorf("invalid email or password"), http.StatusUnauthorized)
	}

	return acc, nil
}

func (s *PostgresStorage) AltCreateAccount(acc *t.NewAccount) *u.ApiError {
	// begin transaction

	err := s.CreateAccount(acc)

	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStorage) AltGetAccount(req *t.AltSigninRequest) (*t.Account, *u.ApiError) {

	// get email and password
	row, err := s.db.Query(
		`SELECT * FROM accounts_view WHERE email = $1`, req.Email)

	if err != nil {
		return nil, u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return nil, error
		}
	}

	if acc.Firstname == "" {
		return nil, u.NewError(fmt.Errorf("user not found"), http.StatusBadGateway)
	}

	return acc, nil

}
