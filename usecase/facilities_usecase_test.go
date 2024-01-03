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

var expectedFasilities = entity.Facilities{
	ID:        "1",
	Name:      "This is name",
	Quantity:  10,
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

type FacilitiesUseCaseTestSuite struct {
	suite.Suite
	frm *repo_mock.FacilitiesRepoMock
	fuc FacilitiesUseCase
}

func (suite *FacilitiesUseCaseTestSuite) SetupTest() {
	suite.frm = new(repo_mock.FacilitiesRepoMock)
	suite.fuc = NewFacilitiesUseCase(suite.frm)
}

// test create
func (suite *FacilitiesUseCaseTestSuite) TestRegisterNewFacilities_Success() {
	suite.frm.On("Create", expectedFasilities).Return(expectedFasilities, nil)
	actual, err := suite.fuc.RegisterNewFacilities(expectedFasilities)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedFasilities.Name, actual.Name)
}

func (suite *FacilitiesUseCaseTestSuite) TestRegisterNewFacilities_EmptyField() {
	payloadMock := entity.Facilities{
		ID:        "1",
		Name:      "",
		Quantity:  10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := suite.fuc.RegisterNewFacilities(payloadMock)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *FacilitiesUseCaseTestSuite) TestRegisterNewFacilities_Fail() {
	suite.frm.On("Create", expectedFasilities).Return(entity.Facilities{}, fmt.Errorf("error"))
	_, err := suite.fuc.RegisterNewFacilities(expectedFasilities)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

// test edit
func (suite *FacilitiesUseCaseTestSuite) TestEditFacilities_Success() {
	suite.frm.On("UpdateById", expectedFasilities).Return(expectedFasilities, nil)
	actual, err := suite.fuc.EditFacilities(expectedFasilities)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), expectedFasilities.Name, actual.Name)
}

func (suite *FacilitiesUseCaseTestSuite) TestEditFacilities_EmptyField() {
	payloadMock := entity.Facilities{
		ID:        "1",
		Name:      "",
		Quantity:  10,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := suite.fuc.EditFacilities(payloadMock)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

func (suite *FacilitiesUseCaseTestSuite) TestEditFacilities_Fail() {
	suite.frm.On("UpdateById", expectedFasilities).Return(entity.Facilities{}, fmt.Errorf("error"))
	_, err := suite.fuc.EditFacilities(expectedFasilities)
	assert.NotNil(suite.T(), err)
	assert.Error(suite.T(), err)
}

// test find all
func (suite *FacilitiesUseCaseTestSuite) TestFindAllFacilities_Success() {
	mockData := []entity.Facilities{expectedFasilities}
	mockPaging := model.Paging{
		Page:        1,
		RowsPerPage: 1,
		TotalRows:   5,
		TotalPages:  1,
	}
	suite.frm.On("List", 1, 5).Return(mockData, mockPaging, nil)
	actual, paging, err := suite.fuc.FindAllFacilities(1, 5)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Len(suite.T(), actual, 1)
	assert.Equal(suite.T(), mockPaging.Page, paging.Page)
}

// test FindFacilitiesById
func (suite *FacilitiesUseCaseTestSuite) TestFindFacilitiesById_Success() {

	suite.frm.On("GetById", expectedFasilities.ID).Return(expectedFasilities, nil)
	actual, err := suite.fuc.FindFacilitiesById(expectedFasilities.ID)
	assert.Nil(suite.T(), err)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), actual.Name, expectedFasilities.Name)
}

func TestFacilitiesUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(FacilitiesUseCaseTestSuite))
}
