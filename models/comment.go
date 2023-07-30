package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type DeleteCommentRequest struct {
	ID        string `json:"id"`
	AuthorID  string `json:"author_id"`
	ArticleID string `json:"article_id"`
}
type NewCommentRequest struct {
	ID              uuid.UUID `json:"id"`
	UserID          string    `json:"user_id"`
	ArticleID       string    `json:"article_id"`
	ParentCommentID string    `json:"parent_comment_id"`
	DateCreated     time.Time `json:"date_created"`
	Content         string    `json:"content"`
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

func NewComment(articleID, user_id, parent_comment_id, content string) (*Comment, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &Comment{
		ID:              id,
		UserID:          user_id,
		ParentCommentID: parent_comment_id,
		ArticleID:       articleID,
		Likes:           0,
		Dislikes:        0,
		Comments:        0,
		DateCreated:     time.Now().UTC(),
		Content:         content,
	}, nil

}
