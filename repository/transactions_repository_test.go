package repository

import (
	"booking-room-app/config"
	"booking-room-app/entity"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)



var expectedTransactions = entity.Transaction{
	ID:        "1",
    EmployeeId: "1",
    RoomId:    "1",
	RoomFacilities: []entity.RoomFacility{expectedRoomFacilities},
	Status: "pending",
	StartTime:  time.Now(),
	EndTime:  time.Now(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}

var expectedTransaction = []entity.Transaction{
	{ID:        "1",
    EmployeeId: "1",
    RoomId:    "1",
	RoomFacilities: []entity.RoomFacility{
		{
			ID:        "1",
			RoomId:    "1",
			FacilityId: "1",
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			RoomId:    "2",
			FacilityId: "2",
			Quantity:  2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},
	Status: "pending",
	StartTime:  time.Now(),
	EndTime:  time.Now(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
	},
	{ID:        "2",
    EmployeeId: "2",
    RoomId:    "2",
	RoomFacilities: []entity.RoomFacility{
		{
			ID:        "1",
			RoomId:    "1",
			FacilityId: "1",
			Quantity:  1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        "2",
			RoomId:    "2",
			FacilityId: "2",
			Quantity:  2,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
	},
	Status: "pending",
	StartTime:  time.Now(),
	EndTime:  time.Now(),
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
	},
}

var expectedRoomFacilities = entity.RoomFacility {
	ID:        "1",
    RoomId:    "1",
    FacilityId: "1",
    Quantity: 1,
	Description: "test",
    CreatedAt: time.Now(),
    UpdatedAt: time.Now(),
}


type TransactionsRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    TransactionsRepository
}

func (suite *TransactionsRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewTransactionsRepository(suite.mockDb)
}

func (suite *TransactionsRepositoryTestSuite) TestUpdatePermission_Success() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdatePermission)).WithArgs(expectedTransactions.Status, expectedTransactions.ID).WillReturnRows(
	sqlmock.NewRows([]string{"employee_id", "room_id", "description", "start_time", "end_time", "created_at"}).AddRow(
		expectedTransactions.EmployeeId, 
		expectedTransactions.RoomId,
		expectedTransactions.Description,
		expectedTransactions.StartTime,
		expectedTransactions.EndTime,
		expectedTransactions.CreatedAt,
		))

	actual, err := suite.repo.UpdatePemission(expectedTransactions)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expectedTransactions.Description, actual.Description)
}

func (suite *TransactionsRepositoryTestSuite) TestUpdatePermission_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdatePermission)).WithArgs(expectedTransactions.Status, expectedTransactions.ID).WillReturnRows(sqlmock.NewRows([]string{"employee_id"}).AddRow(expectedTransactions.EmployeeId))

	_, err := suite.repo.UpdatePemission(expectedTransactions)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetByEmployeeId_Success() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByEmployeeID)).WithArgs(expectedTransactions.EmployeeId, size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.EmployeeId, 
		expectedTransactions.RoomId, 
		expectedTransactions.Description, 
		expectedTransactions.Status, 
		expectedTransactions.StartTime, 
		expectedTransactions.EndTime, 
		expectedTransactions.CreatedAt, 
		expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"r.id", "r.facility_id", "r.quantity", "r.description", "r.created_at", "r.updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.FacilityId,
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.Description, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployeeIdListTransaction)).WithArgs(expectedTransactions.EmployeeId).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))
    _, _, err := suite.repo.GetTransactionByEmployeId(expectedTransactions.ID, page, size)
    assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetByEmployeeId_Fail() {
	_, _, err := suite.repo.GetTransactionByEmployeId(expectedTransactions.EmployeeId, size, offset)
	suite.Error(err)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByEmployeeID)).WithArgs(expectedTransactions.EmployeeId, size, offset).
		WillReturnRows(sqlmock.NewRows([]string{"employee_id"}).
		AddRow(expectedTransactions.EmployeeId))

    _, _, err = suite.repo.GetTransactionByEmployeId("1", page, size)
    assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetIdEmployeeId_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByEmployeeID)).WithArgs(expectedTransactions.EmployeeId, size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(expectedTransactions.ID, expectedTransactions.EmployeeId, expectedTransactions.RoomId, expectedTransactions.Description, expectedTransactions.Status, expectedTransactions.StartTime, expectedTransactions.EndTime, expectedTransactions.CreatedAt, expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WillReturnRows(sqlmock.NewRows([]string{"r.id", "r.facility_id", "r.quantity", "r.description", "r.created_at", "r.updated_at"}).AddRow(expectedRoomFacilities.ID, expectedRoomFacilities.FacilityId, expectedRoomFacilities.Quantity, expectedRoomFacilities.Description,expectedRoomFacilities.CreatedAt, expectedRoomFacilities.UpdatedAt))

    _, _, err := suite.repo.GetTransactionByEmployeId(expectedTransactions.ID, page, size)
    assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestListEmployeeGetByEmployeeId_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByEmployeeID)).WithArgs(expectedTransactions.EmployeeId, size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(expectedTransactions.ID, expectedTransactions.EmployeeId, expectedTransactions.RoomId, expectedTransactions.Description, expectedTransactions.Status, expectedTransactions.StartTime, expectedTransactions.EndTime, expectedTransactions.CreatedAt, expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"r.id"}).AddRow(expectedRoomFacilities.ID))

    _, _, err := suite.repo.GetTransactionByEmployeId("1", page, size)
    assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
} 

