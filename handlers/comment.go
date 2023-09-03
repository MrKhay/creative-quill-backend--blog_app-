package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func (d *Database) CreateNewArticleSubComment(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.NewArticleSubCommentRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ArticleID == "" || req.UserID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	req, error := t.NewArticleSubCommentFunc(req)
	if err != nil {
		return error
	}

	error = d.db.CreateArticleSubComment(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) LikeComment(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.NewCommentLikeRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.UserID == "" || req.CommentID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	error := d.db.LikeComment(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) DislikeComment(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.NewCommentDislikeRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.UserID == "" || req.CommentID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	error := d.db.DislikeComment(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) DeleteComment(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.DeleteCommentRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ID == "" || req.UserID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	error := d.db.DeleteComment(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}
func (d *Database) GetCommentSubComments(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.GetCommentSubCommentsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ArticleID == "" || req.ParentCommentID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	res, error := d.db.GetCommentSubComments(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, res)
	return nil
}
