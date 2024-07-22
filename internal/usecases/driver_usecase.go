package usecases

import (
	"github.com/google/uuid"
	"go-driver-register/internal/domain/entities"
	"go-driver-register/internal/repository"
)

type DriverUsecase interface {
	CreateDriver(driver *entities.Driver) error
	GetAllDrivers() ([]entities.Driver, error)
	GetDriverByID(id uuid.UUID) (entities.Driver, error)
	UpdateDriver(driver *entities.Driver) error
	DeleteDriver(id uuid.UUID) error
}

type driverUsecase struct {
	driverRepo repository.DriverRepository
}

func NewDriverUsecase(driverRepo repository.DriverRepository) DriverUsecase {
	return &driverUsecase{driverRepo: driverRepo}
}

func (u *driverUsecase) CreateDriver(driver *entities.Driver) error {
	driver.ID = uuid.New()
	return u.driverRepo.Create(driver)
}

func (u *driverUsecase) GetAllDrivers() ([]entities.Driver, error) {
	return u.driverRepo.GetAll()
}

func (u *driverUsecase) GetDriverByID(id uuid.UUID) (entities.Driver, error) {
	return u.driverRepo.GetByID(id)
}

func (u *driverUsecase) UpdateDriver(driver *entities.Driver) error {
	return u.driverRepo.Update(driver)
}

func (u *driverUsecase) DeleteDriver(id uuid.UUID) error {
	return u.driverRepo.Delete(id)
}
