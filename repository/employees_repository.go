package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"log"
	"math"
	"time"
)

type EmployeeRepository interface {
	GetEmployeesByID(id string) (entity.Employee, error)
	GetEmployeesByUsername(username string) (entity.Employee, error)
	GetEmployeesByUsernameForLogin(username, password string) (entity.Employee, error)
	CreateEmployee(payload entity.Employee) (entity.Employee, error)
	UpdateEmployee(payload entity.Employee) (entity.Employee, error)
	List(page, size int) ([]entity.Employee, model.Paging, error)
}

type employeeRepository struct {
	db *sql.DB
}

func (e *employeeRepository) CreateEmployee(payload entity.Employee) (entity.Employee, error) {
	var employee entity.Employee

	payload.UpdatedAt = time.Now()

	err := e.db.QueryRow(config.InsertEmployee,
		payload.Name,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Division,
		payload.Position,
		payload.Contact).Scan(&employee.ID, &employee.CreatedAt, &employee.UpdatedAt)

	if err != nil {
		log.Println("employeeRepository.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}

	employee.Name = payload.Name
	employee.Username = payload.Username
	employee.Password = payload.Password
	employee.Role = payload.Role
	employee.Division = payload.Division
	employee.Position = payload.Position
	employee.Contact = payload.Contact

	return employee, nil

}

// GetEmployeesByID implements EmployeeRepository.
func (e *employeeRepository) GetEmployeesByID(id string) (entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow(config.SelectEmployeeByID, id).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Username,
		&employee.Password,
		&employee.Role,
		&employee.Division,
		&employee.Position,
		&employee.Contact,
		&employee.CreatedAt,
		&employee.UpdatedAt)
	if err != nil {
		log.Println("employeeRepository.GetEmployeeByID.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) GetEmployeesByUsername(username string) (entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow(config.SelectEmployeeByUsername, username).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Username,
		&employee.Password,
		&employee.Role,
		&employee.Division,
		&employee.Position,
		&employee.Contact,
		&employee.CreatedAt,
		&employee.UpdatedAt)
	if err != nil {
		log.Println("employeeRepository.GetEmployeeByID.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	return employee, nil
}

func (e *employeeRepository) GetEmployeesByUsernameForLogin(username, password string) (entity.Employee, error) {
	var employee entity.Employee
	err := e.db.QueryRow(config.SelectEmployeeForLogin, username, password).Scan(
		&employee.ID,
		&employee.Name,
		&employee.Username,
		&employee.Password,
		&employee.Role)
	if err != nil {
		log.Println("employeeRepository.GetEmployeeByID.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	return employee, nil
}

// UpdateEmployee implements EmployeeRepository.
func (e *employeeRepository) UpdateEmployee(payload entity.Employee) (entity.Employee, error) {
	var employee entity.Employee
	payload.UpdatedAt = time.Now()

	err := e.db.QueryRow(config.UpdateEmployee,
		payload.Name,
		payload.Username,
		payload.Password,
		payload.Role,
		payload.Division,
		payload.Position,
		payload.Contact,
		payload.ID).Scan(&employee.CreatedAt, &employee.UpdatedAt)

	if err != nil {
		log.Println("employeeRepository.QueryRow: ", err.Error())
		return entity.Employee{}, err
	}
	employee.ID = payload.ID
	employee.Name = payload.Name
	employee.Username = payload.Username
	employee.Password = payload.Password
	employee.Role = payload.Role
	employee.Division = payload.Division
	employee.Position = payload.Position
	employee.Contact = payload.Contact
	employee.UpdatedAt = payload.UpdatedAt

	return employee, nil
}

func (e *employeeRepository) List(page, size int) ([]entity.Employee, model.Paging, error) {
	var employees []entity.Employee
	offset := (page - 1) * size
	rows, err := e.db.Query(config.SelectAllEmployee, size, offset)
	if err != nil {
		log.Println("employeeRepository.Query:", err.Error())
		return nil, model.Paging{}, err
	}
	for rows.Next() {
		var emp entity.Employee
		err := rows.Scan(
			&emp.ID,
			&emp.Name,
			&emp.Username,
			&emp.Password,
			&emp.Role,
			&emp.Division,
			&emp.Position,
			&emp.Contact,
			&emp.CreatedAt,
			&emp.UpdatedAt,
		)
		if err != nil {
			log.Println("employeeRepository.Rows.Next():", err.Error())
			return nil, model.Paging{}, err
		}

		employees = append(employees, emp)
	}

	totalRows := 0
	if err := e.db.QueryRow("SELECT COUNT(*) FROM employees").Scan(&totalRows); err != nil {
		return nil, model.Paging{}, err
	}

	paging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   totalRows,
		TotalPages:  int(math.Ceil(float64(totalRows) / float64(size))),
	}
	return employees, paging, nil
}

func NewEmployeeRepository(db *sql.DB) EmployeeRepository {
	return &employeeRepository{db: db}
}
