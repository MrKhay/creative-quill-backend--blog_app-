package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	t "github.com/mrkhay/creative-quill-backend/models"
)

type Storage interface {
	CreateTables() *t.ApiError
	Signup(acc *t.Account) *t.ApiError
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, *t.ApiError) {

	connStr := os.Getenv("POSTGRES_URI")

	if connStr == "" {
		return nil, t.NewError(fmt.Errorf("you must set your 'POSTGRES_URI' environmental variable"), http.StatusBadGateway)
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, t.NewError(err, http.StatusBadGateway)
	}

	if err := db.Ping(); err != nil {
		return nil, t.NewError(err, http.StatusBadGateway)
	}

	return &PostgresStorage{
		db: db,
	}, nil

}

func (s *PostgresStorage) Init() *t.ApiError {

	return s.CreateTables()

}

func (s *PostgresStorage) CreateTables() *t.ApiError {

	querey := `CREATE TABLE IF NOT EXISTS accounts (
	id serial,
	first_name varchar(50),
	last_name varchar(50),
	display_name varchar(50),
	user_name varchar(20),
	email varchar(30) unique,
	created_at timestamp,
    password varchar(200)
  )`

	_, err := s.db.Exec(querey)

	if err != nil {
		return t.NewError(err, http.StatusConflict)
	}

	return nil
}

func (s *PostgresStorage) Signup(acc *t.Account) *t.ApiError {

	// begin transaction
	tx, err := s.db.Begin()
	if err != nil {
		return t.NewError(err, http.StatusBadRequest)
	}
	ch := make(chan *t.ApiError)

	go func() {

		row, err := tx.Query(`SELECT count(*) FROM accounts WHERE email = $1`, acc.Email)
		if err != nil {
			tx.Rollback()
			ch <- t.NewError(err, http.StatusBadRequest)
		}

		count := 0
		for row.Next() {

			err := row.Scan(
				&count,
			)

			if err != nil {
				tx.Rollback()
				ch <- t.NewError(err, http.StatusBadRequest)
			}

		}

		if count > 0 {
			tx.Rollback()
			ch <- t.NewError(fmt.Errorf("email already in use"), http.StatusConflict)
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
(first_name, last_name, display_name ,user_name,email, created_at ,password)
values($1,$2,$3,$4,$5,$6,$7)`

	_, err = tx.Exec(
		query,
		acc.Firstname,
		acc.Lastname,
		acc.Displayname,
		acc.Username,
		acc.Email,
		acc.CreatedAt,
		acc.Password,
	)

	if err != nil {
		tx.Rollback()
		return t.NewError(err, http.StatusConflict)
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return t.NewError(err, http.StatusConflict)
	}

	return nil
}

func scanIntoAccount(rows *sql.Rows) (*t.Account, *t.ApiError) {

	account := new(t.Account)
	err := rows.Scan(
		&account.ID,
		&account.Firstname,
		&account.Lastname,
		&account.Displayname,
		&account.Username,
		&account.Email,
		&account.CreatedAt,
		&account.Password,
	)

	return account, t.NewError(err, http.StatusConflict)

}
