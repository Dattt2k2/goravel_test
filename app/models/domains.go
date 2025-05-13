package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Domain struct {
	orm.Model
	ID uint64 `gorm:"primaryKey"`
	Name string
	Description string
	CreatedAt time.Time 
	UpdatedAt time.Time
	DeletedAt *time.Time
	orm.SoftDeletes

	RecuimentCampaign []RecuitmentCampaign `gorm:"foreignKey:ID"`
}