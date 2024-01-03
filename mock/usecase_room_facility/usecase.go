package usecase_room_facility

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type RoomFacilityUseCaseMock struct {
	mock.Mock
}

func (r *RoomFacilityUseCaseMock) AddRoomFacilityTransaction(payload entity.RoomFacility) (entity.RoomFacility, int, error) {
	args := r.Called()
	return args.Get(0).(entity.RoomFacility), args.Int(1), args.Error(2)
}

func (r *RoomFacilityUseCaseMock) FindAllRoomFacility(int, int) ([]entity.RoomFacility, model.Paging, error) {
	args := r.Called()
	return args.Get(0).([]entity.RoomFacility), args.Get(1).(model.Paging), args.Error(2)
}

func (r *RoomFacilityUseCaseMock) FindRoomFacilityById(string) (entity.RoomFacility, int, error) {
	args := r.Called()
	return args.Get(0).(entity.RoomFacility), args.Int(1), args.Error(2)
}

func (r *RoomFacilityUseCaseMock) UpdateRoomFacilityTransaction(entity.RoomFacility) (entity.RoomFacility, int, error) {
	args := r.Called()
	return args.Get(0).(entity.RoomFacility), args.Int(1), args.Error(2)
}
