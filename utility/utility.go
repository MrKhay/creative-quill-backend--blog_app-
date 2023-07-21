package utility

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
	t "github.com/mrkhay/creative-quill-backend/models"
)

func WriteJson(w http.ResponseWriter, status int, v any) *t.ApiError {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		return &t.ApiError{
			Error:  err,
			Status: http.StatusUnprocessableEntity,
		}
	}

	return nil
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

func CreateJWT(account *t.Account) (string, *t.ApiError) {
	secreat := os.Getenv("JWT_SECRET")
	// create claims

	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"email":     account.Email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString([]byte(secreat))

	if err != nil {
		return "", t.NewError(err, http.StatusBadGateway)
	}

	return res, nil

}
