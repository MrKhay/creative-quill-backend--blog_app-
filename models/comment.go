package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type DeleteCommentRequest struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
}
type GetCommentSubCommentsRequest struct {
	ArticleID       string `json:"article_id"`
	ParentCommentID string `json:"parent_comment_id"`
}

type NewArticleCommentRequest struct {
	ID          uuid.UUID `json:"id"`
	UserID      string    `json:"user_id"`
	ArticleID   string    `json:"article_id"`
	DateCreated time.Time `json:"date_created"`
	Content     string    `json:"content"`
}
type NewCommentLikeRequest struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}
type NewCommentDislikeRequest struct {
	UserID    string `json:"user_id"`
	CommentID string `json:"comment_id"`
}

type Comment struct {
	ID              uuid.UUID `json:"id"`
	UserID          string    `json:"user_id"`
	ArticleID       string    `json:"article_id"`
	ParentCommentID string    `json:"parent_comment_id"`
	Likes           int       `json:"likes"`
	Dislikes        int       `json:"dislikes"`
	Comments        int       `json:"comments"`
	DateCreated     time.Time `json:"date_created"`
	Content         string    `json:"content"`
}
type NewArticleSubCommentRequest struct {
	ID              uuid.UUID `json:"id"`
	UserID          string    `json:"user_id"`
	ArticleID       string    `json:"article_id"`
	ParentCommentID uuid.UUID `json:"parent_comment_id"`
	DateCreated     time.Time `json:"date_created"`
	Content         string    `json:"content"`
}

func NewArticleCommentFunc(req *NewArticleCommentRequest) (*NewArticleCommentRequest, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &NewArticleCommentRequest{
		ID:          id,
		UserID:      req.UserID,
		ArticleID:   req.ArticleID,
		DateCreated: time.Now(),
		Content:     req.Content,
	}, nil

}

func NewArticleSubCommentFunc(req *NewArticleSubCommentRequest) (*NewArticleSubCommentRequest, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &NewArticleSubCommentRequest{
		ID:              id,
		UserID:          req.UserID,
		ArticleID:       req.ArticleID,
		ParentCommentID: req.ParentCommentID,
		DateCreated:     time.Now(),
		Content:         req.Content,
	}, nil

}
