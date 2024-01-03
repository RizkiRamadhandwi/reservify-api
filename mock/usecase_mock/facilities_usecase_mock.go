package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type FacilitiesUseCaseMock struct {
	mock.Mock
}

func (m *FacilitiesUseCaseMock) FindAllFacilities(page, size int) ([]entity.Facilities, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.Facilities), args.Get(1).(model.Paging), args.Error(2)
}
func (m *FacilitiesUseCaseMock) RegisterNewFacilities(payload entity.Facilities) (entity.Facilities, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Facilities), args.Error(1)
}
func (m *FacilitiesUseCaseMock) FindFacilitiesById(id string) (entity.Facilities, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Facilities), args.Error(1)
}
func (m *FacilitiesUseCaseMock) EditFacilities(payload entity.Facilities) (entity.Facilities, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Facilities), args.Error(1)
}
