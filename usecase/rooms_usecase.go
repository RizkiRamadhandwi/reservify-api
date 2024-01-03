package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"fmt"
	"strings"
)

type RoomUseCase interface {
	RegisterNewRoom(payload entity.Room) (entity.Room, error)
	FindRoomByID(id string) (entity.Room, error)
	FindAllRoom(page, size int) ([]entity.Room, model.Paging, error)
	FindAllRoomStatus(status string, page, size int) ([]entity.Room, model.Paging, error)
	UpdateRoomDetail(payload entity.Room) (entity.Room, error)
	UpdateRoomStatus(payload entity.Room) (entity.Room, error)
}

type roomUseCase struct {
	repo repository.RoomRepository
}

// FindAllRoom implements RoomUseCase.
func (r *roomUseCase) FindAllRoom(page, size int) ([]entity.Room, model.Paging, error) {
	return r.repo.List(page, size)
}

// FindAllRoomStatus implements RoomUseCase.
func (r *roomUseCase) FindAllRoomStatus(status string, page, size int) ([]entity.Room, model.Paging, error) {
	return r.repo.ListStatus(status, page, size)
}

// FindRoomByID implements RoomUseCase.
func (r *roomUseCase) FindRoomByID(id string) (entity.Room, error) {
	return r.repo.Get(id)
}

// RegisterNewRoom implements RoomUseCase.
func (r *roomUseCase) RegisterNewRoom(payload entity.Room) (entity.Room, error) {
	if payload.Name == "" || payload.RoomType == "" || payload.Capacity == 0 {
		return entity.Room{}, fmt.Errorf("oops, field required")
	}
	if payload.Status == "" {
		payload.Status = "available"
	} else {
		payload.Status = strings.ToLower(payload.Status)
	}
	// payload.UpdatedAt = time.Now()
	room, err := r.repo.Create(payload)
	if err != nil {
		return entity.Room{}, fmt.Errorf("failed to create a new room list: %v", err.Error())
	}
	return room, nil
}

// UpdateRoomDetail implements RoomUseCase.
func (r *roomUseCase) UpdateRoomDetail(payload entity.Room) (entity.Room, error) {
	if payload.ID == "" || payload.Name == "" || payload.RoomType == "" || payload.Capacity == 0 {
		return entity.Room{}, fmt.Errorf("oops, field required")
	}
	if payload.Status == "" {
		payload.Status = "available"
	} else {
		payload.Status = strings.ToLower(payload.Status)
	}

	// payload.UpdatedAt = time.Now()
	room, err := r.repo.Update(payload)
	if err != nil {
		return entity.Room{}, fmt.Errorf("failed to update room with ID %s: %v", payload.ID, err.Error())
	}
	return room, nil
}

// UpdateRoomStatus implements RoomUseCase.
func (r *roomUseCase) UpdateRoomStatus(payload entity.Room) (entity.Room, error) {
	if payload.ID == "" {
		return entity.Room{}, fmt.Errorf("oops, field required")
	}
	if payload.Status == "" {
		payload.Status = "available"
	} else {
		payload.Status = strings.ToLower(payload.Status)
	}

	// payload.UpdatedAt = time.Now()
	room, err := r.repo.UpdateStatus(payload)
	if err != nil {
		return entity.Room{}, fmt.Errorf("failed to update room with ID %s: %v", payload.ID, err.Error())
	}
	return room, nil
}

func NewRoomUseCase(repo repository.RoomRepository) RoomUseCase {
	return &roomUseCase{repo: repo}
}
