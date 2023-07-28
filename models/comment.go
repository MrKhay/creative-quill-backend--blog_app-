package models

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type Comment struct {
	ID          uuid.UUID `json:"id"`
	ArticleID   string    `json:"article_id"`
	DateCreated time.Time `json:"date_created"`
	Content     string    `json:"content"`
}

func NewComment(articleID, content string) (*Comment, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &Comment{
		ID:          id,
		ArticleID:   articleID,
		DateCreated: time.Now().UTC(),
		Content:     content,
	}, nil

}
