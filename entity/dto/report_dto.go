package dto

import (
	"booking-room-app/entity"
	"time"
)

type ReportDto struct {
	ID             string            `json:"id"`
	EmployeeId     string            `json:"employeeId,omitempty"`
	RoomId         string            `json:"roomId,omitempty"`
	Employee       entity.Employee   `json:"employee"`
	Room           entity.Room       `json:"room"`
	RoomFacilities []RoomFacilityDto `json:"roomFacilities"`
	Description    string            `json:"description"`
	Status         string            `json:"status"`
	StartTime      time.Time         `json:"startTime"`
	EndTime        time.Time         `json:"endTime"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}
