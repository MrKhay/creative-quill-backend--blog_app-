package models

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

type NewArticleRequest struct {
	ID           uuid.UUID       `json:"id"`
	AuthorID     string          `json:"author_id"`
	Title        string          `json:"title"`
	ThumbnailUrl string          `json:"thumbnail_url"`
	LastModified time.Time       `json:"last_modified"`
	DateCreated  time.Time       `jaon:"date_created"`
	Content      json.RawMessage `json:"content"`
}
type ModifieArticleRequest struct {
	AuthorID     string          `json:"author_id"`
	ArticleID    string          `json:"article_id"`
	Title        string          `json:"title"`
	ThumbnailUrl string          `json:"thumbnail_url"`
	LastModified time.Time       `json:"last_modified"`
	Content      json.RawMessage `json:"content"`
}
type DeleteArticleRequest struct {
	AuthorID  string `json:"author_id"`
	ArticleID string `json:"article_id"`
}
type Article struct {
	ID           uuid.UUID       `json:"id"`
	AuthorID     string          `json:"author_id"`
	Title        string          `json:"title"`
	ThumbnailUrl string          `json:"thumbnail_url"`
	Likes        int             `json:"likes"`
	Dislikes     int             `json:"dislikes"`
	Comments     int             `json:"comments"`
	DateCreated  time.Time       `json:"date_created"`
	LastModified time.Time       `json:"last_modified"`
	Content      json.RawMessage `json:"content"`
}

func NewArticle(authorid, title, thumbnail_url string, content json.RawMessage) (*NewArticleRequest, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &NewArticleRequest{
		ID:           id,
		AuthorID:     authorid,
		Title:        title,
		ThumbnailUrl: thumbnail_url,
		LastModified: time.Now().UTC(),
		DateCreated:  time.Now().UTC(),
		Content:      content,
	}, nil
}
func ModifieArticleFunc(article *ModifieArticleRequest) *ModifieArticleRequest {

	return &ModifieArticleRequest{
		ArticleID:    article.ArticleID,
		AuthorID:     article.AuthorID,
		Title:        article.Title,
		ThumbnailUrl: article.ThumbnailUrl,
		LastModified: time.Now().UTC(),
		Content:      article.Content,
	}
}
