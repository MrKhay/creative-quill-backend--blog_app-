package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrkhay/creative-insider-backend/database"
	"github.com/mrkhay/creative-insider-backend/utility"
)

type ApiError struct {
	Error string `json:"error"`
}
type ApiSuccess struct {
	Success string `json:"success"`
}

type APISTORAGE struct {
	storage database.Storage
}

func NewApiStorage(s database.Storage) *APISTORAGE {
	return &APISTORAGE{
		storage: s,
	}
}

type apiFunc func(http.ResponseWriter, *http.Request) error

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			utility.WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}

	}
}

func (s *APISTORAGE) SetupRoutes(app *mux.Router) {

	// check health
	app.HandleFunc("/health", makeHttpHandleFunc(HealthCheckHandler))
	app.HandleFunc("/happy", makeHttpHandleFunc(s.Happy))

}
