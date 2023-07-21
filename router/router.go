package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/mrkhay/creative-quill-backend/handlers"
	t "github.com/mrkhay/creative-quill-backend/models"
	"github.com/mrkhay/creative-quill-backend/utility"
)

type ApiError struct {
	Error string `json:"error"`
}
type ApiSuccess struct {
	Success string `json:"success"`
}

type apiFunc func(http.ResponseWriter, *http.Request) *t.ApiError

func makeHttpHandleFunc(f apiFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		if err := f(w, r); err != nil {
			utility.WriteJson(w, err.Status, err.Error.Error())
		}

	}
}

func SetupRoutes(app *mux.Router, h handlers.Handlers) {
	// check health
	app.HandleFunc("/health", makeHttpHandleFunc(handlers.HealthCheckHandler))
	app.HandleFunc("/signup", makeHttpHandleFunc(h.Signup)).Methods("POST")

}
