package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const secretKey = "supersecret"

func GenerateToken(email string, userId int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parseToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return 0, errors.New("could not parse token")
	}
	if !parseToken.Valid {
		return 0, errors.New("invalid token")
	}
	claims, ok := parseToken.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}
	userId, ok := claims["userId"].(float64)
	if !ok {
		return 0, errors.New("userId is not a valid number")
	}
	return int64(userId), nil
}
