package helpers

import (
	"errors"
	"fmt"
	"goravel/app/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/goravel/framework/facades"
	
)

func GetAuthIdentifierName(u models.User) string {
	return "id"
}

func GetAuthIdentifier(u models.User) string {
	return fmt.Sprintf("%d", u.ID)
}

func GetAuthPassword(u models.User) string {
	return u.Password
}

func SetPassword(u *models.User, password string) error {
	hasedPassword , err := facades.Hash().Make(password)
	if err != nil {
		return err
	}
	u.Password = hasedPassword
	return nil
}

func CheckPassword(u models.User, password string) bool {
	return facades.Hash().Check(password, GetAuthPassword(u))
}

func GenerateTokens(u models.User) (string, string, error) {
	jwtSerect := facades.Config().GetString("app.jwt_secret")
	if jwtSerect == "" {
		return "", "", errors.New("jwt secret not set")
	}

	accessTokenExpirationTime := time.Now().Add(time.Minute * time.Duration(facades.Config().GetInt("jwt.tll, 60")))
	accessTokenClaims := jwt.MapClaims{
		"user_id": u.ID,
		"email":  u.Email,
		"type": u.Type,
		"exp":   accessTokenExpirationTime.Unix(),
		"token_type": "access",
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessTokenClaims)
	accessTokenSting, err := accessToken.SignedString([]byte(jwtSerect))
	if err != nil {
		return "", "", err 
	}
	refreshTokenExpirationTime := time.Now().Add(time.Hour *24 * 15)
	refreshTokenClaims := jwt.MapClaims{
		"user_id": u.ID,
		"exp":   refreshTokenExpirationTime.Unix(),
		"token_type": "refresh",
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaims)
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSerect))
	if err != nil {
		return "", "", err 
	}

	return accessTokenSting, refreshTokenString, nil
}

func  VerifyAccessToken(tokenString string) (uint64, error) {
	jwtSecret := facades.Config().GetString("jwt.secret")
	if jwtSecret == "" {
		return 0, errors.New("jwt secret not set")
	}

	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: ")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, err 
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		if claims["token_type"] != "access" {
			return 0, errors.New("invalid token type")
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid user id")
		}

		return uint64(userID), nil
	}

	return 0, errors.New("invalid token")
}

func VerifyRefreshToken(tokenString string) (uint64, error) {
	jwtSecret := facades.Config().GetString("jwt.secret")
	if jwtSecret == "" {
		return 0, errors.New("jwt secret not set")
	}

	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return 0, err 
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		if claims["token_type"] != "refresh" {
			return 0, errors.New("invalid token type")
		}

		userID, ok := claims["user_id"].(float64)
		if !ok {
			return 0, errors.New("invalid user id")
		}

		return uint64(userID), nil
	}
	return 0, errors.New("invalid token")
}

func GetTokenClaims(token string) (jwt.MapClaims, error) {
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(facades.Config().GetString("jwt.secret")), nil
	})

	if err != nil && err.Error() != "token is expired" {
		return nil, err
	}

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok {
		return claims, nil 
	}

	return nil, errors.New("invalid token")
}

// // SetAccessTokenCookie thiết lập access token vào cookie
// func SetAccessTokenCookie(ctx http.Context, token string) {
//     ctx.Cookie().Set("access_token", token, 60*24) // Hết hạn sau 24 giờ
// }

// // SetRefreshTokenCookie thiết lập refresh token vào cookie
// func SetRefreshTokenCookie(ctx http.Context, token string) {
//     ctx.Cookie().Set("refresh_token", token, 60*24*7) // Hết hạn sau 7 ngày
// }

// // ClearTokenCookies xóa cả access và refresh token khỏi cookies
// func ClearTokenCookies(ctx http.Context) {
//     ctx.Cookie().Forget("access_token")
//     ctx.Cookie().Forget("refresh_token")
// }

// // ExtractTokenFromRequest lấy token từ cookie hoặc header Authorization
// func ExtractTokenFromRequest(ctx http.Context) string {
//     // Thử lấy từ cookie trước
//     token := ctx.Cookie().Get("access_token")
//     if token != "" {
//         return token
//     }
    
//     // Nếu không có trong cookie, thử lấy từ header Authorization
//     authHeader := ctx.Request().Header("Authorization", "")
//     if strings.HasPrefix(authHeader, "Bearer ") {
//         return strings.TrimPrefix(authHeader, "Bearer ")
//     }
    
//     return ""
// }

// // ExtractRefreshTokenFromRequest lấy refresh token từ cookie
// func ExtractRefreshTokenFromRequest(ctx http.Context) string {
//     return ctx.Cookie().Get("refresh_token")
// }

// // GenerateAccessToken tạo JWT access token
// func GenerateAccessToken(userID uint) (string, error) {
//     // Lấy secret key từ config
//     secretKey := facades.Config().GetString("jwt.secret")
    
//     // Thiết lập claims
//     claims := jwt.MapClaims{
//         "user_id": userID,
//         "exp":     time.Now().Add(time.Hour * 24).Unix(), // Hết hạn sau 24 giờ
//         "iat":     time.Now().Unix(),
//     }
    
//     // Tạo token
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
//     // Ký token với secret key
//     return token.SignedString([]byte(secretKey))
// }

// // GenerateRefreshToken tạo JWT refresh token
// func GenerateRefreshToken(userID uint) (string, error) {
//     // Lấy secret key từ config
//     secretKey := facades.Config().GetString("jwt.refresh_secret")
    
//     // Thiết lập claims
//     claims := jwt.MapClaims{
//         "user_id": userID,
//         "exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Hết hạn sau 7 ngày
//         "iat":     time.Now().Unix(),
//     }
    
//     // Tạo token
//     token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
//     // Ký token với secret key
//     return token.SignedString([]byte(secretKey))
// }

// // ParseToken phân tích JWT token và trả về claims
// func ParseToken(tokenString string) (*TokenClaims, error) {
//     // Lấy secret key từ config
//     secretKey := facades.Config().GetString("jwt.secret")
    
//     // Parse token
//     token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
//         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
//             return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
//         }
//         return []byte(secretKey), nil
//     })
    
//     if err != nil {
//         return nil, err
//     }
    
//     if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
//         userID, ok := claims["user_id"].(float64)
//         if !ok {
//             return nil, errors.New("invalid user ID in token")
//         }
        
//         return &TokenClaims{
//             UserID: uint(userID),
//         }, nil
//     }
    
//     return nil, errors.New("invalid token")
// }

// // TokenClaims chứa thông tin từ JWT token
// type TokenClaims struct {
//     UserID uint
// }