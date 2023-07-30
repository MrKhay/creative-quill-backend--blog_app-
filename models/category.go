package models

import "github.com/google/uuid"

type CategoryName string
type Category struct {
	ID   uuid.UUID    `json:"id"`
	Name CategoryName `json:"category"`
}
