package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type WorkingHistory struct {
	orm.Model
	ID uint64 `gorm:"primaryKey"`
	CandidateID uint64
	CompanyID uint64
	ApplicationID uint64
	StartDate string
	EndDate string
	Status string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	orm.SoftDeletes

	Candidate Candidate `gorm:"foreignKey:CandidateID, references:UserID"`
	Company Company `gorm:"foreignKey:CompanyID, references:CompanyID"`
	Application Application `gorm:"foreignKey:ApplicationID, references:ID"`
}