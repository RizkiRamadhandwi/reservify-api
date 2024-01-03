package repository

import (
	"booking-room-app/entity"
	"booking-room-app/shared/model"
	"database/sql"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var expectedRoomFacility = entity.RoomFacility{
	ID:          "1 id did",
	RoomId:      "room id 123",
	FacilityId:  "facility id 1",
	Quantity:    2,
	Description: "Description 1",
	CreatedAt:   time.Time{},
	UpdatedAt:   time.Time{},
}

type RoomFacilityRepositoryTestSuite struct {
	suite.Suite
	mockDB  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    RoomFacilityRepository
}

func (suite *RoomFacilityRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDB = db
	suite.mockSql = mock
	suite.repo = NewRoomFacilityRepository(suite.mockDB)
}

/*  GetRoomFacilityById Success*/
func (suite *RoomFacilityRepositoryTestSuite) TestGetRoomFacilityById_Success() {

	rows := sqlmock.NewRows([]string{"id", "roomId", "facilityId", "quantity", "description", "createdAt", "updatedAt"}).AddRow(
		expectedRoomFacility.ID,
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("SELECT").WithArgs(expectedRoomFacility.ID).WillReturnRows(rows)
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.GetRoomFacilityById(expectedRoomFacility.ID)

	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusOK, actualStatusCode)
	assert.Equal(suite.T(), actualRoomFacility.RoomId, expectedRoomFacility.RoomId)
}

/*  GetRoomFacilityById Fail*/
func (suite *RoomFacilityRepositoryTestSuite) TestGetRoomFacilityById_Fail() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs("invalid id").WillReturnError(fmt.Errorf("just error"))
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.GetRoomFacilityById("invalid id")

	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)

	suite.mockSql.ExpectQuery("SELECT").WithArgs("id that not found").WillReturnError(sql.ErrNoRows)
	actualRoomFacility, actualStatusCode, actualErr = suite.repo.GetRoomFacilityById("id that not found")
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusBadRequest, actualStatusCode)
	assert.Equal(suite.T(), sql.ErrNoRows, actualErr)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* GetQuantityFacilityByID Success */
func (suite *RoomFacilityRepositoryTestSuite) TestGetQuantityFacilityByID_Success() {
	expectedQuantity := 99
	rows := sqlmock.NewRows([]string{"quantity"}).AddRow(
		expectedQuantity,
	)
	suite.mockSql.ExpectQuery("SELECT").WithArgs("1").WillReturnRows(rows)
	actualQuantity, actualStatusCode, actualErr := suite.repo.GetQuantityFacilityByID("1")
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusOK, actualStatusCode)
	assert.Equal(suite.T(), expectedQuantity, actualQuantity)
}

/* GetQuantityFacilityByID Fail */
func (suite *RoomFacilityRepositoryTestSuite) TestGetQuantityFacilityByID_Fail() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs("1").WillReturnError(fmt.Errorf("internal server error"))
	actualQuantity, actualStatusCode, actualErr := suite.repo.GetQuantityFacilityByID("1")
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), 0, actualQuantity)
}

/* GetQuantityFacilityByID Fail No rows */
func (suite *RoomFacilityRepositoryTestSuite) TestGetQuantityFacilityByID_NoRowsFail() {
	suite.mockSql.ExpectQuery("SELECT").WithArgs("1").WillReturnError(sql.ErrNoRows)
	actualQuantity, actualStatusCode, actualErr := suite.repo.GetQuantityFacilityByID("1")
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusBadRequest, actualStatusCode)
	assert.Equal(suite.T(), 0, actualQuantity)
}

/* List Success */
func (suite *RoomFacilityRepositoryTestSuite) TestListRoomFacility_Success() {
	expectedRoomFacility := []entity.RoomFacility{
		{ID: "1", RoomId: "room id 1", FacilityId: "facilit id 1", Quantity: 1, Description: "Description 1", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
		{ID: "2", RoomId: "room id 2", FacilityId: "facilit id 2", Quantity: 2, Description: "Description 2", CreatedAt: time.Time{}, UpdatedAt: time.Time{}},
	}

	rows := sqlmock.NewRows([]string{"id", "roomId", "facilityId", "quantity", "description", "createdAt", "updatedAt"}).AddRow(
		expectedRoomFacility[0].ID,
		expectedRoomFacility[0].RoomId,
		expectedRoomFacility[0].FacilityId,
		expectedRoomFacility[0].Quantity,
		expectedRoomFacility[0].Description,
		expectedRoomFacility[0].CreatedAt,
		expectedRoomFacility[0].UpdatedAt,
	).AddRow(
		expectedRoomFacility[1].ID,
		expectedRoomFacility[1].RoomId,
		expectedRoomFacility[1].FacilityId,
		expectedRoomFacility[1].Quantity,
		expectedRoomFacility[1].Description,
		expectedRoomFacility[1].CreatedAt,
		expectedRoomFacility[1].UpdatedAt,
	)

	expectedPaging := model.Paging{
		Page:        1,
		RowsPerPage: 5,
		TotalRows:   10,
		TotalPages:  2,
	}

	size, page := 5, 1
	suite.mockSql.ExpectQuery("SELECT").WithArgs(size, (page-1)*size).WillReturnRows(rows)
	suite.mockSql.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"totalRows"}).AddRow(10))

	actualRoomFacility, actualPaging, actualErr := suite.repo.ListRoomFacility(page, size)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), expectedPaging, actualPaging)
	assert.Equal(suite.T(), expectedRoomFacility, actualRoomFacility)
}