func (suite *TransactionsRepositoryTestSuite) TestListGetByEmployeeIdScanCount_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByEmployeeID)).WithArgs(expectedTransactions.EmployeeId, size, offset).WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.EmployeeId, 
		expectedTransactions.RoomId, 
		expectedTransactions.Description, 
		expectedTransactions.Status, 
		expectedTransactions.StartTime, 
		expectedTransactions.EndTime, 
		expectedTransactions.CreatedAt, 
		expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"r.id", "r.facility_id", "r.quantity", "r.description", "r.created_at", "r.updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.FacilityId,
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.Description, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.GetEmployeeIdListTransaction)).WithArgs(expectedTransactions.EmployeeId).WillReturnError(fmt.Errorf("error"))
    _, _, err := suite.repo.GetTransactionByEmployeId(expectedTransactions.ID, page, size)
	
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
} 

func (suite *TransactionsRepositoryTestSuite) TestGetById_Success() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByID)).WithArgs(expectedTransactions.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(expectedTransactions.ID, expectedTransactions.EmployeeId, expectedTransactions.RoomId, expectedTransactions.Description, expectedTransactions.Status, expectedTransactions.StartTime, expectedTransactions.EndTime, expectedTransactions.CreatedAt, expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{"r.id", "r.facility_id", "r.quantity", "r.description", "r.created_at", "r.updated_at"}).AddRow(expectedRoomFacilities.ID, expectedRoomFacilities.FacilityId, expectedRoomFacilities.Quantity, expectedRoomFacilities.Description, expectedRoomFacilities.CreatedAt, expectedRoomFacilities.UpdatedAt))

    _, err := suite.repo.GetTransactionById(expectedTransactions.ID)
    suite.NoError(err)
}

func (suite *TransactionsRepositoryTestSuite) TestScanGetById_Fail() {
	_, err := suite.repo.GetTransactionById(expectedTransactions.ID)
	suite.Error(err)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByID)).
        WithArgs(expectedTransactions.ID).
        WillReturnRows(sqlmock.NewRows([]string{"id"}).
            AddRow(expectedTransactions.ID))

    _, err = suite.repo.GetTransactionById(expectedTransactions.ID)
    suite.Error(err)
} 

