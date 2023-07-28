package models

import (
	"net/http"

	"github.com/google/uuid"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

// like
// dislike
type Impression struct {
	ID     uuid.UUID `json:"id"`
	UserID string    `json:"account_id"`
}

func NewImpression(userid string) (*Impression, *u.ApiError) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, u.NewError(err, http.StatusConflict)
	}

	return &Impression{
		ID:     id,
		UserID: userid,
	}, nil
}
