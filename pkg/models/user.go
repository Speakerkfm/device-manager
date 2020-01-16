package models

type User struct {
	ID    string `json:"id" jwt:"id"`
	Email string `json:"email" jwt:"email"`

	Secret string `json:"secret" jwt:"-"`
}
