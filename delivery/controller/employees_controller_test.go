package controller

import (
	"booking-room-app/entity"
	"booking-room-app/mock/middleware_mock"
	"booking-room-app/mock/usecase_mock"
	"booking-room-app/shared/model"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"

	// "strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type EmployeeControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	eum *usecase_mock.EmployeeUseCaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

var expect = entity.Employee{
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

func (suite *EmployeeControllerTestSuite) SetupTest() {
	suite.eum = new(usecase_mock.EmployeeUseCaseMock)
	suite.amm = new(middleware_mock.AuthMiddlewareMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/api/v1")
	rg.Use(suite.amm.RequireToken("admin"))
	suite.rg = rg
}


func (suite *EmployeeControllerTestSuite) TestCreateHandler_Success(){
	mockPayload := entity.Employee{
		Name:      "a",
		Username:  "a",
		Password:  "a",
		Role:      "admin",
		Division:  "a",
		Position:  "a",
		Contact:   "1",
	}

	mockEmployee := expect
	suite.eum.On("RegisterNewEmployee", mockPayload).Return(mockEmployee, nil)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	handlerFunc.Route()
	requestBody := `{
		"name": "a",
		"username": "a",
		"password": "a",
		"role": "admin",
		"division": "a",
		"position": "a",
		"contact": "1"
	}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/employees", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.createHandler(ctx)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestCreateHandlerPayload_BadRequestFailure(){
	mockPayload := entity.Employee{}
	mockError := errors.New("example error message")

	mockEmployee := expect
	suite.eum.On("RegisterNewEmployee", &mockPayload).Return(mockEmployee, mockError)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodPost, "/api/v1/employees", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.createHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestCreateHandler_InternalServerError(){
	mockPayload := entity.Employee{
		Name:      "a",
		Username:  "a",
		Password:  "a",
		Role:      "admin",
		Division:  "a",
		Position:  "a",
		Contact:   "1",
	}
	mockError := errors.New("example error message")
	mockEmployee := expect
	suite.eum.On("RegisterNewEmployee", mockPayload).Return(mockEmployee, mockError)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	handlerFunc.Route()
	requestBody := `{
		"name": "a",
		"username": "a",
		"password": "a",
		"role": "admin",
		"division": "a",
		"position": "a",
		"contact": "1"
	}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/employees", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.createHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestUpdateHandler_Success(){
	mockPayload := entity.Employee{
		ID: "1",
		Name:      "a",
		Username:  "a",
		Password:  "a",
		Role:      "admin",
		Division:  "a",
		Position:  "a",
		Contact:   "1",
	}

	mockEmployee := expect
	suite.eum.On("UpdateEmployee", mockPayload).Return(mockEmployee, nil)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	handlerFunc.Route()
	requestBody := `{
		"id": "1",
		"name": "a",
		"username": "a",
		"password": "a",
		"role": "admin",
		"division": "a",
		"position": "a",
		"contact": "1"
	}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/employees", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.putHandler(ctx)

	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestUpdateHandler_BadRequest(){
	mockPayload := entity.Employee{}
	mockError := errors.New("example error message")

	mockEmployee := expect
	suite.eum.On("UpdateEmployee", mockPayload).Return(mockEmployee, mockError)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)

	request, err := http.NewRequest(http.MethodPost, "/api/v1/employees", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.putHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestUpdateHandler_NotFound(){
	mockPayload := entity.Employee{
		ID: "error",
	}
	mockError := errors.New("not found ID " + mockPayload.ID)

	mockEmployee := expect
	suite.eum.On("UpdateEmployee", mockPayload).Return(mockEmployee, mockError)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	handlerFunc.Route()
	requestBody := `{"id": "error"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/employees", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.putHandler(ctx)

	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestListHandler_Success(){
	employee := []entity.Employee{expect}
	paging := model.Paging{
		Page: 1,
		RowsPerPage: 1,
		TotalRows: 5,
		TotalPages: 1,
	}
	suite.eum.On("ListAll", 1, 5).Return(employee, paging, nil)
	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	handlerFunc.Route()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees?page=1&size=5", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request
	ctx.Set("employees", employee)
	handlerFunc.ListHandler(ctx)

	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestListHandler_Fail(){
	employee := []entity.Employee{expect}
	errMock := errors.New("something went wrong")

	suite.eum.On("ListAll", 1, 5).Return(employee, model.Paging{}, errMock)
	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees?page=1&size=5", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.ListHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestGetEmployeesByID_Success(){
	suite.eum.On("FindEmployeesByID", "").Return(expect, nil)
	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getByIdHandler(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestGetEmployeesByUsername_Success(){
	suite.eum.On("FindEmployeesByUsername", "").Return(expect, nil)
	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees/username/johndoe", nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getByUsernameHandler(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestGetEmployeesByID_Error(){
	mockError := errors.New("employee not found")
	suite.eum.On("FindEmployeesByID", "").Return(expect, mockError)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees/1", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getByIdHandler(ctx)

	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func (suite *EmployeeControllerTestSuite) TestGetEmployeesByUsername_Error(){
	mockError := errors.New("employee not found")
	suite.eum.On("FindEmployeesByUsername", "").Return(expect, mockError)

	handlerFunc := NewEmployeeController(suite.eum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/employees/username/johndoe", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getByUsernameHandler(ctx)

	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func TestEmployeeControllerTestSuite(e *testing.T){
	suite.Run(e, new(EmployeeControllerTestSuite))
}