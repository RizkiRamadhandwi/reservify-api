package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type RoomFacilityUseCaseMock struct {
	mock.Mock
}

func (r *RoomFacilityUseCaseMock) AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.RoomFacility), args.Error(1)
}

func (r *RoomFacilityUseCaseMock) FindAllRoomFacility(page, size int) ([]entity.RoomFacility, model.Paging, error) {
	args := r.Called(page, size)
	return args.Get(0).([]entity.RoomFacility), args.Get(1).(model.Paging), args.Error(2)
}

func (r *RoomFacilityUseCaseMock) FindRoomFacilityById(id string) (entity.RoomFacility, error) {
	args := r.Called(id)
	return args.Get(0).(entity.RoomFacility), args.Error(1)
}

func (r *RoomFacilityUseCaseMock) UpdateRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, error) {
	args := r.Called(payload)
	return args.Get(0).(entity.RoomFacility), args.Error(1)
}
