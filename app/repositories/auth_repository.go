package repositories

import (
	"goravel/app/http/helpers"
	"goravel/app/models"

	// "goravel/app/helpers"
	"errors"

	"github.com/goravel/framework/facades"
	// "github.com/goravel/framework/http"
)

type AuthRepository interface {
	Register(email, password string) (*models.User, error)
	// FindUserByEmail(email string) (uint64, error)
	SignIn(email, password string) (*models.User, error)
	SignOut() error
	RefreshToken(refreshToken string) (string, string, *models.User, error)
	// StoreToken(userID uint64, token string) error
	// VerifyToken(token string) (uint64, error)
	FindUserById(id uint64) (*models.User, error)
}

type AuthRepositoryImpl struct{}

func NewAuthRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (r *AuthRepositoryImpl) Register(email, password string) (*models.User, error) {
	var count int64
	if err := facades.Orm().Query().Model(&models.User{}).Where("email", email).Count(&count).Error; err != nil {
		return nil, errors.New("failed to check email existence")
	}
	if count > 0 {
		return nil, errors.New("email already exists")
	}

	user := &models.User{
		Email:    email,
		Password: password,
	}
	if err := helpers.SetPassword(user, password); err != nil {
		return nil, err
	}
	if err := facades.Orm().Query().Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *AuthRepositoryImpl) SignIn(email, password string) (*models.User, error) {
	var user models.User
	if err := facades.Orm().Query().Where("email", email).First(&user); err != nil {
		return nil, errors.New("invalid credentials")
	}

	if !helpers.CheckPassword(user, password) {
		return nil, errors.New("invalid credentials")
	}

	return &user, nil
}

// RefreshToken tạo token mới từ refresh token
func (r *AuthRepositoryImpl) RefreshToken(refreshToken string) (string, string, *models.User, error) {
	// Xác thực refresh token
	userID, err := helpers.VerifyRefreshToken(refreshToken)
	if err != nil {
		return "", "", nil, err
	}

	// Lấy thông tin user từ database
	var user models.User
	if err := facades.Orm().Query().Where("id", userID).First(&user); err != nil {
		return "", "", nil, errors.New("user not found")
	}

	// Tạo cặp token mới
	accessToken, newRefreshToken, err := helpers.GenerateTokens(user)
	if err != nil {
		return "", "", nil, err
	}

	return accessToken, newRefreshToken, &user, nil
}

func (r *AuthRepositoryImpl) FindUserById(id uint64) (*models.User, error) {
	var user models.User
	if err := facades.Orm().Query().Where("id", id).First(&user); err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// SignOut xử lý đăng xuất người dùng
func (r *AuthRepositoryImpl) SignOut() error {
	// Trong repository, chúng ta có thể thêm các logic liên quan đến việc xóa token
	// trong cơ sở dữ liệu hoặc danh sách token bị thu hồi nếu cần.
	// Với JWT, thông thường chúng ta không cần làm gì ở mức repository
	// vì chỉ cần xóa token ở client (thông qua xóa cookie)
	return nil
}