/* Test ListRoomFacility Fail Select Room-Facility */
func (suite *RoomFacilityRepositoryTestSuite) TestListRoomFacility_QueryFail() {
	size, page := 5, 1
	suite.mockSql.ExpectQuery("SELECT").WithArgs(5, (page-1)*size).WillReturnError(fmt.Errorf("error"))

	actualRoomFacility, actualPaging, actualErr := suite.repo.ListRoomFacility(page, size)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), model.Paging{}, actualPaging)
	assert.Equal(suite.T(), []entity.RoomFacility(nil), actualRoomFacility)
}

/* Test ListRoomFacility Fail Row Scan */
func (suite *RoomFacilityRepositoryTestSuite) TestListRoomFacility_ScanRowFail() {
	// make sql with insuffient field/atribute to cause error in scanner [eliminated room_id]
	rows := sqlmock.NewRows([]string{"id", "facility_id", "quantity", "description", "created_at", "updated_at"}).AddRow(
		expectedRoomFacility.ID,
		//expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("SELECT id, room_id, facility_id, quantity, description, created_at").WithArgs(5, 0).WillReturnRows(rows)

	actualRoomFacility, actualPaging, actualErr := suite.repo.ListRoomFacility(page, size)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), model.Paging{}, actualPaging)
	assert.Equal(suite.T(), []entity.RoomFacility(nil), actualRoomFacility)
}

/* Test ListRoomFacility Fail Count Room-Facility */
func (suite *RoomFacilityRepositoryTestSuite) TestListRoomFacility_CountRoomFacilityFail() {
	rows := sqlmock.NewRows([]string{"id", "room_id", "facility_id", "quantity", "description", "created_at", "updated_at"}).AddRow(
		expectedRoomFacility.ID,
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("SELECT id, room_id, facility_id, quantity, description, created_at").WithArgs(5, 0).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error count room-facility"))
	actualRoomFacility, actualPaging, actualErr := suite.repo.ListRoomFacility(page, size)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), model.Paging{}, actualPaging)
	assert.Equal(suite.T(), []entity.RoomFacility(nil), actualRoomFacility)
}

/* Test CreateRoomFacility Success */
func (suite *RoomFacilityRepositoryTestSuite) TestCreateRoomFacility_Success() {
	suite.mockSql.ExpectBegin().WillReturnError(nil)
	rows := sqlmock.NewRows([]string{"id", "create_at", "updated_at"}).AddRow(
		expectedRoomFacility.ID,
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)
	suite.mockSql.ExpectQuery(`INSERT`).WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
	).WillReturnRows(rows)
	suite.mockSql.ExpectExec(`UPDATE`).WithArgs(5, expectedRoomFacility.FacilityId).WillReturnResult(sqlmock.NewResult(4545, 01101))
	suite.mockSql.ExpectCommit().WillReturnError(nil)
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.CreateRoomFacility(expectedRoomFacility, 5)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusOK, actualStatusCode)
	assert.Equal(suite.T(), expectedRoomFacility, actualRoomFacility)
}

