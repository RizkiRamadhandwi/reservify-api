package repository

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
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
var offset int = (page - 1) * size

var expectedPaging = model.Paging{
	Page:        page,
	RowsPerPage: size,
	TotalRows:   2,
	TotalPages:  1,
}

type RoomRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    RoomRepository
}

func (suite *RoomRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewRoomRepository(suite.mockDb)
}

func (suite *RoomRepositoryTestSuite) TestCreate_Success() {
	suite.mockSql.ExpectQuery(`INSERT INTO rooms`).WithArgs(expectedRoom.Name, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.Status).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(expectedRoom.ID, expectedRoom.CreatedAt, expectedRoom.UpdatedAt))

	actual, err := suite.repo.Create(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomRepositoryTestSuite) TestCreate_Failure() {
	suite.mockSql.ExpectQuery(`INSERT INTO rooms`).WillReturnError(fmt.Errorf("error"))

	actual, err := suite.repo.Create(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.Room{}, actual)
}

func (suite *RoomRepositoryTestSuite) TestGet_Success() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "room_type", "capacity", "status", "created_at", "updated_at"}).AddRow(expectedRoom.ID, expectedRoom.Name, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.Status, expectedRoom.CreatedAt, expectedRoom.UpdatedAt))

	actual, err := suite.repo.Get(expectedRoom.ID)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomRepositoryTestSuite) TestGet_Failure() {
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))

	actual, err := suite.repo.Get(expectedRoom.ID)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.Room{}, actual)
}

func (suite *RoomRepositoryTestSuite) TestList_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "room_type", "capacity", "status", "created_at", "updated_at"}).AddRow(expectedRooms[0].ID, expectedRooms[0].Name, expectedRooms[0].RoomType, expectedRooms[0].Capacity, expectedRooms[0].Status, expectedRooms[0].CreatedAt, expectedRooms[0].UpdatedAt).AddRow(expectedRooms[1].ID, expectedRooms[1].Name, expectedRooms[1].RoomType, expectedRooms[1].Capacity, expectedRooms[1].Status, expectedRooms[0].CreatedAt, expectedRooms[1].UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(size, offset).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	actual, paging, err := suite.repo.List(page, size)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRooms[0].Name, actual[0].Name)
	assert.Equal(suite.T(), expectedPaging.Page, paging.Page)
}

func (suite *RoomRepositoryTestSuite) TestList_Failure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(size, offset).WillReturnError(fmt.Errorf("error"))

	_, _, err := suite.repo.List(page, size)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestList_ScanFailure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(expectedRoom.ID, expectedRoom.Name))

	_, _, err := suite.repo.List(page, size)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestList_TotalRowsFailure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "room_type", "capacity", "status", "created_at", "updated_at"}).AddRow(expectedRoom.ID, expectedRoom.Name, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.Status, expectedRoom.CreatedAt, expectedRoom.UpdatedAt))

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(size, offset).WillReturnError(fmt.Errorf("error"))

	_, _, err := suite.repo.List(page, size)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestListStatus_Success() {
	rows := sqlmock.NewRows([]string{"id", "name", "room_type", "capacity", "status", "created_at", "updated_at"}).AddRow(expectedRooms[0].ID, expectedRooms[0].Name, expectedRooms[0].RoomType, expectedRooms[0].Capacity, expectedRooms[0].Status, expectedRooms[0].CreatedAt, expectedRooms[0].UpdatedAt).AddRow(expectedRooms[1].ID, expectedRooms[1].Name, expectedRooms[1].RoomType, expectedRooms[1].Capacity, expectedRooms[1].Status, expectedRooms[0].CreatedAt, expectedRooms[1].UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.Status, size, offset).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.Status).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	actual, paging, err := suite.repo.ListStatus(expectedRoom.Status, page, size)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRooms[0].Name, actual[0].Name)
	assert.Equal(suite.T(), expectedPaging.Page, paging.Page)
}

func (suite *RoomRepositoryTestSuite) TestListStatus_Failure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.Status, size, offset).WillReturnError(fmt.Errorf("error"))

	_, _, err := suite.repo.ListStatus(expectedRoom.Status, page, size)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestListStatus_ScanFailure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.Status, size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(expectedRoom.ID, expectedRoom.Name))

	_, _, err := suite.repo.ListStatus(expectedRoom.Status, page, size)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestListStatus_TotalRowsFailure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.Status, size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "name", "room_type", "capacity", "status", "created_at", "updated_at"}).AddRow(expectedRoom.ID, expectedRoom.Name, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.Status, expectedRoom.CreatedAt, expectedRoom.UpdatedAt))

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedRoom.Status, size, offset).WillReturnError(fmt.Errorf("error"))

	_, _, err := suite.repo.ListStatus(expectedRoom.Status, page, size)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *RoomRepositoryTestSuite) TestUpdate_Success() {
	suite.mockSql.ExpectQuery(`UPDATE`).WithArgs(expectedRoom.ID, expectedRoom.Name, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.Status).WillReturnRows(sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(expectedRoom.CreatedAt, expectedRoom.UpdatedAt))

	actual, err := suite.repo.Update(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Status, actual.Status)
}

func (suite *RoomRepositoryTestSuite) TestUpdate_Failure() {
	suite.mockSql.ExpectQuery(`UPDATE`).WillReturnError(fmt.Errorf("error"))

	actual, err := suite.repo.Update(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.Room{}, actual)
}

func (suite *RoomRepositoryTestSuite) TestUpdateStatus_Success() {
	suite.mockSql.ExpectQuery(`UPDATE`).WithArgs(expectedRoom.ID, expectedRoom.Status).WillReturnRows(sqlmock.NewRows([]string{"name", "room_type", "capacity", "created_at", "updated_at"}).AddRow(expectedRoom.Name, expectedRoom.RoomType, expectedRoom.Capacity, expectedRoom.CreatedAt, expectedRoom.UpdatedAt))

	actual, err := suite.repo.UpdateStatus(expectedRoom)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedRoom.Name, actual.Name)
}

func (suite *RoomRepositoryTestSuite) TestUpdateStatus_Failure() {
	suite.mockSql.ExpectQuery(`UPDATE`).WillReturnError(fmt.Errorf("error"))

	actual, err := suite.repo.UpdateStatus(expectedRoom)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), entity.Room{}, actual)
}

func TestRoomRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomRepositoryTestSuite))
}

