package controller

import (
	"booking-room-app/entity/dto"
	"booking-room-app/mock/usecase_mock"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// loginHandler

type AuthControllerTestSuite struct {
	suite.Suite
	rg  *gin.RouterGroup
	aum *usecase_mock.AuthUseCaseMock
}

func (suite *AuthControllerTestSuite) SetupTest() {
	suite.aum = new(usecase_mock.AuthUseCaseMock)
	router := gin.Default()
	gin.SetMode(gin.TestMode)
	rg := router.Group("/api/v1")
	suite.rg = rg
}

func (suite *AuthControllerTestSuite) TestLoginHandler_Success() {
	mockLogin := dto.AuthRequestDto{
		User:     "user1",
		Password: "password",
	}

	mockAuthResponse := dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}

	suite.aum.On("Login", mockLogin).Return(mockAuthResponse, nil)

	handlerFunc := NewAuthController(suite.aum, suite.rg)
	handlerFunc.Route()
	requestBody := `{"username": "user1", "password": "password"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.loginHandler(ctx)

	assert.Equal(suite.T(), http.StatusCreated, responseRecorder.Code)
}

func (suite *AuthControllerTestSuite) TestLoginHandler_BadRequest() {
	// Simulate a scenario where binding the JSON payload fails
	mockLogin := dto.AuthRequestDto{}
	mockAuthResponse := dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}
	mockError := errors.New("example error message")

	// Mock the ShouldBindJSON method to return an error
	suite.aum.On("Login", &mockLogin).Return(mockAuthResponse, mockError)

	handlerFunc := NewAuthController(suite.aum, suite.rg)
	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", nil)
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.loginHandler(ctx)

	assert.Equal(suite.T(), http.StatusBadRequest, responseRecorder.Code)
}

func (suite *AuthControllerTestSuite) TestLoginHandler_InternalServerError() {
	mockLogin := dto.AuthRequestDto{
		User:     "user1",
		Password: "password",
	}

	mockAuthResponse := dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}

	mockError := errors.New("example error message")

	// Mock the ShouldBindJSON method to return an error
	suite.aum.On("Login", mockLogin).Return(mockAuthResponse, mockError)

	handlerFunc := NewAuthController(suite.aum, suite.rg)
	requestBody := `{"username": "user1", "password": "password"}`
	request, err := http.NewRequest(http.MethodPost, "/api/v1/auth/login", strings.NewReader(requestBody))
	assert.NoError(suite.T(), err)

	responseRecorder := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(responseRecorder)
	ctx.Request = request

	handlerFunc.loginHandler(ctx)

	assert.Equal(suite.T(), http.StatusInternalServerError, responseRecorder.Code)
}

func TestAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthControllerTestSuite))
}
