package repo_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type RoomRepoMock struct {
	mock.Mock
}

func (r *RoomRepoMock) Create(payload entity.Room) (entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomRepoMock) Get(id string) (entity.Room, error) {
	args := r.Called(id)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomRepoMock) List(page, size int) ([]entity.Room, model.Paging, error) {
	args := r.Called(page, size)
	return args.Get(0).([]entity.Room), args.Get(1).(model.Paging), args.Error(2)
}

func (r *RoomRepoMock) ListStatus(status string, page, size int) ([]entity.Room, model.Paging, error) {
	args := r.Called(status, page, size)
	return args.Get(0).([]entity.Room), args.Get(1).(model.Paging), args.Error(2)
}

func (r *RoomRepoMock) Update(payload entity.Room) (entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.Room), args.Error(1)
}

func (r *RoomRepoMock) UpdateStatus(payload entity.Room) (entity.Room, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.Room), args.Error(1)
}