func (suite *TransactionsRepositoryTestSuite) TestGetRoomFacilities_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByID)).
        WithArgs(expectedTransactions.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id", "description", "status", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(expectedTransactions.ID, expectedTransactions.EmployeeId, expectedTransactions.RoomId, expectedTransactions.Description, expectedTransactions.Status, expectedTransactions.StartTime, expectedTransactions.EndTime, expectedTransactions.CreatedAt, expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WillReturnError(fmt.Errorf("error"))

    _, err := suite.repo.GetTransactionById("1")
    assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestGetRoomFacilitiRows_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionByID)).
        WithArgs(expectedTransactions.ID).
		WillReturnRows(sqlmock.NewRows([]string{"id", "employee_id", "room_id", "description", "status", "start_time", "end_time", "created_at", "updated_at"}).
		AddRow(expectedTransactions.ID, expectedTransactions.EmployeeId, expectedTransactions.RoomId, expectedTransactions.Description, expectedTransactions.Status, expectedTransactions.StartTime, expectedTransactions.EndTime, expectedTransactions.CreatedAt, expectedTransactions.UpdatedAt))

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs("1").WillReturnRows(sqlmock.NewRows([]string{ "r.facility_id", "r.quantity", "r.created_at", "r.updated_at"}).AddRow(expectedRoomFacilities.FacilityId, expectedRoomFacilities.Quantity, expectedRoomFacilities.CreatedAt, expectedRoomFacilities.UpdatedAt))

    _, err := suite.repo.GetTransactionById("1")
    assert.Error(suite.T(), err)
	assert.NotNil(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_Success() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at", "updated_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt,
		expectedTransactions.UpdatedAt))
		
		suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertRoomFacility)).WithArgs(
			expectedRoomFacilities.RoomId, 
			expectedRoomFacilities.FacilityId, 
			expectedRoomFacilities.Quantity,
			expectedRoomFacilities.Description).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
				expectedRoomFacilities.ID, 
				expectedRoomFacilities.CreatedAt, 
				expectedRoomFacilities.UpdatedAt))
				
				
	rows = sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectQuantityFacility)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`UPDATE facilities SET quantity = quantity - $1, updated_at = CURRENT_TIMESTAMP WHERE id = $2`)).WithArgs(expectedRoomFacilities.Quantity, expectedFasilities.ID).RowsWillBeClosed().WillReturnRows()
	
	actual, err := suite.repo.Create(expectedTransactions)
	assert.Nil(suite.T(), err)			
    assert.Equal(suite.T(), expectedTransactions.Description, actual.Description)
}

func (suite *TransactionsRepositoryTestSuite) TestGetStatusRoom_Fail() {
	var expectedStatus = "err"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
	
	_, err := suite.repo.Create(expectedTransactions)
	assert.NotNil(suite.T(), err)
    assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_Fail() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime).WillReturnError(fmt.Errorf("error"))

    _, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
    assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesNil() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	var expected = entity.Transaction{
		ID:        "1",
		EmployeeId: "1",
		RoomId:    "1",
		RoomFacilities: nil,
		Status: "pending",
		StartTime:  time.Now(),
		EndTime:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expected.EmployeeId,
        expected.RoomId,
        expected.Description,
        expected.StartTime,
        expected.EndTime).WillReturnRows(
		sqlmock.NewRows([]string{"id", "status", "created_at", "updated_at"}).AddRow(
			expectedTransactions.ID, 
			expectedTransactions.Status,
			expectedTransactions.CreatedAt,
			expectedTransactions.UpdatedAt))

	actual, _ := suite.repo.Create(expected)
    // assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), expected.RoomFacilities, actual.RoomFacilities)
} 

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesScanFaill() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime).WillReturnRows(
		sqlmock.NewRows([]string{"id", "status", "created_at", "updated_at"}).AddRow(
			expectedTransactions.ID, 
			expectedTransactions.Status,
			expectedTransactions.CreatedAt,
			expectedTransactions.UpdatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertRoomFacility)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.Description, 
		expectedRoomFacilities.Description).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
			expectedRoomFacilities.ID, 
			expectedRoomFacilities.CreatedAt, 
			expectedRoomFacilities.UpdatedAt))
		
	_, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesScanQuantityFaill() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at", "updated_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt,
		expectedTransactions.UpdatedAt))
		
		suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertRoomFacility)).WithArgs(
			expectedRoomFacilities.RoomId, 
			expectedRoomFacilities.FacilityId, 
			expectedRoomFacilities.Quantity,
			expectedRoomFacilities.Description).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
				expectedRoomFacilities.ID, 
				expectedRoomFacilities.CreatedAt, 
				expectedRoomFacilities.UpdatedAt))

		rows = sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
		suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectQuantityFacility)).WithArgs("xxx").WillReturnRows(rows)
		
	_, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)	
}

