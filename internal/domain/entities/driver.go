package entities

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Driver struct {
	gorm.Model
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:uuid_generate_v4()"`
	Name      string    `json:"name"`
	LicenseID string    `json:"license_id"`
}
