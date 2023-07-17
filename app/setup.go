package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/mrkhay/creative-insider-backend/config"
	"github.com/mrkhay/creative-insider-backend/database"
	"github.com/mrkhay/creative-insider-backend/handlers"
	"github.com/mrkhay/creative-insider-backend/middleware"
	"github.com/mrkhay/creative-insider-backend/utility"
)

func SetupAndRun() error {

	// load env
	err := config.LoadENV()
	if err != nil {
		return err
	}

	// start database
	db, err := database.StartMongoDB()
	if err != nil {
		return err
	}

	defer database.CloseMongoDB()

	app := mux.NewRouter()

	// attach middleware
	app.Use(middleware.LoggerMiddleware)

	utility.SetupLogger()

	server := handlers.NewApiStorage(db)
	// server := handlers.NewApiStorage(nil)

	// setup routes
	server.SetupRoutes(app)

	// setup swagger doc
	config.AddSwaggerRoutes(app)

	// get the port and start
	port := os.Getenv("PORT")

	fmt.Println("Searving on port: ", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), app)

	return nil
}
