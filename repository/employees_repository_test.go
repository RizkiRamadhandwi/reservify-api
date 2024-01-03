package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectEmployee = entity.Employee{
	ID:        "1",
	Name:      "John Doe",
	Username:  "JohnDoe123",
	Password:  "abc5dasar",
	Role:      "GA",
	Division:  "Marketing",
	Position:  "Manager",
	Contact:   "612906347",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type EmployeeRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    EmployeeRepository
}

func (suite *EmployeeRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewEmployeeRepository(suite.mockDb)
}

func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByID_success() {

	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE id = $1`)).WithArgs(expectEmployee.ID).WillReturnRows(rows)

	_, actualError := suite.repo.GetEmployeesByID(expectEmployee.ID)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
}
func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByUsername_success() {
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, username, password, role, division, position, contact, created_at, updated_at FROM employees WHERE username = $1`)).WithArgs(expectEmployee.Username).WillReturnRows(rows)

	_, actualError := suite.repo.GetEmployeesByUsername(expectEmployee.Username)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
}

func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeById_Fail() {
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.GetEmployeesByID("12")
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}
func (suite *EmployeeRepositoryTestSuite) TestGetEmployeeByUsername_Fail() {
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.GetEmployeesByUsername("12")
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeRepositoryTestSuite) TestCreateEmployee_success() {
	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expectEmployee.ID, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertEmployee)).WithArgs(
		expectEmployee.Name,
		expectEmployee.Username,
		expectEmployee.Password,
		expectEmployee.Role,
		expectEmployee.Division,
		expectEmployee.Position,
		expectEmployee.Contact).WillReturnRows(rows)

	_, err := suite.repo.CreateEmployee(expectEmployee)
	suite.NoError(err)
}

func (suite *EmployeeRepositoryTestSuite) TestCreateEmployee_Fail() {
	rows := sqlmock.NewRows([]string{"id", "created_at"}).AddRow(expectEmployee.ID, expectEmployee.CreatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertEmployee)).WithArgs(
		expectEmployee.Name,
		expectEmployee.Username,
		expectEmployee.Password,
		expectEmployee.Role,
		expectEmployee.Division,
		expectEmployee.Position,
		expectEmployee.Contact).WillReturnRows(rows)

	_, err := suite.repo.CreateEmployee(expectEmployee)
	suite.Error(err)
}

func (suite *EmployeeRepositoryTestSuite) TestGetUsernameForLogin_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role"}).AddRow(
		expectEmployee.ID,
		expectEmployee.Name,
		expectEmployee.Username,
		expectEmployee.Password,
		expectEmployee.Role)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectEmployeeForLogin)).WithArgs(expectEmployee.Username, expectEmployee.Password).WillReturnRows(rows)

	_, actualError := suite.repo.GetEmployeesByUsernameForLogin(expectEmployee.Username, expectEmployee.Password)
	assert.Nil(suite.T(), actualError)
	assert.NoError(suite.T(), actualError)
}

func (suite *EmployeeRepositoryTestSuite) TestGetUsernameForLogin_Fail() {
	rows := sqlmock.NewRows([]string{"id", "name", "username", "password"}).AddRow(
		expectEmployee.ID,
		expectEmployee.Name,
		expectEmployee.Username,
		expectEmployee.Password)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectEmployeeForLogin)).WithArgs(expectEmployee.Username, expectEmployee.Password).WillReturnRows(rows)

	_, actualError := suite.repo.GetEmployeesByUsernameForLogin(expectEmployee.Username, expectEmployee.Password)
	assert.NotNil(suite.T(), actualError)
	assert.Error(suite.T(), actualError)
}

func (suite *EmployeeRepositoryTestSuite) TestUpdateEmployee_success() {
	rows := sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateEmployee)).WithArgs(
		expectEmployee.Name,
		expectEmployee.Username,
		expectEmployee.Password,
		expectEmployee.Role,
		expectEmployee.Division,
		expectEmployee.Position,
		expectEmployee.Contact,
		expectEmployee.ID).WillReturnRows(rows)

	_, err := suite.repo.UpdateEmployee(expectEmployee)
	suite.NoError(err)
}

func (suite *EmployeeRepositoryTestSuite) TestUpdateEmployee_Fail() {
	rows := sqlmock.NewRows([]string{"created_at"}).AddRow(expectEmployee.CreatedAt)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateEmployee)).WithArgs(
		expectEmployee.Name,
		expectEmployee.Username,
		expectEmployee.Password,
		expectEmployee.Role,
		expectEmployee.Division,
		expectEmployee.Position,
		expectEmployee.Contact,
		expectEmployee.ID).WillReturnRows(rows)

	_, err := suite.repo.UpdateEmployee(expectEmployee)
	suite.Error(err)
}

func (suite *EmployeeRepositoryTestSuite) TestList_Success() {
	page := 1
	size := 10
	expectEmployee := []entity.Employee{
		{ID: "1", Name: "John Doe", Username: "johndoe123", Password: "abc5dasar", Role: "Admin", Division: "PM", Position: "Manager", Contact: "62654398564", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Name: "John", Username: "johndoe", Password: "abc5dasar", Role: "Admin", Division: "PM", Position: "Manager", Contact: "62654398564", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	expectedPaging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   2,
		TotalPages:  1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).
		AddRow("1", "John Doe", "johndoe123", "abc5dasar", "Admin", "PM", "Manager", "62654398564", time.Now(), time.Now()).
		AddRow("2", "John", "johndoe", "abc5dasar", "Admin", "PM", "Manager", "62654398564", time.Now(), time.Now())

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).
		WillReturnRows(sqlmock.NewRows([]string{"total_rows"}).AddRow(2))

	employees, paging, err := suite.repo.List(page, size)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, employees)
	assert.Equal(suite.T(), expectedPaging, paging)
}

func (suite *EmployeeRepositoryTestSuite) TestScanEmployee_Fail() {
	page := 1
	size := 10
	expectEmployee := []entity.Employee{
		{ID: "1", Name: "John Doe", Username: "johndoe123", Password: "abc5dasar", Role: "Admin", Division: "PM", Position: "Manager", Contact: "62654398564", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Name: "John", Username: "johndoe", Password: "abc5dasar", Role: "Admin", Division: "PM", Position: "Manager", Contact: "62654398564", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}

	rows := sqlmock.NewRows([]string{"name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).
		AddRow(
			expectEmployee[0].Name,
			expectEmployee[0].Username,
			expectEmployee[0].Password,
			expectEmployee[0].Role,
			expectEmployee[0].Division,
			expectEmployee[0].Position,
			expectEmployee[0].Contact,
			expectEmployee[0].CreatedAt,
			expectEmployee[0].UpdatedAt,
		)

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnRows(rows)

	_, _, err := suite.repo.List(page, size)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeRepositoryTestSuite) TestList_Fail() {
	page := 1
	size := 10

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnError(errors.New("some SQL error"))

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeRepositoryTestSuite) TestList_ScanTotalRows() {
	page := 1
	size := 10

	rows := sqlmock.NewRows([]string{"id", "name", "username", "password", "role", "division", "position", "contact", "created_at", "updated_at"}).AddRow(
		expectEmployee.ID, expectEmployee.Name, expectEmployee.Username, expectEmployee.Password, expectEmployee.Role, expectEmployee.Division, expectEmployee.Position, expectEmployee.Contact, expectEmployee.CreatedAt, expectEmployee.UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).
		WillReturnError(errors.New("error"))

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestEmployeeRepositoryTestSuite(e *testing.T) {
	suite.Run(e, new(EmployeeRepositoryTestSuite))
}
