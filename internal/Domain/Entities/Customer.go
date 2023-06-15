package entities

import "time"

type Customer struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Document  string    `json:"document"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"createdAt"`
}
