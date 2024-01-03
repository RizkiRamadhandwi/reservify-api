package entity

import "time"

type Facilities struct {
	ID        string    `json:"id"`
	Name      string    `json:"name" validate:"required,min=3,max=100"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
