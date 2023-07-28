package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type Article struct {
	ID            uuid.UUID       `json:"id"`
	AuthorID      string          `json:"author_id"`
	Title         string          `json:"title"`
	ThumbnailUrl  string          `json:"thumbnail_url"`
	LikesCount    int             `json:"likes"`
	DislikesCount int             `json:"dislikes"`
	CommentsCount int             `json:"comments"`
	DateCreated   time.Time       `jaon:"date_created"`
	Content       json.RawMessage `json:"content"`
}

func NewArticle(authorid, title, thumbnail_url string, content json.RawMessage) (*Article, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &Article{
		ID:            id,
		AuthorID:      authorid,
		Title:         title,
		ThumbnailUrl:  thumbnail_url,
		LikesCount:    0,
		DislikesCount: 0,
		CommentsCount: 0,
		DateCreated:   time.Now().UTC(),
		Content:       content,
	}, nil
}
