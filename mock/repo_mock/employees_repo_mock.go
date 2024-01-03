package repo_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type EmployeeRepoMock struct {
	mock.Mock
}

// GetEmployeesByUsernameForLogin implements repository.EmployeeRepository.
func (e *EmployeeRepoMock) GetEmployeesByUsernameForLogin(username string, password string) (entity.Employee, error) {
	args := e.Called(username, password)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeRepoMock) GetEmployeesByID(id string) (entity.Employee, error) {
	args := e.Called(id)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeRepoMock) GetEmployeesByUsername(username string) (entity.Employee, error) {
	args := e.Called(username)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeRepoMock) CreateEmployee(payload entity.Employee) (entity.Employee, error) {
	args := e.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeRepoMock) UpdateEmployee(payload entity.Employee) (entity.Employee, error) {
	args := e.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeRepoMock) List(page, size int) ([]entity.Employee, model.Paging, error) {
	args := e.Called(page, size)
	return args.Get(0).([]entity.Employee), args.Get(1).(model.Paging), args.Error(2)
}