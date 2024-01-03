package repository

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedFasilities = entity.Facilities{
	ID:        "1",
	Name:      "This is name",
	Quantity:  10,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type FasilitiesRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    FasilitiesRepository
}

func (suite *FasilitiesRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewFasilitesRepository(suite.mockDb)
}

// test create
func (suite *FasilitiesRepositoryTestSuite) TestCreate_Success() {

	rows := sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
		expectedFasilities.ID,
		expectedFasilities.CreatedAt,
		expectedFasilities.UpdatedAt)

	suite.mockSql.ExpectQuery(`INSERT`).WithArgs(
		expectedFasilities.Name,
		expectedFasilities.Quantity).WillReturnRows(rows)

	actual, err := suite.repo.Create(expectedFasilities)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedFasilities.Name, actual.Name)
}

func (suite *FasilitiesRepositoryTestSuite) TestCreate_Fail() {
	suite.mockSql.ExpectQuery(`INSERT`).WithArgs(
		expectedFasilities.Name,
		expectedFasilities.Quantity).WillReturnError(errors.New("error"))

	_, err := suite.repo.Create(expectedFasilities)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

// test get by id
func (suite *FasilitiesRepositoryTestSuite) TestGetById_Success() {

	rows := sqlmock.NewRows([]string{"id", "name", "quantity", "created_at", "updated_at"}).AddRow(
		expectedFasilities.ID, expectedFasilities.Name, expectedFasilities.Quantity, expectedFasilities.CreatedAt, expectedFasilities.UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedFasilities.ID).WillReturnRows(rows)

	actual, err := suite.repo.GetById(expectedFasilities.ID)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedFasilities.Name, actual.Name)
}

func (suite *FasilitiesRepositoryTestSuite) TestGetById_Fail() {
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.GetById("12")
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

// test update
func (suite *FasilitiesRepositoryTestSuite) TestUpdate_Success() {

	rows := sqlmock.NewRows([]string{"created_at", "updated_at"}).AddRow(
		expectedFasilities.CreatedAt, expectedFasilities.UpdatedAt)

	suite.mockSql.ExpectQuery(`UPDATE`).WithArgs(
		expectedFasilities.Name,
		expectedFasilities.Quantity,
		expectedFasilities.ID).WillReturnRows(rows)

	actual, err := suite.repo.UpdateById(expectedFasilities)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedFasilities.Name, actual.Name)
}

func (suite *FasilitiesRepositoryTestSuite) TestUpdate_Fail() {
	suite.mockSql.ExpectQuery(`UPDATE`).WithArgs(
		expectedFasilities.Name,
		expectedFasilities.Quantity,
		expectedFasilities.ID).WillReturnError(errors.New("error"))

	_, err := suite.repo.UpdateById(expectedFasilities)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

// test list
func (suite *FasilitiesRepositoryTestSuite) TestList_Success() {
	page := 1
	size := 10
	expectedFacility := []entity.Facilities{
		{ID: "1", Name: "Facility1", Quantity: 10, CreatedAt: time.Now(), UpdatedAt: time.Now()},
		{ID: "2", Name: "Facility2", Quantity: 15, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	}
	expectedPaging := model.Paging{
		Page:        page,
		RowsPerPage: size,
		TotalRows:   2,
		TotalPages:  1,
	}

	rows := sqlmock.NewRows([]string{"id", "name", "quantity", "created_at", "updated_at"}).
		AddRow("1", "Facility1", 10, time.Now(), time.Now()).
		AddRow("2", "Facility2", 15, time.Now(), time.Now())

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).
		WillReturnRows(sqlmock.NewRows([]string{"total_rows"}).AddRow(2))

	facilities, paging, err := suite.repo.List(page, size)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedFacility, facilities)
	assert.Equal(suite.T(), expectedPaging, paging)
}

func (suite *FasilitiesRepositoryTestSuite) TestList_Fail() {
	page := 1
	size := 10

	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).
		WillReturnError(errors.New("some SQL error"))

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *FasilitiesRepositoryTestSuite) TestList_ScanFail() {
	page := 1
	size := 10

	rows := sqlmock.NewRows([]string{"id", "name", "created_at", "updated_at"}).AddRow(
		expectedFasilities.ID, expectedFasilities.Name, expectedFasilities.CreatedAt, expectedFasilities.UpdatedAt)

	// Mocking the query
	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).WillReturnRows(rows)

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *FasilitiesRepositoryTestSuite) TestList_ScanTotalRows() {
	page := 1
	size := 10

	rows := sqlmock.NewRows([]string{"id", "name", "quantity", "created_at", "updated_at"}).AddRow(
		expectedFasilities.ID, expectedFasilities.Name, expectedFasilities.Quantity, expectedFasilities.CreatedAt, expectedFasilities.UpdatedAt)

	// Mocking the query
	suite.mockSql.ExpectQuery(`SELECT`).
		WithArgs(size, (page-1)*size).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(`SELECT`).
		WillReturnError(errors.New("error"))

	_, _, err := suite.repo.List(page, size)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestFasilitiesRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(FasilitiesRepositoryTestSuite))
}
