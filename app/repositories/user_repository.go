package repositories

import (
	"goravel/app/models"
)

type UserRepository interface {
	// Find a user by ID
	Find(id uint64) (*models.User, error)

	// Find a user by email
	FindByEmail(email string) (*models.User, error)

	// Create a new user
	Create(user *models.User) error

	// Update an existing user
	Update(user *models.User) error

	// Delete a user by ID
	Delete(id uint64) error

	// List all users with pagination
	Paginate(page, limit int) ([]models.User, int64, error)
}
