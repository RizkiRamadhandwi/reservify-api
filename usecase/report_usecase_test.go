package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/entity/dto"
	"booking-room-app/mock/repo_mock"
	"fmt"
	"testing"
	"time"

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

var expectedReport = []dto.ReportDto{
	{
		ID:         "1",
		EmployeeId: "1",
		Employee:   employee,
		RoomId:     "1",
		Room:       room,
		Status:     "pending",
		StartTime:  time.Now(),
		EndTime:    time.Now().Truncate(time.Second),
		CreatedAt:  time.Now().Truncate(time.Second),
		UpdatedAt:  time.Now().Truncate(time.Second),
	},
}

type ReportUseCaseTestSuite struct {
	suite.Suite
	rrm *repo_mock.ReportRepoMock
	ruc ReportUseCase
}

func (suite *ReportUseCaseTestSuite) SetupTest() {
	suite.rrm = new(repo_mock.ReportRepoMock)
	suite.ruc = NewReportUseCase(suite.rrm)
}

func (suite *ReportUseCaseTestSuite) TestPrintAllReports_DaySuccess() {
	expectedReport[0].StartTime = time.Now().AddDate(0, 0, -1).Truncate(time.Second)
	suite.rrm.On("List", expectedReport[0].StartTime, expectedReport[0].EndTime).Return(expectedReport, nil)

	_, err := suite.ruc.PrintAllReports("day")
	expectedReport[0].StartTime = time.Now()

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *ReportUseCaseTestSuite) TestPrintAllReports_WeekSuccess() {
	expectedReport[0].StartTime = time.Now().AddDate(0, 0, -7).Truncate(time.Second)
	suite.rrm.On("List", expectedReport[0].StartTime, expectedReport[0].EndTime).Return(expectedReport, nil)

	_, err := suite.ruc.PrintAllReports("week")
	expectedReport[0].StartTime = time.Now()

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *ReportUseCaseTestSuite) TestPrintAllReports_MonthSuccess() {
	expectedReport[0].StartTime = time.Now().AddDate(0, -1, 0).Truncate(time.Second)
	suite.rrm.On("List", expectedReport[0].StartTime, expectedReport[0].EndTime).Return(expectedReport, nil)

	_, err := suite.ruc.PrintAllReports("month")
	expectedReport[0].StartTime = time.Now()

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *ReportUseCaseTestSuite) TestPrintAllReports_YearSuccess() {
	expectedReport[0].StartTime = time.Now().AddDate(-1, 0, 0).Truncate(time.Second)
	suite.rrm.On("List", expectedReport[0].StartTime, expectedReport[0].EndTime).Return(expectedReport, nil)

	_, err := suite.ruc.PrintAllReports("year")
	expectedReport[0].StartTime = time.Now()

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *ReportUseCaseTestSuite) TestPrintAllReports_Failure() {
	expectedReport[0].StartTime = time.Now().AddDate(0, 0, -1).Truncate(time.Second)
	suite.rrm.On("List", expectedReport[0].StartTime, expectedReport[0].EndTime).Return(expectedReport, fmt.Errorf("error"))
	expectedReport[0].StartTime = time.Now()

	_, err := suite.ruc.PrintAllReports("day")

	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestReportUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(ReportUseCaseTestSuite))
}
