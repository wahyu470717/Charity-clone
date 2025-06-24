package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTUtil struct {
	secret []byte
}

var JWTUtilInstance *JWTUtil

func NewJWTUtil(secret string) *JWTUtil {
	return &JWTUtil{secret: []byte(secret)}
}

func InitJWTUtil(secret string) {
	JWTUtilInstance = NewJWTUtil(secret)
}
func GetJWTUtil() *JWTUtil {
	return JWTUtilInstance
}

type JWTUtilInterface interface {
	GenerateJWT(username string, id int64, role string) (string, error)
}

func (j *JWTUtil) GenerateJWT(username string, id int64, role string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"id":       id,
		"role":     role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // Ubah dari ES256 ke HS256
	return token.SignedString(j.secret)
}

func (j *JWTUtil) VerifyJWT(tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unpexted signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})
}

func (j *JWTUtil) parseClaimsFromHeader(authHeader string) (jwt.MapClaims, error) {
	parts := strings.SplitN(authHeader, " ", 2)
	if len(parts) != 2 || parts[0] != "Bearer" {
		return nil, errors.New("unexpected signing method")
	}

	tokenString := parts[1]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return j.secret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token tidak valid atau klaim tidak bisa dibaca")
	}

	return claims, nil
}

func (j *JWTUtil) VerifyToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")

}



// Generate reset password token (shorter expiry)
func (j *JWTUtil) GenerateResetToken(userID int64, email string) (string, error) {
	claims := jwt.MapClaims{
		"id":    userID,
		"email": email,
		"type":  "reset_password",
		"exp":   time.Now().Add(1 * time.Hour).Unix(), // 1 hour expiry
		"iat":   time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}