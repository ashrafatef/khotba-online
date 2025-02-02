package models

import (
	"time"
)

type MasjedStatus string

const (
	MasjedStatusActive   MasjedStatus = "active"
	MasjedStatusInactive MasjedStatus = "inactive"
)

type Masjed struct {
	ID        int          `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string       `json:"name"`
	Languages []string     `gorm:"type:text[]" json:"languages"`
	EmamId    int          `json:"emam_id"`
	Status    MasjedStatus `json:"status"`
	Location  interface{}  `gorm:"type:jsonb" json:"location"`
	UpdatedAt time.Time    `json:"updated_at"`
	CreatedAt time.Time    `json:"created_at"`
}
