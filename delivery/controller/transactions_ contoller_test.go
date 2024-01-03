package controller

import (
	"booking-room-app/entity"
	"booking-room-app/mock/middleware_mock"
	"booking-room-app/mock/usecase_mock"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedTransactions = entity.Transaction{
	ID:             "1",
	EmployeeId:     "1",
	RoomId:         "1",
	RoomFacilities: nil,
	Description:    "Test",
	Status:         "pending",
	StartTime:      time.Date(2023, time.December, 25, 12, 0, 0, 0, time.UTC),
	EndTime:        time.Date(2023, time.December, 25, 12, 0, 0, 0, time.UTC),
	CreatedAt:      time.Now(),
	UpdatedAt:      time.Now(),
}

var transactionsPoint = "/transactions"

type TransactionsControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	tum *usecase_mock.TransactionsUseCaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

func (suite *TransactionsControllerTestSuite) SetupTest() {
	suite.tum = new(usecase_mock.TransactionsUseCaseMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	suite.rg = router.Group(apiGroup)
}

func (suite *TransactionsControllerTestSuite) TestCreateHandler_Success() {
	mockPayload := entity.Transaction{
		ID:          "1",
		EmployeeId:  "1",
		RoomId:      "1",
		Description: "Test",
		StartTime:   time.Date(2023, time.December, 25, 12, 0, 0, 0, time.UTC),
		EndTime:     time.Date(2023, time.December, 25, 15, 0, 0, 0, time.UTC),
	}

	suite.tum.On("RequestNewBookingRooms", mockPayload).Return(expectedTransactions, nil)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	requestBody := `{
        "id": "1",
        "employeeId": "1",
        "roomId": "1",
        "description": "Test",
		"startTime": "2023-12-25T12:00:00Z",
        "endTime": "2023-12-25T15:00:00Z"
		}`

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.createHandler(c)
	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestCreateHandler_fail() {
	mockPayload := entity.Transaction{}

	suite.tum.On("RequestNewBookingRooms", &mockPayload).Return(expectedTransactions, fmt.Errorf("error"))

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.createHandler(c)
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestCreateHandler_InternalServerErrorFailure() {
	mockPayload := entity.Transaction{
		ID:          "1",
		EmployeeId:  "1",
		RoomId:      "1",
		Description: "Test",
		StartTime:   time.Date(2023, time.December, 25, 12, 0, 0, 0, time.UTC),
		EndTime:     time.Date(2023, time.December, 25, 15, 0, 0, 0, time.UTC),
	}

	suite.tum.On("RequestNewBookingRooms", mockPayload).Return(expectedTransactions, fmt.Errorf("error"))

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	requestBody := `{
        "id": "1",
        "employeeId": "1",
        "roomId": "1",
        "description": "Test",
		"startTime": "2023-12-25T12:00:00Z",
        "endTime": "2023-12-25T15:00:00Z"
		}`

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.createHandler(c)
	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestListHandler_Success() {
	mockTransactions := []entity.Transaction{expectedTransactions}
	suite.tum.On("FindAllTransactions", page, size, time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(3000, time.December, 31, 0, 0, 0, 0, time.UTC)).Return(mockTransactions, expectedPaging, nil)
	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockTransactions)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestListHandler_StartDateTimeBadRequest() {
	mockTransactions := []entity.Transaction{expectedTransactions}
	suite.tum.On("FindAllTransactions", page, size, time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(3000, time.December, 31, 0, 0, 0, 0, time.UTC)).Return(mockTransactions, expectedPaging, nil)
	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?startDate=err", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockTransactions)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestListHandler_EndDateTimeBadRequest() {
	mockTransactions := []entity.Transaction{expectedTransactions}
	suite.tum.On("FindAllTransactions", page, size, time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(3000, time.December, 31, 0, 0, 0, 0, time.UTC)).Return(mockTransactions, expectedPaging, nil)
	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?endDate=err", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockTransactions)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestListHandler_PaginationBadRequest() {
	mockTransactions := []entity.Transaction{expectedTransactions}
	mockError := errors.New("something went wrong")
	suite.tum.On("FindAllTransactions", page, size, time.Date(1000, time.January, 1, 0, 0, 0, 0, time.UTC), time.Date(3000, time.December, 31, 0, 0, 0, 0, time.UTC)).Return(mockTransactions, expectedPaging, mockError)
	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?page=err&size=err", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()

	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockTransactions)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestgetTransactionById_Success() {
	// mockID := "1"
	suite.tum.On("FindTransactionsById", "").Return(expectedTransactions, nil)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getTransactionById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestGetTransactionById_Fail() {
	mockError := errors.New("transaction not found")
	suite.tum.On("FindTransactionsById", "").Return(expectedTransactions, mockError)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getTransactionById(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestgetTransactionByEmployeeId_Success() {
	mockTransactions := []entity.Transaction{expectedTransactions}
	suite.tum.On("FindTransactionsByEmployeeId", "").Return(mockTransactions, expectedPaging, nil)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getTransactionByEmployeeId(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestgetTransactionByEmployeeId_Fail() {
	mockTransactions := []entity.Transaction{expectedTransactions}

	mockError := errors.New("transaction not found")
	suite.tum.On("FindTransactionsByEmployeeId", "").Return(mockTransactions, expectedPaging, mockError)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getTransactionByEmployeeId(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestUpdateHandler_Success() {
	mockPayload := entity.Transaction{
		ID:     "1",
		Status: "accepted",
	}

	suite.tum.On("AccStatusBooking", mockPayload).Return(mockPayload, nil)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	requestBody := `{"id": "1","status": "accepted"}`
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)
	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockPayload)
	handlerFunc.updateStatusHandler(c)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestUpdateHandler_BadRequest() {
	mockPayload := entity.Transaction{}
	mockError := errors.New("example error message")

	suite.tum.On("AccStatusBooking", &mockPayload).Return(mockPayload, mockError)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockPayload)
	handlerFunc.updateStatusHandler(c)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *TransactionsControllerTestSuite) TestUpdateHandler_NotFound() {
	mockPayload := entity.Transaction{
		ID: "nonexistent_id",
	}
	mockError := errors.New("not found ID " + mockPayload.ID)

	suite.tum.On("AccStatusBooking", mockPayload).Return(mockPayload, mockError)

	handlerFunc := NewTransactionsController(suite.tum, suite.rg, suite.amm)
	requestBody := `{"id": "nonexistent_id"}`
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, transactionsPoint), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)
	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockPayload)
	handlerFunc.updateStatusHandler(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func TestTransactionControllerTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsControllerTestSuite))
}

