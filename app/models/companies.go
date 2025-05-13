package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Company struct {
	orm.Model
	CompanyID uint64 `gorm:"primaryKey"`
	Name string 
	Address string
	Phone string
	Email string
	Website string
	Logo string
	Domain string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	orm.SoftDeletes

	Recruiters []Recruiter `gorm:"foreignKey:CompanyID"`
	WorkingHistory []WorkingHistory `gorm:"foreignKey:CompanyID"`
	Campaigns []RecuitmentCampaign `gorm:"foreignKey:CompanyID"`
}