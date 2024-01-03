package usecase

import (
	"booking-room-app/entity"
	"booking-room-app/entity/dto"
	"booking-room-app/mock/service_mock"
	"booking-room-app/mock/usecase_mock"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthUseCaseTestSuite struct {
	suite.Suite
	aum *usecase_mock.UserUseCaseMock
	jsm *service_mock.JwtServiceMock
	au  AuthUseCase
}

func (suite *AuthUseCaseTestSuite) SetupTest() {
	suite.aum = new(usecase_mock.UserUseCaseMock)
	suite.jsm = new(service_mock.JwtServiceMock)
	suite.au = NewAuthUseCase(suite.aum, suite.jsm)
}

var mockLogin = dto.AuthRequestDto{
	User:     "user1",
	Password: "password",
}
var mockUser = entity.Employee{
	ID:        "1",
	Name:      "neymar",
	Username:  "user1",
	Password:  "password",
	Role:      "admin",
	Division:  "Human Department",
	Position:  "HRD",
	Contact:   "083612",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var mockAuthResponse = dto.AuthResponseDto{
	Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
}

func (suite *AuthUseCaseTestSuite) TestLogin_Success() {
	mockLogin := dto.AuthRequestDto{
		User:     "user1",
		Password: "password",
	}
	mockUser := entity.Employee{
		ID:        "1",
		Name:      "neymar",
		Username:  "user1",
		Password:  "password",
		Role:      "admin",
		Division:  "Human Department",
		Position:  "HRD",
		Contact:   "083612",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	mockAuthResponse := dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}
	suite.aum.On("FindEmployeForLogin", mockLogin.User, mockLogin.Password).Return(mockUser, nil)
	suite.jsm.On("CreateToken", mockUser).Return(mockAuthResponse, nil)
	actual, err := suite.au.Login(mockLogin)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), mockAuthResponse, actual)
}

func (suite *AuthUseCaseTestSuite) TestLogin_Fail() {
	suite.aum.On("FindEmployeForLogin", mockLogin.User, mockLogin.Password).Return(entity.Employee{}, fmt.Errorf("error"))
	_, err := suite.au.Login(mockLogin)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *AuthUseCaseTestSuite) TestLogin_CreateTokenFail() {
	mockLogin := dto.AuthRequestDto{
		User:     "user1",
		Password: "password",
	}
	mockUser := entity.Employee{
		ID:        "1",
		Name:      "neymar",
		Username:  "user1",
		Password:  "password",
		Role:      "admin",
		Division:  "Human Department",
		Position:  "HRD",
		Contact:   "083612",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockAuthResponse := dto.AuthResponseDto{
		Token: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c",
	}
	suite.aum.On("FindEmployeForLogin", mockLogin.User, mockLogin.Password).Return(mockUser, nil)
	suite.jsm.On("CreateToken", mockUser).Return(mockAuthResponse, fmt.Errorf("error"))
	_, err := suite.au.Login(mockLogin)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func TestAuthUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(AuthUseCaseTestSuite))
}
