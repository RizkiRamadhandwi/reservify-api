package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/repository"
	"booking-room-app/shared/model"
	"errors"
	"fmt"
)

type EmployeesUseCase interface {
	FindEmployeesByID(id string) (entity.Employee, error)
	FindEmployeesByUsername(username string) (entity.Employee, error)
	FindEmployeForLogin(username ,password string) (entity.Employee, error)
	RegisterNewEmployee(payload entity.Employee) (entity.Employee, error)
	UpdateEmployee(payload entity.Employee) (entity.Employee, error)
	ListAll(page, size int) ([]entity.Employee, model.Paging, error)
}

type employeesUseCase struct {
	repo repository.EmployeeRepository
}

// FindEmployeesByUsername implements EmployeesUseCase.
func (e *employeesUseCase) FindEmployeesByUsername(username string) (entity.Employee, error) {
	if username == "" {
		return entity.Employee{}, errors.New("username harus diisi")
	}
	return e.repo.GetEmployeesByUsername(username)
}

func (e *employeesUseCase)FindEmployeForLogin(username, password string) (entity.Employee, error){
	if username == "" {
		return entity.Employee{}, errors.New("username harus diisi")
	}
	return e.repo.GetEmployeesByUsernameForLogin(username, password)
}

// ListAll implements EmployeesUseCase.
func (e *employeesUseCase) ListAll(page int, size int) ([]entity.Employee, model.Paging, error) {
	// if page == 0 && size == 0 {
	// 	page = 1
	// 	size = 5
	// }
	return e.repo.List(page, size)
}

// FindEmployeesByID implements EmployeesUseCase.
func (e *employeesUseCase) FindEmployeesByID(id string) (entity.Employee, error) {
	if id == "" {
		return entity.Employee{}, errors.New("id harus diisi")
	}
	return e.repo.GetEmployeesByID(id)
}

// RegisterNewEmployee implements EmployeesUseCase.
func (e *employeesUseCase) RegisterNewEmployee(payload entity.Employee) (entity.Employee, error) {
	if payload.Name == "" || payload.Password == "" || payload.Role == "" || payload.Division == "" || payload.Position == "" || payload.Contact == "" {
		return entity.Employee{}, fmt.Errorf("oops, field required")
	}

	employee, err := e.repo.CreateEmployee(payload)
	if err != nil {
		return entity.Employee{}, fmt.Errorf("oppps, failed to save data employee :%v", err.Error())
	}
	return employee, nil
}

// UpdateEmployee implements EmployeesUseCase.
func (e *employeesUseCase) UpdateEmployee(payload entity.Employee) (entity.Employee, error) {
	if payload.ID == "" ||payload.Name == "" || payload.Password == "" || payload.Role == "" || payload.Division == "" || payload.Position == "" || payload.Contact == "" {
		return entity.Employee{}, fmt.Errorf("oops, field required")
	}

	employee, err := e.repo.UpdateEmployee(payload)
	if err != nil {
		return entity.Employee{}, fmt.Errorf("oppps, failed to save data employee :%v", err.Error())
	}
	return employee, nil
}

func NewEmployeeUseCase(repo repository.EmployeeRepository) EmployeesUseCase {
	return &employeesUseCase{repo: repo}
}
