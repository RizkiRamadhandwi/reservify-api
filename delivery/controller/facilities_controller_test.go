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
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type FacilitiesControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	fum *usecase_mock.FacilitiesUseCaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

func (suite *FacilitiesControllerTestSuite) SetupTest() {
	suite.fum = new(usecase_mock.FacilitiesUseCaseMock)
	suite.amm = new(middleware_mock.AuthMiddlewareMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/api/v1")
	rg.Use(suite.amm.RequireToken("userssss"))
	suite.rg = rg
}

var expectedFasilities = entity.Facilities{
	ID:        "1",
	Name:      "This is name",
	Quantity:  10,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

// test list
func (suite *FacilitiesControllerTestSuite) TestListHandler_Success() {
	mockFacility := []entity.Facilities{expectedFasilities}
	mockPaging := model.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}
	suite.fum.On("FindAllFacilities", 1, 5).Return(mockFacility, mockPaging, nil)
	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	handlerFunc.Route()
	request, err := http.NewRequest(http.MethodGet, "/api/v1/facilities?page=1&size=5", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request
	ctx.Set("facilities", mockFacility)
	handlerFunc.listHandler(ctx)

	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestListHandler_Fail() {
	mockFacility := []entity.Facilities{expectedFasilities}
	mockError := errors.New("something went wrong")

	suite.fum.On("FindAllFacilities", 1, 5).Return(mockFacility, model.Paging{}, mockError)
	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/facilities?page=1&size=5", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.listHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

// test gethandler
func (suite *FacilitiesControllerTestSuite) TestGetHandler_Success() {
	// mockID := "1"
	suite.fum.On("FindFacilitiesById", "").Return(expectedFasilities, nil)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/facilities/1", nil)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getHandler(ctx)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestGetHandler_Error() {

	mockError := errors.New("facility not found")
	suite.fum.On("FindFacilitiesById", "").Return(expectedFasilities, mockError)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodGet, "/api/v1/facilities/1", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.getHandler(ctx)

	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestCreateHandler_Success() {
	// Simulate a successful scenario
	mockPayload := entity.Facilities{
		Name:     "This is name",
		Quantity: 10,
	}
	mockFacility := expectedFasilities

	suite.fum.On("RegisterNewFacilities", mockPayload).Return(mockFacility, nil)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	requestBody := `{"name": "This is name", "quantity": 10}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/facilities", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.createHandler(ctx)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestCreateHandler_BadRequest() {
	// Simulate a scenario where binding the JSON payload fails
	mockPayload := entity.Facilities{}
	mockError := errors.New("example error message")

	// Mock the ShouldBindJSON method to return an error
	suite.fum.On("RegisterNewFacilities", &mockPayload).Return(expectedFasilities, mockError)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodPost, "/api/v1/facilities", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.createHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestCreateHandler_InternalServerError() {
	mockPayload := entity.Facilities{
		Name:     "This is name",
		Quantity: 10,
	}
	mockError := errors.New("example error message")

	suite.fum.On("RegisterNewFacilities", mockPayload).Return(expectedFasilities, mockError)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	requestBody := `{"name": "This is name", "quantity": 10}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/facilities", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.createHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestUpdateHandler_Success() {
	// Simulate a successful scenario
	mockPayload := entity.Facilities{
		ID:       "1",
		Name:     "This is name",
		Quantity: 10,
	}
	mockFacility := expectedFasilities

	suite.fum.On("EditFacilities", mockPayload).Return(mockFacility, nil)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	requestBody := `{"id": "1","name": "This is name", "quantity": 10}`
	request, err := http.NewRequest(http.MethodPut, "/api/v1/facilities", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.updateHandler(ctx)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestUpdateHandler_BadRequest() {
	mockPayload := entity.Facilities{}
	mockError := errors.New("example error message")

	suite.fum.On("EditFacilities", &mockPayload).Return(expectedFasilities, mockError)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	request, err := http.NewRequest(http.MethodPut, "/api/v1/facilities", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.updateHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *FacilitiesControllerTestSuite) TestUpdateHandler_NotFound() {
	mockPayload := entity.Facilities{
		ID: "nonexistent_id",
	}
	mockError := errors.New("not found ID " + mockPayload.ID)

	suite.fum.On("EditFacilities", mockPayload).Return(expectedFasilities, mockError)

	handlerFunc := NewFacilitiesController(suite.fum, suite.rg, suite.amm)
	requestBody := `{"id": "nonexistent_id"}`
	request, err := http.NewRequest(http.MethodPut, "/api/v1/facilities", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.updateHandler(ctx)

	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func TestFacilitiesControllerTestSuite(t *testing.T) {
	suite.Run(t, new(FacilitiesControllerTestSuite))
}
