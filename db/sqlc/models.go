// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type Account struct {
	ID        int64     `json:"id"`
	Owner     string    `json:"owner"`
	Message   string    `json:"message"`
	Mistakes  string    `json:"mistakes"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
}
