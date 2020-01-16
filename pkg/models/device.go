package models

type Device struct {
	ID   string `json:"id" jwt:"id"`
	Name string `json:"name" jwt:"name"`

	OwnerID string `json:"owner_id" jwt:"owner_id"`
}