/* Test CreateRoomFacility Fail Begin */
func (suite *RoomFacilityRepositoryTestSuite) TestCreateRoomFacility_BeginTxFail() {
	suite.mockSql.ExpectBegin().WillReturnError(fmt.Errorf("failed to begin transaction tx"))
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.CreateRoomFacility(expectedRoomFacility, 5)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test CreateRoomFacility Fail Insert */
func (suite *RoomFacilityRepositoryTestSuite) TestCreateRoomFacility_InsertFail() {
	suite.mockSql.ExpectBegin().WillReturnError(nil)
	suite.mockSql.ExpectQuery(`INSERT`).WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
	).WillReturnError(fmt.Errorf("failed to insert data"))
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.CreateRoomFacility(expectedRoomFacility, 5)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test CreateRoomFacility Fail Reduce Quantity in Facility */
func (suite *RoomFacilityRepositoryTestSuite) TestCreateRoomFacility_ReduceQuantityFail() {
	suite.mockSql.ExpectBegin().WillReturnError(nil)
	rows := sqlmock.NewRows([]string{"id", "create_at", "updated_at"}).AddRow(
		expectedRoomFacility.ID,
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)
	suite.mockSql.ExpectQuery(`INSERT`).WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
	).WillReturnRows(rows)
	suite.mockSql.ExpectExec("UPDATE").WithArgs(5, expectedRoomFacility.FacilityId).WillReturnError(fmt.Errorf("failed to insert data"))
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.CreateRoomFacility(expectedRoomFacility, 5)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test CreateRoomFacility Fail Commit Transaction */

func (suite *RoomFacilityRepositoryTestSuite) TestCreateRoomFacility_CommitFail() {
	suite.mockSql.ExpectBegin().WillReturnError(nil)
	rows := sqlmock.NewRows([]string{"id", "create_at", "updated_at"}).AddRow(
		expectedRoomFacility.ID,
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)
	suite.mockSql.ExpectQuery(`INSERT`).WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
	).WillReturnRows(rows)
	suite.mockSql.ExpectExec("UPDATE").WithArgs(5, expectedRoomFacility.FacilityId).WillReturnResult(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectCommit().WillReturnError(fmt.Errorf("failed to commit"))
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.CreateRoomFacility(expectedRoomFacility, 5)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test UpdateRoomFacility Success */
func (suite *RoomFacilityRepositoryTestSuite) TestUpdateRoomFacility_Success() {
	newFacilityQuantity := 5

	suite.mockSql.ExpectBegin().WillReturnError(nil)

	rows := sqlmock.NewRows([]string{"create_at", "updated_at"}).AddRow(
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("UPDATE").WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.ID,
	).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE").WithArgs(newFacilityQuantity, expectedRoomFacility.FacilityId).WillReturnResult(sqlmock.NewResult(123, 4567))

	suite.mockSql.ExpectCommit().WillReturnError(nil)

	actualRoomFacility, actualStatusCode, actualErr := suite.repo.UpdateRoomFacility(expectedRoomFacility, newFacilityQuantity)
	assert.Nil(suite.T(), actualErr)
	assert.NoError(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusCreated, actualStatusCode)
	assert.Equal(suite.T(), expectedRoomFacility, actualRoomFacility)
}

/* Test UpdateRoomFacility Failed Begin */
func (suite *RoomFacilityRepositoryTestSuite) TestUpdateRoomFacility_BeginFail() {
	newFacilityQuantity := 5
	suite.mockSql.ExpectBegin().WillReturnError(fmt.Errorf("failed to begin transaction tx"))
	actualRoomFacility, actualStatusCode, actualErr := suite.repo.UpdateRoomFacility(expectedRoomFacility, newFacilityQuantity)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test UpdateRoomFacility Failed Update Room-Facility*/
func (suite *RoomFacilityRepositoryTestSuite) TestUpdateRoomFacility_UpdateRoomFacilityFail() {
	newFacilityQuantity := 5

	suite.mockSql.ExpectBegin().WillReturnError(nil)
	suite.mockSql.ExpectQuery("UPDATE").WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.ID,
	).WillReturnError(fmt.Errorf("failed to update data"))

	actualRoomFacility, actualStatusCode, actualErr := suite.repo.UpdateRoomFacility(expectedRoomFacility, newFacilityQuantity)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test UpdateRoomFacility Failed Update Facility Quantity*/
func (suite *RoomFacilityRepositoryTestSuite) TestUpdateRoomFacility_UpdateFacilityQuantityFail() {
	newFacilityQuantity := 5

	suite.mockSql.ExpectBegin().WillReturnError(nil)

	rows := sqlmock.NewRows([]string{"create_at", "updated_at"}).AddRow(
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("UPDATE").WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.ID,
	).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE").WithArgs(newFacilityQuantity, expectedRoomFacility.FacilityId).WillReturnError(fmt.Errorf("failed to update facility quantity"))

	actualRoomFacility, actualStatusCode, actualErr := suite.repo.UpdateRoomFacility(expectedRoomFacility, newFacilityQuantity)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

/* Test UpdateRoomFacility Failed Commit Transaction*/
func (suite *RoomFacilityRepositoryTestSuite) TestUpdateRoomFacility_CommitFail() {
	newFacilityQuantity := 5

	suite.mockSql.ExpectBegin().WillReturnError(nil)

	rows := sqlmock.NewRows([]string{"create_at", "updated_at"}).AddRow(
		expectedRoomFacility.CreatedAt,
		expectedRoomFacility.UpdatedAt,
	)

	suite.mockSql.ExpectQuery("UPDATE").WithArgs(
		expectedRoomFacility.RoomId,
		expectedRoomFacility.FacilityId,
		expectedRoomFacility.Quantity,
		expectedRoomFacility.Description,
		expectedRoomFacility.ID,
	).WillReturnRows(rows)

	suite.mockSql.ExpectExec("UPDATE").WithArgs(newFacilityQuantity, expectedRoomFacility.FacilityId).WillReturnResult(sqlmock.NewResult(123, 4567))

	suite.mockSql.ExpectCommit().WillReturnError(fmt.Errorf("failed to commit transaction"))

	actualRoomFacility, actualStatusCode, actualErr := suite.repo.UpdateRoomFacility(expectedRoomFacility, newFacilityQuantity)
	assert.NotNil(suite.T(), actualErr)
	assert.Error(suite.T(), actualErr)
	assert.Equal(suite.T(), http.StatusInternalServerError, actualStatusCode)
	assert.Equal(suite.T(), entity.RoomFacility{}, actualRoomFacility)
}

func TestRoomFacilityRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(RoomFacilityRepositoryTestSuite))
}
