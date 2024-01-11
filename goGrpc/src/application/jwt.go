package application

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func CreateAllToken(userNo string) (string, string, error) {
	// Access Token 생성
	accClaims := jwt.MapClaims{}
	accClaims["authorized"] = true
	accClaims["userNo"] = userNo
	accClaims["exp"] = time.Now().Add(time.Hour * 24).Unix() // 액세스 토큰 만료 시간을 24시간으로 설정
	accToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accClaims)
	accTokenStr, err := accToken.SignedString([]byte(os.Getenv("ACCESS_KEY")))

	if err != nil {
		return "", "", err
	}

	// Refresh Token 생성
	refClaims := jwt.MapClaims{}
	refClaims["authorized"] = true
	refClaims["userNo"] = userNo
	refClaims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix() // 리프레시 토큰 만료 시간을 1년으로 설정
	refToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refClaims)
	refTokenStr, err := refToken.SignedString([]byte(os.Getenv("REFRESH_KEY")))

	if err != nil {
		return "", "", err
	}

	return accTokenStr, refTokenStr, nil
}

func CreateToken(userNo string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userNo"] = userNo
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("ACCESS_KEY")))
}

func CreateRefreshToken(userNo string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["userNo"] = userNo
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("REFRESH_KEY")))
}

func VerifyToken(tokenString string) (string, error) {
	if tokenString == "" {
		return "", errors.New("Access-Token not provided")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	userNoStr, ok := claims["userNo"].(string)
	if !ok {
		return "", errors.New("userNo is not a string")
	}

	return userNoStr, nil
}

func VerifyRefreshToken(tokenString string) (string, error) {
	fmt.Println("tokenString", tokenString)
	if tokenString == "" {
		return "", errors.New("Refresh-Token not provided")
	}

	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

	claims := jwt.MapClaims{}

	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("REFRESH_KEY")), nil
	})

	if err != nil {
		return "", err
	}

	userNoStr, ok := claims["userNo"].(string)
	if !ok {
		return "", errors.New("userNo is not a string")
	}

	return userNoStr, nil
}
