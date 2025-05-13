package services

import (
	"goravel/app/models"
)

type UserService interface {
	// Get user by ID
	Get(id uint64) (*models.User, error)

	// Find user by email
	FindByEmail(email string) (*models.User, error)

	// Register a new user
	Register(name, email, password, phone, userType string) (*models.User, error)

	// Update user profile
	Update(id uint64, name, email, phone, avatar string) (*models.User, error)

	// Change user password
	ChangePassword(id uint64, oldPassword, newPassword string) error

	// Delete a user
	Delete(id uint64) error

	// Get users with pagination
	Paginate(page, limit int) ([]models.User, int64, error)
}
