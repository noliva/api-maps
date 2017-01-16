package models

type MyMap struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Boundaries  string    `json:"boundaries"`
	GroupOwner  string    `json:"groupOwner"db:"group_owner"`
	Addresses   []address `json:"addresses"`
	CreatedAt   string    `json:"created_at"db:"created_at"`
}