package handlers

import (
	"net/http"

	"github.com/mrkhay/creative-insider-backend/utility"
)

// CreateAccount returns account with token.
//
// swagger:route POST /account account CreateAccount
//
// Returns account with token..
//
// Responses:
//
//	200: AccountResponse
//
// 500:
// 300
func (s *APISTORAGE) Happy(w http.ResponseWriter, r *http.Request) error {

	utility.WriteJson(w, http.StatusOK, "Happy")
	return nil
}
