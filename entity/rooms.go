package entity

import "time"

type Room struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	RoomType  string    `json:"room_type"`
	Capacity  int       `json:"capacity"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
