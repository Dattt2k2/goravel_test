package controllers

import (
	"fmt"
	"goravel/app/models"
	"goravel/app/services"
	"time"

	// "net/http"
	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

type AuthController struct {
	authService services.AuthService
}

func NewAuthController(authservice services.AuthService) *AuthController {
	return &AuthController{
		authService: authservice,
	}
}

func (ctrl *AuthController) Register(ctx http.Context) http.Response {

	validator, err := facades.Validation().Make(ctx.Request().All(), map[string]string{
		"email":    "required|email",
		"password": "required|min:6",
		"type":   	"required|in:admin,user",
	})

	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Validation failed",
		})
	}

	if validator.Fails() {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": validator.Errors(),
		})
	}

	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	user, accessToken, refreshToken, err := ctrl.authService.Register(email, password)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Registration failed",
		})
	}

	return ctx.Response().Json(200, map[string]string{
		"message":       "Registration successful",
		"user_id":       fmt.Sprintf("%d", user.ID),
		"email":         user.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
	// ctx.Response().Json(http.StatusOK, http.Json{
	// 	"message":       "Registration successful",
	// 	"user":          user,
	// 	"access_token":  accessToken,
	// 	"refresh_token": refreshToken,
	// })
}

func (ctrl *AuthController) Login(ctx http.Context) http.Response {
	validator, err := facades.Validation().Make(ctx.Request().All(), map[string]string{
		"email":    "required|email",
		"password": "required|min:6",
	})

	if err != nil {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": "Validation failed",
		})

	}
	if validator.Fails() {
		return ctx.Response().Json(http.StatusBadRequest, http.Json{
			"error": validator.Errors(),
		})

	}

	email := ctx.Request().Input("email")
	password := ctx.Request().Input("password")
	user, accessToken, refreshToken, err := ctrl.authService.SignIn(email, password)
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Login failed",
		})

	}

	// Thiết lập cookies với token
	accessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    accessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	refreshTokenCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	ctx.Response().Cookie(accessTokenCookie)
	ctx.Response().Cookie(refreshTokenCookie)

	// ctx.Response().Json(http.StatusOK, http.Json{
	// 	"message":       "Login successful",
	// 	"user":          user,
	// 	"access_token":  accessToken,
	// 	"refresh_token": refreshToken,
	// })
	return ctx.Response().Json(200, map[string]string{
		"message":       "Login successful",
		"user_id":       fmt.Sprintf("%d", user.ID),
		"email":         user.Email,
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})

}

func (ctrl *AuthController) Me(ctx http.Context) http.Response {
	user, ok := ctx.Value("user").(*models.User)
	if !ok {
		// ctx.Response().Json(http.StatusUnauthorized, http.Json{
		// 	"error": "Unauthorized",
		// })
		return ctx.Response().Json(401, map[string]string{
			"error": "Unauthorized",
		})
	}
	// ctx.Response().Json(http.StatusOK, http.Json{
	// 	"message": "User information",
	// 	"user":    user,
	// })
	return ctx.Response().Json(200, map[string]string{
		"message": "User information",
		"user_id": fmt.Sprintf("%d", user.ID),
		"email":   user.Email,
		"name":    user.Name,
	})
}

// RefreshToken xử lý việc làm mới token dựa trên refresh token
func (ctrl *AuthController) RefreshToken(ctx http.Context) http.Response {
	// Lấy refresh token từ cookie hoặc request
	refreshToken := ctx.Request().Cookie("refresh_token", "")
	if refreshToken == "" {
		refreshToken = ctx.Request().Input("refresh_token")
	}

	if refreshToken == "" {
		// ctx.Response().Json(http.StatusBadRequest, http.Json{
		// 	"error": "Refresh token is required",
		// })
		return ctx.Response().Json(400, map[string]string{
			"error": "Refresh token is required",
		})

	}

	// Gọi service để làm mới token
	user, newAccessToken, newRefreshToken, err := ctrl.authService.RefreshToken(refreshToken)
	if err != nil {
		// ctx.Response().Json(http.StatusUnauthorized, http.Json{
		// 	"error":   "Invalid refresh token",
		// 	"details": err.Error(),
		// })

		return ctx.Response().Json(401, map[string]string{
			"error":   "Invalid refresh token",
			"details": err.Error(),
		})

	}

	// Thiết lập cookies với token mới
	accessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    newAccessToken,
		Expires:  time.Now().Add(24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	refreshTokenCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    newRefreshToken,
		Expires:  time.Now().Add(7 * 24 * time.Hour),
		Path:     "/",
		HttpOnly: true,
	}
	ctx.Response().Cookie(accessTokenCookie)
	ctx.Response().Cookie(refreshTokenCookie)

	// Trả về token mới
	// ctx.Response().Json(http.StatusOK, http.Json{
	// 	"message":       "Token refreshed successfully",
	// 	"user":          user,
	// 	"access_token":  newAccessToken,
	// 	"refresh_token": newRefreshToken,
	// })

	return ctx.Response().Json(200, map[string]string{
		"message":       "Token refreshed successfully",
		"user_id":       fmt.Sprintf("%d", user.ID),
		"email":         user.Email,
		"access_token":  newAccessToken,
		"refresh_token": newRefreshToken,
	})
}

// Logout xóa các token khỏi cookie và đăng xuất người dùng
func (ctrl *AuthController) Logout(ctx http.Context) http.Response {
	// Gọi đến service để xử lý đăng xuất (nếu cần)
	err := ctrl.authService.SignOut()
	if err != nil {
		return ctx.Response().Json(http.StatusInternalServerError, http.Json{
			"error": "Logout failed",
		})
	}

	// Tạo cookie hết hạn để xóa token
	accessTokenCookie := http.Cookie{
		Name:     "access_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Đặt thời gian hết hạn trong quá khứ
		Path:     "/",
		HttpOnly: true,
	}
	refreshTokenCookie := http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour), // Đặt thời gian hết hạn trong quá khứ
		Path:     "/",
		HttpOnly: true,
	}

	// Thiết lập cookie để xóa
	ctx.Response().Cookie(accessTokenCookie)
	ctx.Response().Cookie(refreshTokenCookie)

	return ctx.Response().Json(http.StatusOK, http.Json{
		"message": "Logout successful",
	})
}
