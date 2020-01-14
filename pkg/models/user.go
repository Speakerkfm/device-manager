package models

import "github.com/satori/go.uuid"

type User struct {
	ID    uuid.UUID
	Email string
}
