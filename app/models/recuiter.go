package models

import (
	"time"
)

type Recruiter struct {
	UserID    uint64 `gorm:"primaryKey"` // UserID là khóa chính
	CompanyID uint64
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time

	// Quan hệ
	User      User                 `gorm:"foreignKey:UserID;references:ID"`          // Sửa định nghĩa quan hệ
	Company   Company              `gorm:"foreignKey:CompanyID;references:ID"`       // Sửa định nghĩa quan hệ
	Campaigns []RecuitmentCampaign `gorm:"foreignKey:RecruiterID;references:UserID"` // Sửa định nghĩa quan hệ
}
