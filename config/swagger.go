package config

import (
	"github.com/gorilla/mux"
	_ "github.com/mrkhay/creative-insider-backend/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

func AddSwaggerRoutes(app *mux.Router) {
	// Setup Swagger documentation
	app.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
}
