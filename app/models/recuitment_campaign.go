package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)


type RecuitmentCampaign struct {
	orm.Model
	ID uint64 `gorm:"primaryKey"`
	RecruiterID uint64
	DomainID uint64
	Name string 
	Description string
	SalaryFrom int
	SalaryTo int
	Currency string
	Experience string
	HiredEndDate time.Time 
	Urgent bool 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	orm.SoftDeletes

	Recruiter Recruiter `gorm:"foreignKey:RecruiterID, references:UserID"`
	Domain Domain `gorm:"foreignKey:DomainID, references:ID"`
	Applications []Application `gorm:"foreignKey:RecuimentCampaignID, references:ID"`

}