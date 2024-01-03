package controller

import (
	"booking-room-app/entity"
	"booking-room-app/entity/dto"
	"booking-room-app/mock/middleware_mock"
	"booking-room-app/mock/usecase_mock"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
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

type ReportControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	rum *usecase_mock.ReportUseCaseMock
	amm *middleware_mock.AuthMiddlewareMock
}

func (suite *ReportControllerTestSuite) SetupTest() {
	suite.rum = new(usecase_mock.ReportUseCaseMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	suite.rg = router.Group(apiGroup)
}

func (suite *ReportControllerTestSuite) TestDownloadHandler_Success() {
	suite.rum.On("PrintAllReports", "day").Return(expectedReport, nil)

	handlerFunc := NewReportController(suite.rum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/reports/download?range=day", apiGroup), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	handlerFunc.downloadHandler(c)

	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
}

func (suite *ReportControllerTestSuite) TestDownloadHandler_Failure() {
	suite.rum.On("PrintAllReports", "day").Return(expectedReport, fmt.Errorf("error"))

	handlerFunc := NewReportController(suite.rum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/reports/download?range=day", apiGroup), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	handlerFunc.downloadHandler(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func (suite *ReportControllerTestSuite) TestDownloadHandler_EmptyRangeFailure() {
	suite.rum.On("PrintAllReports", "da").Return(expectedReport, fmt.Errorf("error"))

	handlerFunc := NewReportController(suite.rum, suite.rg, suite.amm)
	handlerFunc.Route()

	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/reports/download?range=da", apiGroup), nil)

	responseRecorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(responseRecorder)
	c.Request = request
	handlerFunc.downloadHandler(c)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func TestReportControllerTestSuite(t *testing.T) {
	suite.Run(t, new(ReportControllerTestSuite))
}
