package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type UserUseCaseMock struct {
	mock.Mock
}

// FindEmployeForLogin implements usecase.EmployeesUseCase.
func (t *UserUseCaseMock) FindEmployeForLogin(username string, password string) (entity.Employee, error) {
	args := t.Called(username, password)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (m *UserUseCaseMock) FindEmployeesByID(id string) (entity.Employee, error) {
	args := m.Called(id)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (m *UserUseCaseMock) FindEmployeesByUsername(username string) (entity.Employee, error) {
	args := m.Called(username)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (m *UserUseCaseMock) RegisterNewEmployee(payload entity.Employee) (entity.Employee, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (m *UserUseCaseMock) UpdateEmployee(payload entity.Employee) (entity.Employee, error) {
	args := m.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (m *UserUseCaseMock) ListAll(page, size int) ([]entity.Employee, model.Paging, error) {
	args := m.Called(page, size)
	return args.Get(0).([]entity.Employee), args.Get(1).(model.Paging), args.Error(2)
}
