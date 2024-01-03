package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/mock/repo_room_facility"
	"booking-room-app/shared/model"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedRoomFacility = entity.RoomFacility{
	ID:          "id",
	RoomId:      "room id",
	FacilityId:  "facility id",
	Quantity:    7,
	Description: "Description 1",
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
}

type RoomFacilityUseCaseTestSuite struct {
	suite.Suite
	rfrm *repo_room_facility.RoomFacilityRepoMock
	rfuc RoomFacilityUsecase
}

func (suite *RoomFacilityUseCaseTestSuite) SetupTest() {
	suite.rfrm = new(repo_room_facility.RoomFacilityRepoMock)
	suite.rfuc = NewRoomFacilityUsecase(suite.rfrm)
}

/* test AddRoomFacilityTransaction success*/
func (suite *RoomFacilityUseCaseTestSuite) TestAddRoomFacilityTransaction_Success() {
	availableFacilityQuantity := 10
	suite.rfrm.On("GetQuantityFacilityByID").Return(availableFacilityQuantity, http.StatusOK, nil)
	suite.rfrm.On("CreateRoomFacility").Return(expectedRoomFacility, http.StatusCreated, nil)
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.AddRoomFacilityTransaction(expectedRoomFacility)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusCreated, actualStatusCode)
	assert.Equal(suite.T(), expectedRoomFacility, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestAddRoomFacilityTransaction_GetQuantityFail() {
	availableFacilityQuantity := 10
	suite.rfrm.On("GetQuantityFacilityByID").Return(availableFacilityQuantity, http.StatusBadRequest, fmt.Errorf("failed to get facility quntity"))
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.AddRoomFacilityTransaction(expectedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusBadRequest, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestAddRoomFacilityTransaction_InsuffientQuantityFail() {
	availableFacilityQuantity := 0
	suite.rfrm.On("GetQuantityFacilityByID").Return(availableFacilityQuantity, http.StatusOK, nil)
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.AddRoomFacilityTransaction(expectedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusBadRequest, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestAddRoomFacilityTransaction_CreateRoomFacilityFail() {
	availableFacilityQuantity := 10
	suite.rfrm.On("GetQuantityFacilityByID").Return(availableFacilityQuantity, http.StatusOK, nil)
	suite.rfrm.On("CreateRoomFacility").Return(expectedRoomFacility, http.StatusInternalServerError, fmt.Errorf("failed to create room-facility"))
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.AddRoomFacilityTransaction(expectedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* test FindAllRoomFacility success */
func (suite *RoomFacilityUseCaseTestSuite) TestFindAllRoomFacility_Success() {
	suite.rfrm.On("ListRoomFacility").Return([]entity.RoomFacility{}, model.Paging{}, nil)
	actualRoomFacility, actualPaging, actualErr := suite.rfuc.FindAllRoomFacility(0, 0)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), model.Paging{}, actualPaging)
	assert.Equal(suite.T(), []entity.RoomFacility{}, actualRoomFacility)
}

/* Test FindRoomFacilityById Success*/
func (suite *RoomFacilityUseCaseTestSuite) TestFindRoomFacilityById_Success() {
	suite.rfrm.On("GetRoomFacilityById").Return(expectedRoomFacility, http.StatusOK, nil)
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.FindRoomFacilityById(expectedRoomFacility.ID)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusOK, actualStatusCode)
	assert.Equal(suite.T(), expectedRoomFacility, actualRoomFacility)
}

/* Test UpdateRoomFacilityTransaction Success*/
func (suite *RoomFacilityUseCaseTestSuite) TestUpdateRoomFacilityTransaction_Success() {
	var postedRoomFacility = entity.RoomFacility{
		RoomId:      "",
		FacilityId:  "",
		Quantity:    0,
		Description: "",
	}
	var oldRoomFacility = entity.RoomFacility{
		ID:          "id",
		RoomId:      "room id",
		FacilityId:  "facility id",
		Quantity:    7,
		Description: "Description 1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	suite.rfrm.On("GetRoomFacilityById").Return(oldRoomFacility, http.StatusOK, nil)
	suite.rfrm.On("UpdateRoomFacility").Return(expectedRoomFacility, http.StatusOK, nil)
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.UpdateRoomFacilityTransaction(postedRoomFacility)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusOK, actualStatusCode)
	assert.Equal(suite.T(), expectedRoomFacility, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestUpdateRoomFacilityTransaction_GetRoomFacilityByIDFail() {
	suite.rfrm.On("GetRoomFacilityById").Return(entity.RoomFacility{}, http.StatusInternalServerError, fmt.Errorf("failed to get room-facility by id"))
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.UpdateRoomFacilityTransaction(expectedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestUpdateRoomFacilityTransaction_GetQuantityFacilityByIDFail() {
	var postedRoomFacility = entity.RoomFacility{
		RoomId:      "",
		FacilityId:  "",
		Quantity:    1,
		Description: "",
	}
	var oldRoomFacility = entity.RoomFacility{
		ID:          "id",
		RoomId:      "room id",
		FacilityId:  "facility id",
		Quantity:    7,
		Description: "Description 1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	suite.rfrm.On("GetRoomFacilityById").Return(oldRoomFacility, http.StatusOK, nil)
	suite.rfrm.On("GetQuantityFacilityByID").Return(0, http.StatusBadRequest, fmt.Errorf("failed to get quantity facility by id"))
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.UpdateRoomFacilityTransaction(postedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusBadRequest, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestUpdateRoomFacilityTransaction_InsufficientQuantityFail() {
	var postedRoomFacility = entity.RoomFacility{
		RoomId:      "",
		FacilityId:  "",
		Quantity:    8,
		Description: "",
	}
	var oldRoomFacility = entity.RoomFacility{
		ID:          "id",
		RoomId:      "room id",
		FacilityId:  "facility id",
		Quantity:    7,
		Description: "Description 1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	suite.rfrm.On("GetRoomFacilityById").Return(oldRoomFacility, http.StatusOK, nil)
	suite.rfrm.On("GetQuantityFacilityByID").Return(0, http.StatusOK, nil)
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.UpdateRoomFacilityTransaction(postedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusBadRequest, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func (suite *RoomFacilityUseCaseTestSuite) TestUpdateRoomFacilityTransaction_UpdteRoomFacilityFail() {
	var postedRoomFacility = entity.RoomFacility{
		RoomId:      "",
		FacilityId:  "",
		Quantity:    0,
		Description: "",
	}
	var oldRoomFacility = entity.RoomFacility{
		ID:          "id",
		RoomId:      "room id",
		FacilityId:  "facility id",
		Quantity:    7,
		Description: "Description 1",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}
	suite.rfrm.On("GetRoomFacilityById").Return(oldRoomFacility, http.StatusOK, nil)
	suite.rfrm.On("UpdateRoomFacility").Return(entity.RoomFacility{}, http.StatusInternalServerError, fmt.Errorf("failed to update room-facility"))
	actualRoomFacility, actualStatusCode, actualErr := suite.rfuc.UpdateRoomFacilityTransaction(postedRoomFacility)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func TestRoomFacilityUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoomFacilityUseCaseTestSuite))
}
