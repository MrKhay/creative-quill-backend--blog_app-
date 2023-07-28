package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func (h *Database) FollowUser(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.FollowRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.UserID == "" || req.AccountID == "" {
		return u.NewError(fmt.Errorf("invalid"), http.StatusConflict)
	}

	if req.UserID == req.AccountID {

		return u.NewError(fmt.Errorf("invalid"), http.StatusConflict)

	}
	error := h.db.FollowUser(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, "Success")
	return nil
}

func (h *Database) UnFollowUser(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UnFollowRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.UserID == "" || req.AccountID == "" {
		return u.NewError(fmt.Errorf("invalid"), http.StatusConflict)
	}

	if req.UserID == req.AccountID {

		return u.NewError(fmt.Errorf("invalid"), http.StatusConflict)

	}
	error := h.db.UnFollowUser(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, "Success")
	return nil
}
func (h *Database) GetFollowersAccDetail(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.GetAccFollowingDetails)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.UserID == "" {
		return u.NewError(fmt.Errorf("invalid"), http.StatusConflict)
	}

	res, error := h.db.GetAccFollowingDetails(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, res)
	return nil
}

func (h *Database) GetFolloweringAccDetails(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.GetFollowersDetail)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.UserID == "" {
		return u.NewError(fmt.Errorf("invalid"), http.StatusConflict)
	}

	res, error := h.db.GetFollowerDetails(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, res)
	return nil
}

func (d *Database) UpdateAccount(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccountReq)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateAccount(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateFirstName(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateFirstName(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateLastName(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateLastName(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateDisplayName(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateDisplayName(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateUsername(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateUsername(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateWebsiteUrl(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateWebsiteUrl(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateProfilePicUrl(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateProfilePicUrl(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateDescription(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateDescription(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
func (d *Database) UpdateHeaderPicUrl(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.UpdateAccDetailsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	error := d.db.UpdateHeaderPicUrl(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusAccepted, Success)
	return nil
}