func (suite *TransactionsRepositoryTestSuite) TestCreate_RoomFacilitiesQuantityFaill() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
	var expectedF = entity.Facilities{
		ID:        "1",
		Name:      "This is name",
		Quantity:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at", "updated_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt,
		expectedTransactions.UpdatedAt))
		
		suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertRoomFacility)).WithArgs(
			expectedRoomFacilities.RoomId, 
			expectedRoomFacilities.FacilityId, 
			expectedRoomFacilities.Quantity,
			expectedRoomFacilities.Description).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
				expectedRoomFacilities.ID, 
				expectedRoomFacilities.CreatedAt, 
				expectedRoomFacilities.UpdatedAt))

	rows = sqlmock.NewRows([]string{"quantity"}).AddRow(expectedF.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectQuantityFacility)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)

	// expectedF.Quantity < expectedRoomFacilities.Quantity
	_, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestCreateUpdateFacilityQuantity_Fail() {
	var expectedStatus = "available"
	rows := sqlmock.NewRows([]string{"status"}).AddRow(expectedStatus)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomByID2)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertTransactions)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at", "updated_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt,
		expectedTransactions.UpdatedAt))
		
		suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.InsertRoomFacility)).WithArgs(
			expectedRoomFacilities.RoomId, 
			expectedRoomFacilities.FacilityId, 
			expectedRoomFacilities.Quantity,
			expectedRoomFacilities.Description).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
				expectedRoomFacilities.ID, 
				expectedRoomFacilities.CreatedAt, 
				expectedRoomFacilities.UpdatedAt))

	rows = sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectQuantityFacility)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.UpdateFacilityQuantity)).WithArgs(expectedRoomFacilities.Quantity, expectedFasilities.ID).WillReturnError(fmt.Errorf("error"))

    _, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestUpdate_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`INSERT INTO transactions (employee_id, room_id, description, start_time, end_time, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, status, created_at`)).WithArgs(
        expectedTransactions.EmployeeId,
        expectedTransactions.RoomId,
        expectedTransactions.Description,
        expectedTransactions.StartTime,
        expectedTransactions.EndTime,
		expectedTransactions.UpdatedAt).WillReturnRows(
        sqlmock.NewRows([]string{"id", "status", "created_at"}).AddRow(
		expectedTransactions.ID, 
		expectedTransactions.Status,
		expectedTransactions.CreatedAt))
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta( `INSERT INTO trx_room_facility (room_id, facility_id, quantity, updated_at) VALUES ($1, $2, $3, $4) RETURNING id, created_at, updated_at`)).WithArgs(
		expectedRoomFacilities.RoomId, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.UpdatedAt).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	rows := sqlmock.NewRows([]string{"quantity"}).AddRow(expectedFasilities.Quantity)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`SELECT quantity FROM facilities WHERE id = $1`)).WithArgs(expectedRoomFacilities.FacilityId).WillReturnRows(rows)
		
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(`UPDATE facilities SET quantity = quantity - $1 WHERE id = $2 RETURNING id, created_at, updated_at`)).WithArgs(expectedRoomFacilities.Quantity, expectedFasilities.ID).WillReturnRows(sqlmock.NewRows([]string{"id", "created_at"}).AddRow(expectedFasilities.ID, expectedFasilities.CreatedAt))


    _, err := suite.repo.Create(expectedTransactions)
    assert.NotNil(suite.T(), err)
    assert.Error(suite.T(), err)    
}

