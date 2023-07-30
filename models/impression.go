package models

import (
	"net/http"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

// like
// dislike
type Like struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	ContentID string    `json:"content_id"`
}

type Dislike struct {
	ID        uuid.UUID `json:"id"`
	UserID    string    `json:"user_id"`
	ContentID string    `json:"content_id"`
}

func NewLikeFunc(like *Like) (*Like, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &Like{
		ID:        id,
		UserID:    like.UserID,
		ContentID: like.ContentID,
	}, nil
}

func NewDislikeFunc(like *Dislike) (*Dislike, *u.ApiError) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &Dislike{
		ID:        id,
		UserID:    like.UserID,
		ContentID: like.ContentID,
	}, nil
}
