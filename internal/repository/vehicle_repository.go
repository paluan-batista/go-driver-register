package repository

import (
	"github.com/google/uuid"
	"go-driver-register/internal/domain/entities"
	"gorm.io/gorm"
)

type VehicleRepository interface {
	Create(vehicle *entities.Vehicle) error
	GetAll() ([]entities.Vehicle, error)
	GetByID(id uuid.UUID) (entities.Vehicle, error)
	Update(vehicle *entities.Vehicle) error
	Delete(id uuid.UUID) error
	AssignDriver(vehicleID, driverID uuid.UUID) error
}

type vehicleRepository struct {
	db *gorm.DB
}

func NewVehicleRepository(db *gorm.DB) VehicleRepository {
	return &vehicleRepository{db: db}
}

func (r *vehicleRepository) Create(vehicle *entities.Vehicle) error {
	return r.db.Create(vehicle).Error
}

func (r *vehicleRepository) GetAll() ([]entities.Vehicle, error) {
	var vehicles []entities.Vehicle
	err := r.db.Find(&vehicles).Error
	return vehicles, err
}

func (r *vehicleRepository) GetByID(id uuid.UUID) (entities.Vehicle, error) {
	var vehicle entities.Vehicle
	err := r.db.First(&vehicle, "id = ?", id).Error
	return vehicle, err
}

func (r *vehicleRepository) Update(vehicle *entities.Vehicle) error {
	return r.db.Save(vehicle).Error
}

func (r *vehicleRepository) Delete(id uuid.UUID) error {
	return r.db.Delete(&entities.Vehicle{}, "id = ?", id).Error
}

func (r *vehicleRepository) AssignDriver(vehicleID, driverID uuid.UUID) error {
	return r.db.Model(&entities.Vehicle{}).Where("id = ?", vehicleID).Update("driver_id", driverID).Error
}
