package usecases

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/mocks"
	"testing"
)

type VehicleUsecaseTestSuite struct {
	suite.Suite
	usecase        VehicleUsecase
	mockRepo       *mocks.VehicleRepository
	mockDriverRepo *mocks.DriverRepository
}

func (suite *VehicleUsecaseTestSuite) SetupTest() {
	suite.mockRepo = new(mocks.VehicleRepository)
	suite.mockDriverRepo = new(mocks.DriverRepository)
	suite.usecase = NewVehicleUsecase(suite.mockRepo, suite.mockDriverRepo)
}

func (suite *VehicleUsecaseTestSuite) TestCreateVehicle() {
	vehicle := &entities.Vehicle{Plate: "XYZ-1234", VehicleModel: "Car"}

	suite.mockRepo.On("Create", vehicle).Return(nil)

	err := suite.usecase.CreateVehicle(vehicle)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *VehicleUsecaseTestSuite) TestGetAllVehicles() {
	vehicles := []entities.Vehicle{
		{ID: uuid.New(), Plate: "XYZ-1234", VehicleModel: "Car"},
		{ID: uuid.New(), Plate: "ABC-5678", VehicleModel: "Truck"},
	}

	suite.mockRepo.On("GetAll").Return(vehicles, nil)

	result, err := suite.usecase.GetAllVehicles()

	suite.NoError(err)
	suite.Equal(vehicles, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *VehicleUsecaseTestSuite) TestGetVehicleByID() {
	id := uuid.New()
	vehicle := entities.Vehicle{ID: id, Plate: "XYZ-1234", VehicleModel: "Car"}

	suite.mockRepo.On("GetByID", id).Return(vehicle, nil)

	result, err := suite.usecase.GetVehicleByID(id)

	suite.NoError(err)
	suite.Equal(vehicle, result)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *VehicleUsecaseTestSuite) TestUpdateVehicle() {
	vehicle := &entities.Vehicle{ID: uuid.New(), Plate: "XYZ-1234", VehicleModel: "Car"}

	suite.mockRepo.On("Update", vehicle).Return(nil)

	err := suite.usecase.UpdateVehicle(vehicle)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *VehicleUsecaseTestSuite) TestDeleteVehicle() {
	id := uuid.New()

	suite.mockRepo.On("Delete", id).Return(nil)

	err := suite.usecase.DeleteVehicle(id)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func (suite *VehicleUsecaseTestSuite) TestAssignDriver() {
	vehicleID := uuid.New()
	driverID := uuid.New()

	suite.mockRepo.On("AssignDriver", vehicleID, driverID).Return(nil)

	err := suite.usecase.AssignDriver(vehicleID, driverID)

	suite.NoError(err)
	suite.mockRepo.AssertExpectations(suite.T())
}

func TestVehicleUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(VehicleUsecaseTestSuite))
}
