package utility

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/golang-jwt/jwt"
)

type ApiError struct {
	Error  error
	Status int
}

func NewError(error error, code int) *ApiError {
	return &ApiError{
		Error:  error,
		Status: code,
	}
}

func WriteJson(w http.ResponseWriter, status int, v any) *ApiError {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)

	if err != nil {
		return &ApiError{
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

func CreateJWT(email *string) (string, *ApiError) {
	secreat := os.Getenv("JWT_SECRET")
	// create claims
	claims := &jwt.MapClaims{
		"expiresAt": 15000,
		"email":     email,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	res, err := token.SignedString([]byte(secreat))

	if err != nil {
		return "", NewError(err, http.StatusBadGateway)
	}

	return res, nil

}

func ValidateJWT(tokenString string) (*jwt.Token, error) {
	secreat := os.Getenv("JWT_SECRET")

	return jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {

		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}

		return []byte(secreat), nil
	})

}
