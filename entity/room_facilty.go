package entity

import "time"

type RoomFacility struct {
	ID          string    `json:"id"`
	RoomId      string    `json:"roomId,omitempty"`
	FacilityId  string    `json:"facilityId"`
	Quantity    int       `json:"quantity"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

