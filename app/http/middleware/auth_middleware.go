package middleware

import (
	"goravel/app/http/helpers"
	"goravel/app/models"
	"goravel/app/repositories"
	"goravel/app/services"
	"strings"
	"time"

	"github.com/goravel/framework/contracts/http"
	"github.com/goravel/framework/facades"
)

func Auth() http.Middleware {
	return func(ctx http.Context) {
		// Lấy access token từ header hoặc cookie
		token := ctx.Request().Header("Authorization", "")
		if token != "" && len(token) > 7 {
			token = token[7:]
		} else {
			token = ctx.Request().Cookie("access_token", "")
		}

		if token == "" {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": "Unauthorized: Access token required",
			})
			return
		}

		// Kiểm tra access token
		userID, err := helpers.VerifyAccessToken(token)
		if err != nil {
			// Access token không hợp lệ hoặc đã hết hạn
			// Kiểm tra lỗi có phải do token hết hạn không
			tokenExpiredError := "token is expired"
			if !strings.Contains(strings.ToLower(err.Error()), tokenExpiredError) {
				// Nếu không phải lỗi token hết hạn, trả về lỗi unauthorized
				ctx.Response().Json(http.StatusUnauthorized, http.Json{
					"error": "Unauthorized: Invalid access token",
				})
				return
			}

			// Nếu token hết hạn, thử làm mới token bằng refresh token
			refreshToken := ctx.Request().Cookie("refresh_token", "")
			if refreshToken == "" {
				ctx.Response().Json(http.StatusUnauthorized, http.Json{
					"error": "Unauthorized: Refresh token not found",
				})
				return
			}

			// Tạo auth service để làm mới token
			authService := services.NewAuthService(repositories.NewAuthRepository())
			user, newAccessToken, newRefreshToken, err := authService.RefreshToken(refreshToken)
			if err != nil {
				ctx.Response().Json(http.StatusUnauthorized, http.Json{
					"error":   "Unauthorized: Unable to refresh token",
					"details": err.Error(),
				})
				return
			}

			// Cập nhật cookies với token mới - sử dụng đúng cú pháp của Goravel
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

			// Thiết lập user vào context và tiếp tục xử lý request
			ctx.WithValue("user", *user)
			ctx.Request().Next()
			return
		}

		// Nếu access token hợp lệ, tiếp tục với việc lấy thông tin user
		var user models.User
		if err := facades.Orm().Query().Where("id", userID).First(&user); err != nil {
			ctx.Response().Json(http.StatusUnauthorized, http.Json{
				"error": "Unauthorized: User not found",
			})
			return
		}

		ctx.WithValue("user", user)
		ctx.Request().Next()
	}
}
