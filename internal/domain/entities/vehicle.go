package entities

import "github.com/google/uuid"

type Vehicle struct {
	ID       uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Plate    string    `json:"plate"`
	Model    string    `json:"model"`
	Driver   *Driver   `json:"driver,omitempty" gorm:"foreignKey:DriverID"`
	DriverID uuid.UUID `json:"driver_id,omitempty"`
}
