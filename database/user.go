package database

import (
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func (s *PostgresStorage) FollowUser(req *t.FollowRequest) *u.ApiError {

	query :=
		`insert into following (follower_id,following_id,date_followed) values($1,$2,$3)`

	_, err := s.db.Exec(
		query,
		req.UserID,
		req.AccountID,
		req.DateFollowed,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}

func (s *PostgresStorage) UnFollowUser(req *t.UnFollowRequest) *u.ApiError {

	query :=
		`DELETE FROM following WHERE follower_id = $1 AND following_id = $2`

	_, err := s.db.Exec(
		query,
		req.UserID,
		req.AccountID,
	)

	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	return nil

}
func (s *PostgresStorage) GetAccFollowingDetails(req *t.GetAccFollowingDetails) ([]*t.Account, *u.ApiError) {

	query :=
		`SELECT * FROM accounts WHERE id IN (SELECT follower_id FROM following WHERE following_id = $1);`

	row, err := s.db.Query(
		query,
		req.UserID,
	)
	accounts := []*t.Account{}

	for row.Next() {

		acc := new(t.Account)

		err := scanIntoAccount(acc, row)
		acc.Password = ""
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return accounts, nil

}

func (s *PostgresStorage) GetFollowerDetails(req *t.GetFollowersDetail) ([]*t.Account, *u.ApiError) {

	query :=
		`SELECT * FROM accounts WHERE id IN (SELECT following_id FROM following WHERE follower_id = $1);`

	row, err := s.db.Query(
		query,
		req.UserID,
	)
	accounts := []*t.Account{}

	for row.Next() {

		acc := new(t.Account)

		err := scanIntoAccount(acc, row)
		acc.Password = ""
		if err != nil {
			return nil, err
		}
		accounts = append(accounts, acc)
	}

	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return accounts, nil

}

func (s *PostgresStorage) UpdateAccount(req *t.UpdateAccountReq) *u.ApiError {

	query := `UPDATE accounts SET first_name = $1,
	last_name =$2,display_name=$3,
	user_name=$4,headerpic_url=$5,profilepic_url=$6,
	description=$7,website_url=$8,email=$9 WHERE id = $10`
	row, err := s.db.Query(
		query,
		req.Firstname,
		req.Lastname,
		req.Displayname,
		req.Username,
		req.HeaderPicUrl,
		req.ProfilePicUrl,
		req.Description,
		req.WebsiteUrl,
		req.Email,
		req.UserID,
	)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil

}

func (s *PostgresStorage) UpdateFirstName(req *t.UpdateAccDetailsRequest) *u.ApiError {

	row, err := s.db.Query(
		`UPDATE accounts SET first_name = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
func (s *PostgresStorage) UpdateLastName(req *t.UpdateAccDetailsRequest) *u.ApiError {
	row, err := s.db.Query(
		`UPDATE accounts SET last_name = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
func (s *PostgresStorage) UpdateDisplayName(req *t.UpdateAccDetailsRequest) *u.ApiError {

	row, err := s.db.Query(
		`UPDATE accounts SET display_name = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
func (s *PostgresStorage) UpdateUsername(req *t.UpdateAccDetailsRequest) *u.ApiError {
	row, err := s.db.Query(
		`UPDATE accounts SET user_name = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}

func (s *PostgresStorage) UpdateDescription(req *t.UpdateAccDetailsRequest) *u.ApiError {
	row, err := s.db.Query(
		`UPDATE accounts SET description = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
func (s *PostgresStorage) UpdateWebsiteUrl(req *t.UpdateAccDetailsRequest) *u.ApiError {
	row, err := s.db.Query(
		`UPDATE accounts SET website_url = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
func (s *PostgresStorage) UpdateProfilePicUrl(req *t.UpdateAccDetailsRequest) *u.ApiError {
	row, err := s.db.Query(
		`UPDATE accounts SET profilepic_url = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
func (s *PostgresStorage) UpdateHeaderPicUrl(req *t.UpdateAccDetailsRequest) *u.ApiError {
	row, err := s.db.Query(
		`UPDATE accounts SET headerpic_url = $1 WHERE id = $2`, req.Value, req.UserID)

	if err != nil {
		return u.NewError(err, http.StatusBadGateway)
	}

	acc := new(t.Account)
	for row.Next() {

		error := scanIntoAccount(acc, row)

		if error != nil {
			return error
		}

	}

	return nil
}
