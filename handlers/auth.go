package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/mrkhay/creative-quill-backend/database"
	t "github.com/mrkhay/creative-quill-backend/models"
	"github.com/mrkhay/creative-quill-backend/utility"
)

// Signup returns account with token.
//
// swagger:route POST /signup account Signup
//
// Returns account with token..
//
// Responses:
//
//	200: AccountResponse
//
// 500:
// 300

type Handlers struct {
	db database.Storage
}

func SetUpHandlers(s database.Storage) *Handlers {

	return &Handlers{
		db: s,
	}
}

func (h *Handlers) Signup(w http.ResponseWriter, r *http.Request) *t.ApiError {

	acc := new(t.Account)

	err := json.NewDecoder(r.Body).Decode(acc)
	if err != nil {
		return t.NewError(err, http.StatusConflict)
	}

	acc, error := t.NewAccount(acc.Firstname, acc.Lastname, acc.Email, acc.Password)

	if error != nil {
		return error
	}
	error = h.db.Signup(acc)

	if error != nil {
		return error
	}

	utility.WriteJson(w, http.StatusOK, acc)
	return nil
}
