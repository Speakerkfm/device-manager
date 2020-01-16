package models

type User struct {
	ID    string `json:"id" jwt:"id"`
	Email string `json:"email" jwt:"email"`

	DevicesIDs []string `json:"-" jwt:"-"`
}
