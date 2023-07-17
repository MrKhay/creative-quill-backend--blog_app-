package utility

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func WriteJson(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func SetupLogger() {
	// Open log file for writing
	file, err := os.OpenFile("logfile.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	
	if err != nil {
		log.Fatal(err)
	}

	// Set log output to the opened file
	log.SetOutput(file)
	log.SetFlags(0)
}
