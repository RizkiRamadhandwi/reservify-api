package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"booking-room-app/entity"
	"booking-room-app/mock/middleware_mock"
	"booking-room-app/mock/usecase_room_facility"
	"booking-room-app/shared/model"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedRoomFacility = entity.RoomFacility{
	ID:          "id",
	RoomId:      "10be53ca-a3f2-11ee-a506-0242ac120002",
	FacilityId:  "34267130-a3f2-11ee-a506-0242ac120002",
	Quantity:    7,
	Description: "description 1",
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
}

type RoomFacilityControllerTestSuite struct {
	suite.Suite
	rg   *gin.RouterGroup
	rfum *usecase_room_facility.RoomFacilityUseCaseMock
	amm  *middleware_mock.AuthMiddlewareMock
}

func (suite *RoomFacilityControllerTestSuite) SetupTest() {
	suite.rfum = new(usecase_room_facility.RoomFacilityUseCaseMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	suite.rg = router.Group("/api/v1")
}

/* test createRoomFacilityHandler */
func (suite *RoomFacilityControllerTestSuite) TestCreateRoomFacilityHandler_Success() {
	// sending request
	requestBody := `{"roomId": "10be53ca-a3f2-11ee-a506-0242ac120002", "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// call usecase mock
	suite.rfum.On("AddRoomFacilityTransaction").Return(expectedRoomFacility, http.StatusCreated, nil)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [createRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.createRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestCreateRoomFacilityHandler_BindingFail() {
	// sending request
	requestBody := `{"roomId": 10be53ca-a3f2-11ee-a506-0242ac120002, "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [createRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.createRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestCreateRoomFacilityHandler_RequiredFieldFail() {
	// sending request, remove one of required field to cause error
	requestBody := `{"facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [createRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.createRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestCreateRoomFacilityHandler_UuidRoomIdFail() {
	// sending request
	requestBody := `{"roomId": "not uuid", "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [createRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.createRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestCreateRoomFacilityHandler_UuidFacilityIdFail() {
	// sending request
	requestBody := `{"roomId": "10be53ca-a3f2-11ee-a506-0242ac120002", "facilityId": "not uuid", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [createRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.createRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestCreateRoomFacilityHandler_AddRoomFacilityTransactionFail() {
	// sending request
	requestBody := `{"roomId": "10be53ca-a3f2-11ee-a506-0242ac120002", "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// call usecase mock
	suite.rfum.On("AddRoomFacilityTransaction").Return(expectedRoomFacility, http.StatusInternalServerError, fmt.Errorf("fail to add room-facility"))

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [createRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.createRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

/* test listRoomFacilityHandler */
func (suite *RoomFacilityControllerTestSuite) TestListRoomFacilityHandler_Success() {
	// sending request
	request, err := http.NewRequest(http.MethodGet, "/api/v1/roomfacilities?page=1&size=5", nil)

	// assert
	assert.NoError(suite.T(), err)

	mockRoomFacility := []entity.RoomFacility{expectedRoomFacility}
	var mockPaging = model.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}

	// call usecase mock
	suite.rfum.On("FindAllRoomFacility").Return(mockRoomFacility, mockPaging, nil)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [listRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.listRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestListRoomFacilityHandler_FindAllRoomFacilityFail() {
	// sending request
	request, err := http.NewRequest(http.MethodGet, "/api/v1/roomfacilities?page=1&size=5", nil)

	// assert
	assert.NoError(suite.T(), err)

	mockRoomFacility := []entity.RoomFacility{expectedRoomFacility}
	var mockPaging = model.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}

	// call usecase mock
	suite.rfum.On("FindAllRoomFacility").Return(mockRoomFacility, mockPaging, fmt.Errorf("failed to find all room-facility"))

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [listRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.listRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

/* test getRoomFacilityById */
func (suite *RoomFacilityControllerTestSuite) TestGetRoomFacilityById_Success() {
	// sending request
	request, err := http.NewRequest(http.MethodGet, "/api/v1/roomfacilities/1", nil)

	// assert
	assert.NoError(suite.T(), err)

	// call usecase mock
	suite.rfum.On("FindRoomFacilityById").Return(expectedRoomFacility, http.StatusOK, nil)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [getRoomFacilityById]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.getRoomFacilityById(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusOK, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestGetRoomFacilityById_FindRoomFacilityByIdFail() {
	// sending request
	request, err := http.NewRequest(http.MethodGet, "/api/v1/roomfacilities/1", nil)

	// assert
	assert.NoError(suite.T(), err)

	// call usecase mock
	suite.rfum.On("FindRoomFacilityById").Return(expectedRoomFacility, http.StatusNotFound, fmt.Errorf("failed to find room-facility by id"))

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [getRoomFacilityById]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.getRoomFacilityById(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusNotFound, responseRecorder.Code)
}

/* test updateRoomFacilityHandler */
func (suite *RoomFacilityControllerTestSuite) TestUpdateRoomFacilityHandler_Success() {
	// sending request
	requestBody := `{"roomId": "10be53ca-a3f2-11ee-a506-0242ac120002", "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPut, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// call usecase mock
	suite.rfum.On("UpdateRoomFacilityTransaction").Return(expectedRoomFacility, http.StatusCreated, nil)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [updateRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.updateRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestUpdateRoomFacilityHandler_BindingFail() {
	// sending request
	requestBody := `{"roomId": 10be53ca-a3f2-11ee-a506-0242ac120002, "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPut, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [updateRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.updateRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestUpdateRoomFacilityHandler_RequiredFail() {
	// sending request, remove all of required fields to cause error
	requestBody := `{}`
	request, err := http.NewRequest(http.MethodPut, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [updateRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.updateRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *RoomFacilityControllerTestSuite) TestUpdateRoomFacilityHandler_UpdateRoomFacilityFail() {
	// sending request
	requestBody := `{"roomId": "10be53ca-a3f2-11ee-a506-0242ac120002", "facilityId": "34267130-a3f2-11ee-a506-0242ac120002", "quantity": 7, "description": "description 1"}`
	request, err := http.NewRequest(http.MethodPut, "/api/v1/roomfacilities", strings.NewReader(requestBody))

	// assert
	assert.NoError(suite.T(), err)

	// call usecase mock
	suite.rfum.On("UpdateRoomFacilityTransaction").Return(expectedRoomFacility, http.StatusInternalServerError, fmt.Errorf("fail to update room-facility"))

	// record response
	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	// create and call controller function [updateRoomFacilityHandler]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.updateRoomFacilityHandler(ctx)

	// assert
	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

/* test route */
func (suite *RoomFacilityControllerTestSuite) TestRoute_Success() {
	// create and call controller function [Route]
	handlerFunc := NewRoomFacilityController(suite.rfum, suite.rg, suite.amm)
	handlerFunc.Route()
}

func TestRommFacilityControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RoomFacilityControllerTestSuite))
}
