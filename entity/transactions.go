package entity

import "time"

type Transaction struct {
	ID          string `json:"id"`
	EmployeeId  string `json:"employeeId"`
	RoomId      string `json:"roomId"`
	RoomFacilities []RoomFacility `json:"roomFacilities,omitempty"`
	Description string `json:"description"`
	Status      string `json:"status"`
	StartTime time.Time `json:"startTime"`
	EndTime time.Time `json:"endTime"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`

}