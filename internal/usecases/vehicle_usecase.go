package usecases

import (
	"github.com/google/uuid"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/repository"
)

type VehicleUsecase interface {
	CreateVehicle(vehicle *entities.Vehicle) error
	GetAllVehicles() ([]entities.Vehicle, error)
	GetVehicleByID(id uuid.UUID) (entities.Vehicle, error)
	UpdateVehicle(vehicle *entities.Vehicle) error
	DeleteVehicle(id uuid.UUID) error
	AssignDriver(vehicleID, driverID uuid.UUID) error
}

type vehicleUsecase struct {
	vehicleRepo repository.VehicleRepository
	driverRepo  repository.DriverRepository
}

func NewVehicleUsecase(vehicleRepo repository.VehicleRepository, driverRepo repository.DriverRepository) VehicleUsecase {
	return &vehicleUsecase{vehicleRepo: vehicleRepo, driverRepo: driverRepo}
}

func (u *vehicleUsecase) CreateVehicle(vehicle *entities.Vehicle) error {
	vehicle.ID = uuid.New()
	return u.vehicleRepo.Create(vehicle)
}

func (u *vehicleUsecase) GetAllVehicles() ([]entities.Vehicle, error) {
	return u.vehicleRepo.GetAll()
}

func (u *vehicleUsecase) GetVehicleByID(id uuid.UUID) (entities.Vehicle, error) {
	return u.vehicleRepo.GetByID(id)
}

func (u *vehicleUsecase) UpdateVehicle(vehicle *entities.Vehicle) error {
	return u.vehicleRepo.Update(vehicle)
}

func (u *vehicleUsecase) DeleteVehicle(id uuid.UUID) error {
	return u.vehicleRepo.Delete(id)
}

func (u *vehicleUsecase) AssignDriver(vehicleID, driverID uuid.UUID) error {
	return u.vehicleRepo.AssignDriver(vehicleID, driverID)
}
