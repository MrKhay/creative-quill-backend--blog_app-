package models

import "time"

// Account represents a account object.
// swagger:model
type Account struct {
	ID                int       `json:"-"`
	FirstName         string    `json:"firstname"`
	LastName          string    `json:"lastname"`
	AccountNumber     int64     `json:"acc_number"`
	Email             string    `json:"email"`
	EncryptedPassword string    `json:"-"`
	Balance           string    `json:"balance"`
	CreatedAt         time.Time `json:"createdAt"`
}
