package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	t "github.com/mrkhay/creative-quill-backend/models"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func (d *Database) CreateNewArticle(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.NewArticleRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	article, error := t.NewArticle(req.AuthorID, req.Title, req.ThumbnailUrl, req.Content)
	if error != nil {
		return error
	}
	error = d.db.CreateArticle(article)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) GetArticle(w http.ResponseWriter, r *http.Request) *u.ApiError {

	res, err := d.db.GetAllArticles()

	if err != nil {
		return err
	}

	u.WriteJson(w, http.StatusCreated, res)

	return nil
}

func (d *Database) ModifieArticle(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.ModifieArticleRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ArticleID == "" || req.AuthorID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	article := t.ModifieArticleFunc(req)

	error := d.db.ModifieArticle(article)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) DeleteArticle(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.DeleteArticleRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ArticleID == "" || req.AuthorID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	error := d.db.DeleteArticle(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}
func (d *Database) LikeArticle(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.Like)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ContentID == "" || req.UserID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	req, error := t.NewLikeFunc(req)
	if err != nil {
		return error
	}

	error = d.db.LikeArticle(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) DisLikeArticle(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.Dislike)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ContentID == "" || req.UserID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	req, error := t.NewDislikeFunc(req)
	if err != nil {
		return error
	}

	error = d.db.DislikeArticle(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) CreateNewArticleComment(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.NewArticleCommentRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ArticleID == "" || req.UserID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	req, error := t.NewArticleCommentFunc(req)
	if err != nil {
		return error
	}

	error = d.db.CreateArticleNewComment(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, Success)
	return nil
}

func (d *Database) GetArticleComments(w http.ResponseWriter, r *http.Request) *u.ApiError {

	req := new(t.ArticleCommentsRequest)

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return u.NewError(err, http.StatusConflict)
	}

	if req.ArticleID == "" {

		return u.NewError(fmt.Errorf("missing field"), http.StatusBadRequest)
	}

	res, error := d.db.GetArticleComments(req)

	if error != nil {
		return error
	}

	u.WriteJson(w, http.StatusCreated, res)
	return nil
}