func (suite *TransactionsRepositoryTestSuite) TestList_Success() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(
		expectedTransaction[0].ID, 
		expectedTransaction[0].EmployeeId, 
		expectedTransaction[0].RoomId, 
		expectedTransaction[0].Description, 
		expectedTransaction[0].Status, 
		expectedTransaction[0].StartTime, 
		expectedTransaction[0].EndTime,
		expectedTransaction[0].CreatedAt,
		expectedTransaction[0].UpdatedAt,
		)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionList)).WithArgs(size, offset, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs(expectedRoomFacilities.RoomId).WillReturnRows(sqlmock.NewRows([]string{"r.id", "r.facility_id", "r.quantity", "r.description", "r.created_at", "r.updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.Description, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	suite.mockSql.ExpectQuery(`SELECT`).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(2))

	_, _, err := suite.repo.List(page, size, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestList_Fail() {
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionList)).WithArgs(size, offset, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt).WillReturnError(fmt.Errorf("error"))

	_, _, err := suite.repo.List(page, size, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestListScan_Fail() {
	rows := sqlmock.NewRows([]string{"id"}).AddRow(
		expectedTransaction[0].ID, 
		)


	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionList)).WithArgs(size, offset, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt).WillReturnRows(rows)

	_, _, err := suite.repo.List(page, size, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestSelectRoomFacilities_Fail() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(
		expectedTransaction[0].ID, 
		expectedTransaction[0].EmployeeId, 
		expectedTransaction[0].RoomId, 
		expectedTransaction[0].Description, 
		expectedTransaction[0].Status, 
		expectedTransaction[0].StartTime, 
		expectedTransaction[0].EndTime,
		expectedTransaction[0].CreatedAt,
		expectedTransaction[0].UpdatedAt,
		)


	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionList)).WithArgs(size, offset, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt).WillReturnRows(rows)

	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs(expectedRoomFacilities.RoomId).WillReturnError(fmt.Errorf("error"))

	_, _, err := suite.repo.List(page, size, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestScanRoomFacilities_Fail() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(
		expectedTransaction[0].ID, 
		expectedTransaction[0].EmployeeId, 
		expectedTransaction[0].RoomId, 
		expectedTransaction[0].Description, 
		expectedTransaction[0].Status, 
		expectedTransaction[0].StartTime, 
		expectedTransaction[0].EndTime,
		expectedTransaction[0].CreatedAt,
		expectedTransaction[0].UpdatedAt,
		)


	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionList)).WithArgs(size, offset, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs(expectedRoomFacilities.RoomId).WillReturnRows(sqlmock.NewRows([]string{"r.id"}).AddRow(
		expectedRoomFacilities.ID))
	
		_, _, err := suite.repo.List(page, size, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *TransactionsRepositoryTestSuite) TestScanCount_Fail() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "room_id","description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(
		expectedTransaction[0].ID, 
		expectedTransaction[0].EmployeeId, 
		expectedTransaction[0].RoomId, 
		expectedTransaction[0].Description, 
		expectedTransaction[0].Status, 
		expectedTransaction[0].StartTime, 
		expectedTransaction[0].EndTime,
		expectedTransaction[0].CreatedAt,
		expectedTransaction[0].UpdatedAt,
		)


	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectTransactionList)).WithArgs(size, offset, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(regexp.QuoteMeta(config.SelectRoomWithFacilities)).WithArgs(expectedRoomFacilities.RoomId).WillReturnRows(sqlmock.NewRows([]string{"r.id", "r.facility_id", "r.quantity", "r.description", "r.created_at", "r.updated_at"}).AddRow(
		expectedRoomFacilities.ID, 
		expectedRoomFacilities.FacilityId, 
		expectedRoomFacilities.Quantity, 
		expectedRoomFacilities.Description, 
		expectedRoomFacilities.CreatedAt, 
		expectedRoomFacilities.UpdatedAt))

	suite.mockSql.ExpectQuery(`SELECT`).WillReturnError(fmt.Errorf("error"))
	
		_, _, err := suite.repo.List(page, size, expectedTransaction[0].CreatedAt, expectedTransaction[0].CreatedAt)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestTransactionsRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(TransactionsRepositoryTestSuite))
}
	

