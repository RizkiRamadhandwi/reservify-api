package repo_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type RoomFacilityRepoMock struct {
	mock.Mock
}

func (m *RoomFacilityRepoMock) CreateRoomFacility(payload entity.RoomFacility, newFacilityQuantity int) (entity.RoomFacility, error) {
	args := m.Called(payload, newFacilityQuantity)
	return args.Get(0).(entity.RoomFacility), args.Error(1)
}

func (m *RoomFacilityRepoMock) GetQuantityFacilityByID(id string) (int, error) {
	args := m.Called(id)
	return args.Int(0), args.Error(1)
}

func (m *RoomFacilityRepoMock) GetRoomFacilityById(id string) (entity.RoomFacility, error) {
	args := m.Called(id)
	return args.Get(0).(entity.RoomFacility), args.Error(1)
}

func (m *RoomFacilityRepoMock) ListRoomFacility(page int, size int) ([]entity.RoomFacility, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.RoomFacility), args.Get(1).(model.Paging), args.Error(2)
}

func (m *RoomFacilityRepoMock) UpdateRoomFacility(payload entity.RoomFacility, newFacilityQuantity int) (entity.RoomFacility, error) {
	args := m.Called(payload, newFacilityQuantity)
	return args.Get(0).(entity.RoomFacility), args.Error(1)
}
