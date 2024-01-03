package repo_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type FacilitiesRepoMock struct {
	mock.Mock
}

func (m *FacilitiesRepoMock) List(page, size int) ([]entity.Facilities, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.Facilities), args.Get(1).(model.Paging), args.Error(2)
}
func (m *FacilitiesRepoMock) Create(payload entity.Facilities) (entity.Facilities, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Facilities), args.Error(1)
}
func (m *FacilitiesRepoMock) GetById(id string) (entity.Facilities, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Facilities), args.Error(1)
}
func (m *FacilitiesRepoMock) UpdateById(payload entity.Facilities) (entity.Facilities, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Facilities), args.Error(1)
}
