package repository

import (
	"github.com/google/uuid"
	"go-driver-register/internal/domain/entities"
	"gorm.io/gorm"
)

type DriverRepository interface {
	Create(driver *entities.Driver) error
	GetAll() ([]entities.Driver, error)
	GetByID(id uuid.UUID) (entities.Driver, error)
	Update(driver *entities.Driver) error
	Delete(id uuid.UUID) error
}

type driverRepository struct {
	db *gorm.DB
}

func NewDriverRepository(db *gorm.DB) DriverRepository {
	return &driverRepository{db: db}
}

func (r *driverRepository) Create(driver *entities.Driver) error {
	return r.db.Create(driver).Error
}

func (r *driverRepository) GetAll() ([]entities.Driver, error) {
	var drivers []entities.Driver
	err := r.db.Find(&drivers).Error
	return drivers, err
}

func (r *driverRepository) GetByID(id uuid.UUID) (entities.Driver, error) {
	var driver entities.Driver
	err := r.db.First(&driver, "id = ?", id).Error
	return driver, err
}

func (r *driverRepository) Update(driver *entities.Driver) error {
	return r.db.Save(driver).Error
}

func (r *driverRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Driver{}, "id = ?", id).Error
}
