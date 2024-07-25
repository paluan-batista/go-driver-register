package usecases

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/mocks"
	"testing"
)

type DriverUsecaseTestSuite struct {
	suite.Suite
	usecase  DriverUsecase
	mockRepo *mocks.DriverRepository
}

func (suite *DriverUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.DriverRepository)
	suite.usecase = NewDriverUsecase(suite.mockRepo)
}

func (suite *DriverUsecaseTestSuite) TestCreateDriver() {
	driver := &entities.Driver{Name: "John Doe", LicenseID: "12345"}

	suite.mockRepo.On("Create", driver).Return(nil)

	err := suite.usecase.CreateDriver(driver)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *DriverUsecaseTestSuite) TestGetAllDrivers() {
	drivers := []entities.Driver{
		{ID: uuid.New(), Name: "John Doe", LicenseID: "12345"},
		{ID: uuid.New(), Name: "Jane Smith", LicenseID: "67890"},
	}

	suite.mockRepo.On("GetAll").Return(drivers, nil)

	result, err := suite.usecase.GetAllDrivers()

	suite.NoError(err)
	suite.Equal(drivers, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *DriverUsecaseTestSuite) TestGetDriverByID() {
	id := uuid.New()
	driver := entities.Driver{ID: id, Name: "John Doe", LicenseID: "12345"}

	suite.mockRepo.On("GetByID", id).Return(driver, nil)

	result, err := suite.usecase.GetDriverByID(id)

	suite.NoError(err)
	suite.Equal(driver, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *DriverUsecaseTestSuite) TestUpdateDriver() {
	driver := &entities.Driver{ID: uuid.New(), Name: "John Doe", LicenseID: "12345"}

	suite.mockRepo.On("Update", driver).Return(nil)

	err := suite.usecase.UpdateDriver(driver)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *DriverUsecaseTestSuite) TestDeleteDriver() {
	id := uuid.New()

	suite.mockRepo.On("Delete", id).Return(nil)

	err := suite.usecase.DeleteDriver(id)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestDriverUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(DriverUsecaseTestSuite))
}
