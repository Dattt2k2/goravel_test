package models

import (
	"time"

	"github.com/goravel/framework/database/orm"
)

type Application struct {
	orm.Model
	ID          uint64 `gorm:"primaryKey"`
	CandidateID uint64
	RecuimentCampaignID uint64
	CV_url string
	Description string
	Status string
	CreatedAt time.Time 
	UpdatedAt time.Time
	DeletedAt *time.Time
	orm.SoftDeletes

	Candidate Candidate `gorm:"foreignKey:CandidateID, references:UserID"`
	RecruitmentCampaign []RecuitmentCampaign `gorm:"foreignKey:RecuimentCampaignID, references:UserID"`
}