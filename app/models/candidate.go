package models

import (
	"time"
)

type Candidate struct {
	UserID         uint64 `gorm:"primaryKey"` // UserID là khóa chính
	DOb            time.Time
	YearEnrollment time.Time
	YearGraduation time.Time
	IsGraduated    bool
	University     string
	Experience     string // Sửa lỗi chính tả
	CreatedAt      time.Time
	UpdatedAt      time.Time
	DeletedAt      *time.Time

	// Quan hệ
	User             User             `gorm:"foreignKey:UserID;references:ID"`          // Sửa dấu phẩy thành chấm phẩy
	WorkingHistories []WorkingHistory `gorm:"foreignKey:CandidateID;references:UserID"` // Sửa tên trường và cú pháp
	Applications     []Application    `gorm:"foreignKey:CandidateID;references:UserID"` // Sửa cú pháp
}
