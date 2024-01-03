package repository

import (
	"booking-room-app/entity"
	"booking-room-app/entity/dto"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var employee = entity.Employee{
	Name:     "Admin Enigma",
	Username: "adminenigma",
	Division: "IT",
	Position: "Administration Officer",
	Contact:  "088999000007",
}

var room = entity.Room{
	Name:     "Ruang Candradimuka",
	RoomType: "Meeting Room",
	Capacity: 21,
}

var expectedReport = dto.ReportDto{
	ID:         "1",
	EmployeeId: "1",
	Employee:   employee,
	RoomId:     "1",
	Room:       room,
	Status:     "pending",
	StartTime:  time.Now(),
	EndTime:    time.Now(),
	CreatedAt:  time.Now(),
	UpdatedAt:  time.Now(),
}

var expectedRoomFacilityty = dto.RoomFacilityDto{
	FacilityID: "1",
	Name:       "LED Proyektor",
	Quantity:   2,
}

type ReportRepositoryTestSuite struct {
	suite.Suite
	mockDb  *sql.DB
	mockSql sqlmock.Sqlmock
	repo    ReportRepository
}

func (suite *ReportRepositoryTestSuite) SetupTest() {
	db, mock, _ := sqlmock.New()
	suite.mockDb = db
	suite.mockSql = mock
	suite.repo = NewReportRepository(suite.mockDb)
}

func (suite *ReportRepositoryTestSuite) TestList_Success() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "username", "division", "position", "contact", "room_id", "name", "room_type", "capacity", "description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(expectedReport.ID, expectedReport.EmployeeId, expectedReport.Employee.Name, expectedReport.Employee.Username, expectedReport.Employee.Division, expectedReport.Employee.Position, expectedReport.Employee.Contact, expectedReport.RoomId, expectedReport.Room.Name, expectedReport.Room.RoomType, expectedReport.Room.Capacity, expectedReport.Description, expectedReport.Status, expectedReport.StartTime, expectedReport.EndTime, expectedReport.CreatedAt, expectedReport.UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.StartTime, expectedReport.EndTime).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.RoomId).WillReturnRows(sqlmock.NewRows([]string{"facility_id", "name", "quantity"}).AddRow(expectedRoomFacilityty.FacilityID, expectedRoomFacilityty.Name, expectedRoomFacilityty.Quantity))

	actual, err := suite.repo.List(expectedReport.StartTime, expectedReport.EndTime)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedReport.ID, actual[0].ID)
}

func (suite *ReportRepositoryTestSuite) TestList_Failure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.StartTime, expectedReport.EndTime).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.List(expectedReport.StartTime, expectedReport.EndTime)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *ReportRepositoryTestSuite) TestList_ScanFailure() {
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.StartTime, expectedReport.EndTime).WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedReport.ID))

	_, err := suite.repo.List(expectedReport.StartTime, expectedReport.EndTime)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *ReportRepositoryTestSuite) TestList_RoomFacilityFailure() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "username", "division", "position", "contact", "room_id", "name", "room_type", "capacity", "description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(expectedReport.ID, expectedReport.EmployeeId, expectedReport.Employee.Name, expectedReport.Employee.Username, expectedReport.Employee.Division, expectedReport.Employee.Position, expectedReport.Employee.Contact, expectedReport.RoomId, expectedReport.Room.Name, expectedReport.Room.RoomType, expectedReport.Room.Capacity, expectedReport.Description, expectedReport.Status, expectedReport.StartTime, expectedReport.EndTime, expectedReport.CreatedAt, expectedReport.UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.StartTime, expectedReport.EndTime).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.RoomId).WillReturnError(fmt.Errorf("error"))

	_, err := suite.repo.List(expectedReport.StartTime, expectedReport.EndTime)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *ReportRepositoryTestSuite) TestList_ScanRoomFacilityFailure() {
	rows := sqlmock.NewRows([]string{"id", "employee_id", "name", "username", "division", "position", "contact", "room_id", "name", "room_type", "capacity", "description", "status", "start_time", "end_time", "created_at", "updated_at"}).AddRow(expectedReport.ID, expectedReport.EmployeeId, expectedReport.Employee.Name, expectedReport.Employee.Username, expectedReport.Employee.Division, expectedReport.Employee.Position, expectedReport.Employee.Contact, expectedReport.RoomId, expectedReport.Room.Name, expectedReport.Room.RoomType, expectedReport.Room.Capacity, expectedReport.Description, expectedReport.Status, expectedReport.StartTime, expectedReport.EndTime, expectedReport.CreatedAt, expectedReport.UpdatedAt)

	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.StartTime, expectedReport.EndTime).WillReturnRows(rows)
	suite.mockSql.ExpectQuery(`SELECT`).WithArgs(expectedReport.RoomId).WillReturnRows(sqlmock.NewRows([]string{"facility_id"}).AddRow(expectedRoomFacilityty.FacilityID))

	_, err := suite.repo.List(expectedReport.StartTime, expectedReport.EndTime)

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestReportRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(ReportRepositoryTestSuite))
}
