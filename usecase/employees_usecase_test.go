package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/mock/repo_mock"
	"booking-room-app/shared/model"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectEmployee = entity.Employee{
	ID:        "1",
	Name:      "John Doe",
	Username:  "johndoe",
	Password:  "johndoe001",
	Role:      "admin",
	Division:  "HR",
	Position:  "Manager",
	Contact:   "124325463",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type EmployeeUseCaseTestSuite struct {
	suite.Suite
	erm *repo_mock.EmployeeRepoMock
	euc EmployeesUseCase
}

func (suite *EmployeeUseCaseTestSuite) SetupTest(){
	suite.erm = new(repo_mock.EmployeeRepoMock)
	suite.euc = NewEmployeeUseCase(suite.erm)
}

func (suite *EmployeeUseCaseTestSuite) TestListAll_success(){
	suite.erm.On("List", 1, 5).Return([]entity.Employee{}, model.Paging{}, nil)
	actualEmployee, actualPaging, err := suite.euc.ListAll(1, 5)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), model.Paging{}, actualPaging)
	assert.Equal(suite.T(), []entity.Employee{}, actualEmployee)
}
func (suite *EmployeeUseCaseTestSuite) TestUpdateEmployee_success() {
	suite.erm.On("UpdateEmployee", expectEmployee).Return(expectEmployee, nil)

	actual, err := suite.euc.UpdateEmployee(expectEmployee)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actual)
}
func (suite *EmployeeUseCaseTestSuite) TestCreateEmployee_success() {
	suite.erm.On("CreateEmployee", expectEmployee).Return(expectEmployee, nil)

	actual, err := suite.euc.RegisterNewEmployee(expectEmployee)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actual)
}
func (suite *EmployeeUseCaseTestSuite) TestCreateEmployee_Fail() {
	suite.erm.On("CreateEmployee", expectEmployee).Return(entity.Employee{}, fmt.Errorf("error"))

	_, err := suite.euc.RegisterNewEmployee(expectEmployee)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)

}
func (suite *EmployeeUseCaseTestSuite) TestUpdateEmployee_Fail(){
	suite.erm.On("UpdateEmployee", expectEmployee).Return(entity.Employee{}, fmt.Errorf("error"))

	_, err := suite.euc.UpdateEmployee(expectEmployee)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)

}
func (suite *EmployeeUseCaseTestSuite) TestCreateEmployee_emptyField(){
	payload := entity.Employee{
		ID:        "1",
		Name:      "John Doe",
		Username:  "johndoe",
		Password:  "johndoe001",
		Role:      "admin",
		Division:  "HR",
		Position:  "",
		Contact:   "124325463",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := suite.euc.UpdateEmployee(payload)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}
func (suite *EmployeeUseCaseTestSuite) TestUpdateEmployee_emptyField(){
	payload := entity.Employee{
		ID:        "1",
		Name:      "John Doe",
		Username:  "johndoe",
		Password:  "johndoe001",
		Role:      "admin",
		Division:  "HR",
		Position:  "",
		Contact:   "124325463",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := suite.euc.RegisterNewEmployee(payload)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *EmployeeUseCaseTestSuite) TestGetEmployeeByID_success() {
	suite.erm.On("GetEmployeesByID",  expectEmployee.ID).Return(expectEmployee, nil)

	actualEmployee, err := suite.euc.FindEmployeesByID(expectEmployee.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actualEmployee)
}
func (suite *EmployeeUseCaseTestSuite) TestGetEmployeeByID_emptyId() {
	suite.erm.On("GetEmployeesByID", "").Return(entity.Employee{}, errors.New("id harus diisi"))

	_, err := suite.euc.FindEmployeesByID("")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "id harus diisi", err.Error())
}

func (suite *EmployeeUseCaseTestSuite) TestGetEmployeeByUsername_success() {
	suite.erm.On("GetEmployeesByUsername",  expectEmployee.Username).Return(expectEmployee, nil)

	actualEmployee, err := suite.euc.FindEmployeesByUsername(expectEmployee.Username)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actualEmployee)
}

func (suite *EmployeeUseCaseTestSuite) TestGetEmployeeByUsername_emptyId() {
	suite.erm.On("GetEmployeesByUsername", "").Return(entity.Employee{}, errors.New("username harus diisi"))

	_, err := suite.euc.FindEmployeesByUsername("")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "username harus diisi", err.Error())
}

func (suite *EmployeeUseCaseTestSuite) TestEmployeeForLogin_emptyUsername() {
	suite.erm.On("GetEmployeesByUsernameForLogin", "", "").Return(entity.Employee{}, errors.New("username harus diisi"))

	_, err := suite.euc.FindEmployeForLogin("", "")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "username harus diisi", err.Error())
}

func (suite *EmployeeUseCaseTestSuite) TestEmployeeForLogin_success() {
	suite.erm.On("GetEmployeesByUsernameForLogin",  expectEmployee.Username, expectEmployee.Password).Return(expectEmployee, nil)

	actualEmployee, err := suite.euc.FindEmployeForLogin(expectEmployee.Username, expectEmployee.Password)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectEmployee, actualEmployee)
}

func TestEmployeeUseCaseTestSuite(e *testing.T) {
	suite.Run(e, new(EmployeeUseCaseTestSuite))
}
