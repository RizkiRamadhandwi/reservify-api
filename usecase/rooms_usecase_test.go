package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/mock/repo_mock"
	"booking-room-app/shared/model"
	"fmt"
	"testing"
	"time"

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

var expectedRooms = []entity.Room{
	{
		ID:        "a3d8e4ef-2e85-4ea5-9509-795f256226c3",
		Name:      "Ruang Candradimuka",
		RoomType:  "Ruang Meeting",
		Capacity:  42,
		Status:    "available",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
	{
		ID:        "e058c04a-7a41-4299-b618-a15e300b3554",
		Name:      "Ruang Bratasena",
		RoomType:  "Ruang Konferensi",
		Capacity:  21,
		Status:    "booked",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

var page int = 1
var size int = 5

var expectedPaging = model.Paging{
	Page:        page,
	RowsPerPage: size,
	TotalRows:   2,
	TotalPages:  1,
}

type RoomUseCaseTestSuite struct {
	suite.Suite
	rrm *repo_mock.RoomRepoMock
	ruc RoomUseCase
}

func (suite *RoomUseCaseTestSuite) SetupTest() {
	suite.rrm = new(repo_mock.RoomRepoMock)
	suite.ruc = NewRoomUseCase(suite.rrm)
}

func (suite *RoomUseCaseTestSuite) TestFindAllRoom() {
	suite.rrm.On("List", page, size).Return(expectedRooms, expectedPaging, nil)

	actual, paging, err := suite.ruc.FindAllRoom(page, size)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRooms[0].Name, actual[0].Name)
	assert.Equal(suite.T(), expectedPaging.Page, paging.Page)
}

func (suite *RoomUseCaseTestSuite) TestFindAllRoomStatus() {
	suite.rrm.On("ListStatus", expectedRoom.Status, page, size).Return(expectedRooms, expectedPaging, nil)

	actual, paging, err := suite.ruc.FindAllRoomStatus(expectedRoom.Status, page, size)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRooms[0].Name, actual[0].Name)
	assert.Equal(suite.T(), expectedPaging.Page, paging.Page)
}

func (suite *RoomUseCaseTestSuite) TestFindroomByID() {
	suite.rrm.On("Get", expectedRoom.ID).Return(expectedRoom, nil)

	actual, err := suite.ruc.FindRoomByID(expectedRoom.ID)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_Success() {
	suite.rrm.On("Create", expectedRoom).Return(expectedRoom, nil)

	actual, err := suite.ruc.RegisterNewRoom(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_EmptyStatusSuccess() {
	var expectedPayload entity.Room = expectedRoom
	expectedRoom.Status = ""

	suite.rrm.On("Create", expectedPayload).Return(expectedPayload, nil)

	_, err := suite.ruc.RegisterNewRoom(expectedRoom)
	expectedRoom.Status = "available"

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_Failure() {
	suite.rrm.On("Create", expectedRoom).Return(entity.Room{}, fmt.Errorf("error"))

	_, err := suite.ruc.RegisterNewRoom(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestRegisterNewRoom_EmptyFieldFailure() {
	expectedRoom.Name = ""
	_, err := suite.ruc.RegisterNewRoom(expectedRoom)
	expectedRoom.Name = "Ruang Candradimuka"

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomDetail_Success() {
	suite.rrm.On("Update", expectedRoom).Return(expectedRoom, nil)

	actual, err := suite.ruc.UpdateRoomDetail(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomDetail_Failure() {
	suite.rrm.On("Update", expectedRoom).Return(entity.Room{}, fmt.Errorf("error"))

	_, err := suite.ruc.UpdateRoomDetail(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomDetail_EmptyFieldFailure() {
	expectedRoom.Name = ""
	_, err := suite.ruc.UpdateRoomDetail(expectedRoom)
	expectedRoom.Name = "Ruang Candradimuka"

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomDetail_EmptyStatusSuccess() {
	var expectedPayload entity.Room = expectedRoom
	expectedRoom.Status = ""

	suite.rrm.On("Update", expectedPayload).Return(expectedPayload, nil)

	_, err := suite.ruc.UpdateRoomDetail(expectedRoom)
	expectedRoom.Status = "available"

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomStatus_Success() {
	suite.rrm.On("UpdateStatus", expectedRoom).Return(expectedRoom, nil)

	actual, err := suite.ruc.UpdateRoomStatus(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomStatus_Failure() {
	suite.rrm.On("UpdateStatus", expectedRoom).Return(entity.Room{}, fmt.Errorf("error"))

	_, err := suite.ruc.UpdateRoomStatus(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomStatus_EmptyFieldFailure() {
	expectedRoom.ID = ""
	_, err := suite.ruc.UpdateRoomStatus(expectedRoom)
	expectedRoom.ID = "a3d8e4ef-2e85-4ea5-9509-795f256226c3"

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomUseCaseTestSuite) TestUpdateRoomStatus_EmptyStatusSuccess() {
	var expectedPayload entity.Room = expectedRoom
	expectedRoom.Status = ""

	suite.rrm.On("UpdateStatus", expectedPayload).Return(expectedPayload, nil)

	_, err := suite.ruc.UpdateRoomStatus(expectedRoom)
	expectedRoom.Status = "available"

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func TestRoomUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(RoomUseCaseTestSuite))
}
