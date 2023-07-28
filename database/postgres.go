package database

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	_ "github.com/lib/pq"
	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type Storage interface {
	CreateTables() *u.ApiError
	Atuh
	User
}

type Atuh interface {
	CreateAccount(acc *t.NewAccount) *u.ApiError
	GetAccount(req *t.SigninRequest) (*t.Account, *u.ApiError)
	AltCreateAccount(acc *t.NewAccount) *u.ApiError
	AltGetAccount(req *t.AltSigninRequest) (*t.Account, *u.ApiError)
}

type User interface {
	FollowUser(req *t.FollowRequest) *u.ApiError
	UnFollowUser(req *t.UnFollowRequest) *u.ApiError
	GetFollowerDetails(req *t.GetFollowersDetail) ([]*t.Account, *u.ApiError)
	GetAccFollowingDetails(req *t.GetAccFollowingDetails) ([]*t.Account, *u.ApiError)
	UpdateAccount(req *t.UpdateAccountReq) *u.ApiError
	UpdateFirstName(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateLastName(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateDisplayName(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateUsername(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateWebsiteUrl(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateProfilePicUrl(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateDescription(req *t.UpdateAccDetailsRequest) *u.ApiError
	UpdateHeaderPicUrl(req *t.UpdateAccDetailsRequest) *u.ApiError
}

type PostgresStorage struct {
	db *sql.DB
}

func NewPostgresStorage() (*PostgresStorage, *u.ApiError) {

	connStr := os.Getenv("POSTGRES_URI")

	if connStr == "" {
		return nil, u.NewError(fmt.Errorf("you must set your 'POSTGRES_URI' environmental variable"), http.StatusBadGateway)
	}

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, u.NewError(err, http.StatusBadGateway)
	}

	if err := db.Ping(); err != nil {
		return nil, u.NewError(err, http.StatusBadGateway)
	}

	return &PostgresStorage{
		db: db,
	}, nil

}

func (s *PostgresStorage) Init() *u.ApiError {

	return s.CreateTables()

}

// Creats Table If table does not exists.
func (s *PostgresStorage) CreateTables() *u.ApiError {

	querey := `CREATE TABLE IF NOT EXISTS accounts (
	id uuid primary key,
	first_name varchar(50),
	last_name varchar(50),
	display_name varchar(50) DEFAULT '',
	user_name varchar(20) DEFAULT '',
	email varchar(30),
	website_url text DEFAULT '',
	description text DEFAULT '',
	profilepic_url text DEFAULT '',
	headerpic_url text DEFAULT '',
	created_at timestamp,
    password text
    )`

	_, err := s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE TABLE IF NOT EXISTS following (
    follower_id uuid References accounts(id) ON DELETE CASCADE,
    following_id uuid References accounts(id) ON DELETE CASCADE,
	PRIMARY KEY (follower_id,following_id)
    )`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	querey = `CREATE OR REPLACE VIEW following_view AS SELECT id, (SELECT * FROM following WHERE follower_id = id) 
	         AS acc_following_no,(SELECT COUNT(*) FROM following WHERE following_id =id) AS followers_no 
	         FROM accounts;
              )`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}
	querey = `CREATE OR REPLACE VIEW accounts_view AS SELECT a.id,a.first_name,a.last_name,a.display_name,a.user_name,
	          a.email,a.website_url,a.description,a.profilepic_url,a.headerpic_url,a.created_at,fv.acc_following_no,
	          fv.followers_no,a.password FROM accounts a JOIN following_view AS fv ON a.id=fv.id`

	_, err = s.db.Exec(querey)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil
}

func scanIntoAccount(acc *t.Account, rows *sql.Rows) *u.ApiError {

	err := rows.Scan(
		&acc.ID,
		&acc.Firstname,
		&acc.Lastname,
		&acc.Displayname,
		&acc.Username,
		&acc.Email,
		&acc.WebsiteUrl,
		&acc.Description,
		&acc.ProfilePicUrl,
		&acc.HeaderPicUrl,
		&acc.CreatedAt,
		&acc.AccFolloweringCount,
		&acc.FollowersCount,
		&acc.Password,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}
