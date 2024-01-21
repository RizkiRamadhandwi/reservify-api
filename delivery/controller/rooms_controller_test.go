package controller

import (
	"booking-room-app/entity"
	"booking-room-app/mock/middleware_mock"
	"booking-room-app/mock/usecase_mock"
	"booking-room-app/shared/model"
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

var expectedRoom = entity.Room{
	ID:        "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
	Name:      "Ruang Candradimuka",
	RoomType:  "Ruang Meeting",
	Capacity:  42,
	Status:    "available",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var apiGroup = "/api/v1"
var resource = "/rooms"

var page int = 1
var size int = 5

var expectedPaging = model.Paging{
	Page:        page,
	RowsPerPage: size,
	TotalRows:   2,
	TotalPages:  1,
}

type RoomControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	rum *usecase_mock.RoomUseCaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

func (suite *RoomControllerTestSuite) SetupTest() {
	suite.rum = new(usecase_mock.RoomUseCaseMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	suite.rg = router.Group(apiGroup)
}

func (suite *RoomControllerTestSuite) TestCreateHandler_Success() {
	mockPayload := entity.Room{
		Name:     "Ruang Candradimuka",
		RoomType: "Ruang Meeting",
		Capacity: 42,
		Status:   "available",
	}

	suite.rum.On("RegisterNewRoom", mockPayload).Return(expectedRoom, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	requestBody := `{
        "name": "Ruang Candradimuka",
        "room_type": "Ruang Meeting",
        "capacity": 42,
        "status": "available"
    }`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", apiGroup, resource), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.createHandler(c)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestCreateHandler_BadRequestFailure() {
	mockPayload := entity.Room{}

	suite.rum.On("RegisterNewRoom", &mockPayload).Return(expectedRoom, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", apiGroup, resource), nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.createHandler(c)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestCreateHandler_InternalServerErrorFailure() {
	mockPayload := entity.Room{
		Name:     "Ruang Candradimuka",
		RoomType: "Ruang Meeting",
		Capacity: 42,
		Status:   "available",
	}

	suite.rum.On("RegisterNewRoom", mockPayload).Return(expectedRoom, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	requestBody := `{
        "name": "Ruang Candradimuka",
        "room_type": "Ruang Meeting",
        "capacity": 42,
        "status": "available"
    }`
	request, err := http.NewRequest(http.MethodPost, fmt.Sprintf("%s%s", apiGroup, resource), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.createHandler(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestGetHandler_Success() {
	suite.rum.On("FindRoomByID", "").Return(expectedRoom, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s/%s", apiGroup, resource, expectedRoom.ID), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	handlerFunc.getHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestGetHandler_Failure() {
	suite.rum.On("FindRoomByID", "").Return(expectedRoom, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s/f7casja-881241-12313asd", apiGroup, resource), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	handlerFunc.getHandler(c)

	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestListHandler_Success() {
	mockRooms := []entity.Room{expectedRoom}
	suite.rum.On("FindAllRoom", page, size).Return(mockRooms, expectedPaging, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?page=%d&size=%d", apiGroup, resource, page, size), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockRooms)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestListHandler_EmptyPaginationSuccess() {
	mockRooms := []entity.Room{expectedRoom}
	suite.rum.On("FindAllRoom", page, size).Return(mockRooms, expectedPaging, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, resource), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockRooms)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestListHandler_StatusEmptyPaginationSuccess() {
	mockRooms := []entity.Room{expectedRoom}
	suite.rum.On("FindAllRoomStatus", expectedRoom.Status, page, size).Return(mockRooms, expectedPaging, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?status=%s", apiGroup, resource, expectedRoom.Status), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockRooms)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestListHandler_StatusPaginationSuccess() {
	mockRooms := []entity.Room{expectedRoom}
	suite.rum.On("FindAllRoomStatus", expectedRoom.Status, page, size).Return(mockRooms, expectedPaging, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s?page=%d&size=%d&status=%s", apiGroup, resource, page, size, expectedRoom.Status), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockRooms)
	handlerFunc.listHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestListHandler_BadRequestFailure() {
	mockRooms := []entity.Room{expectedRoom}
	suite.rum.On("FindAllRoom", page, size).Return(mockRooms, model.Paging{}, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", apiGroup, resource), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	c.Set(resource, mockRooms)
	handlerFunc.listHandler(c)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateDetailHandler_Success() {
	mockPayload := entity.Room{
		ID:       "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
		Name:     "Ruang Singasari",
		RoomType: "Ruang Meeting",
		Capacity: 37,
		Status:   "available",
	}

	suite.rum.On("UpdateRoomDetail", mockPayload).Return(expectedRoom, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	requestBody := `{
        "id": "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
        "name": "Ruang Singasari",
        "room_type": "Ruang Meeting",
        "capacity": 37,
        "status": "available"
    }`
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, resource), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.updateDetailHandler(c)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateDetailHandler_BadRequestFailure() {
	mockPayload := entity.Room{}

	suite.rum.On("UpdateRoomDetail", &mockPayload).Return(expectedRoom, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, resource), nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.updateDetailHandler(c)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateDetailHandler_InternalServerErrorFailure() {
	mockPayload := entity.Room{
		ID:       "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
		Name:     "Ruang Singasari",
		RoomType: "Ruang Meeting",
		Capacity: 37,
		Status:   "available",
	}

	suite.rum.On("UpdateRoomDetail", mockPayload).Return(entity.Room{}, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	requestBody := `{
        "id": "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
        "name": "Ruang Singasari",
        "room_type": "Ruang Meeting",
        "capacity": 37,
        "status": "available"
    }`
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, resource), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.updateDetailHandler(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateStatusHandler_Success() {
	mockPayload := entity.Room{
		ID:     "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
		Status: "booked",
	}

	suite.rum.On("UpdateRoomStatus", mockPayload).Return(expectedRoom, nil)

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	requestBody := `{
        "id": "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
        "status": "booked"
    }`
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, resource), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.updateStatusHandler(c)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateStatusHandler_BadRequestFailure() {
	mockPayload := entity.Room{}

	suite.rum.On("UpdateRoomStatus", &mockPayload).Return(expectedRoom, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, resource), nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.updateStatusHandler(c)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomControllerTestSuite) TestUpdateStatusHandler_InternalServerErrorFailure() {
	mockPayload := entity.Room{
		ID:     "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
		Status: "available",
	}

	suite.rum.On("UpdateRoomStatus", mockPayload).Return(entity.Room{}, fmt.Errorf("error"))

	handlerFunc := NewRoomController(suite.rum, suite.amm, suite.rg)
	handlerFunc.Route()

	requestBody := `{
        "id": "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
        "status": "available"
    }`
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s%s", apiGroup, resource), strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request

	handlerFunc.updateStatusHandler(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func TestRoomControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoomControllerTestSuite))
}

