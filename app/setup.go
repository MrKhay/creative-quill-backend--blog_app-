package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mrkhay/creative-quill-backend/config"
	"github.com/mrkhay/creative-quill-backend/database"
	"github.com/mrkhay/creative-quill-backend/handlers"
	"github.com/mrkhay/creative-quill-backend/middleware"
	"github.com/mrkhay/creative-quill-backend/router"
	u "github.com/mrkhay/creative-quill-backend/utility"
)

func SetupAndRun() *u.ApiError {

	// load env
	err := config.LoadENV()
	if err != nil {
		fmt.Println("1")
		return u.NewError(err, http.StatusConflict)
	}

	// start database
	db, error := database.NewPostgresStorage()
	if error != nil {
		fmt.Println(err.Error())
		return error
	}

	// create tables
	db.Init()

	app := mux.NewRouter()

	// attach middleware
	app.Use(middleware.LoggerMiddleware)

	// setup logger
	u.SetupLogger()

	// setup handlers
	handler := handlers.SetUpHandlers(db)

	// setup routes
	router.SetupRoutes(app, *handler)

	// setup swagger doc
	config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")

	fmt.Println("Searving on port: ", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), app)

	return nil
}
