package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type RoomUseCaseMock struct {
	mock.Mock
}

func (r *RoomUseCaseMock) RegisterNewRoom(payload entity.Room) (entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) FindRoomByID(id string) (entity.Room, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) FindAllRoom(page, size int) ([]entity.Room, model.Paging, error) {
	args := r.Called(page, size)
	return args.Get(0).([]entity.Room), args.Get(1).(model.Paging), args.Error(2)
}

func (r *RoomUseCaseMock) FindAllRoomStatus(status string, page, size int) ([]entity.Room, model.Paging, error) {
	args := r.Called(status, page, size)
	return args.Get(0).([]entity.Room), args.Get(1).(model.Paging), args.Error(2)
}

func (r *RoomUseCaseMock) UpdateRoomDetail(payload entity.Room) (entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomUseCaseMock) UpdateRoomStatus(payload entity.Room) (entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.Room), args.Error(1)
}
