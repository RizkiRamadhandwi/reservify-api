package usecase_mock

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"

	"github.com/stretchr/testify/mock"
)

type EmployeeUseCaseMock struct {
	mock.Mock
}

// FindEmployeForLogin implements usecase.EmployeesUseCase.
func (e *EmployeeUseCaseMock) FindEmployeForLogin(username string, password string) (entity.Employee, error) {
	args := e.Called(username, password)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeUseCaseMock) RegisterNewEmployee(payload entity.Employee) (entity.Employee, error) {
	args := e.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}

func (e *EmployeeUseCaseMock) FindEmployeesByID(id string) (entity.Employee, error) {
	args := e.Called(id)
	return args.Get(0).(entity.Employee), args.Error(1)

}
func (e *EmployeeUseCaseMock) FindEmployeesByUsername(username string) (entity.Employee, error) {
	args := e.Called(username)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (e *EmployeeUseCaseMock) FindEmployeeForLogin(username, password string) (entity.Employee, error) {
	args := e.Called(username, password)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (e *EmployeeUseCaseMock) UpdateEmployee(payload entity.Employee) (entity.Employee, error) {
	args := e.Called(payload)
	return args.Get(0).(entity.Employee), args.Error(1)
}
func (e *EmployeeUseCaseMock) ListAll(page, size int) ([]entity.Employee, model.Paging, error) {
	args := e.Called(page, size)
	return args.Get(0).([]entity.Employee), args.Get(1).(model.Paging), args.Error(2)
}
