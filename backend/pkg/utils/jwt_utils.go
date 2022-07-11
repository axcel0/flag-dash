package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)


type Claims struct {
	Data interface {}
	jwt.StandardClaims
}


func GenerateJWT(data interface{}, secretKey string, expiredTime uint8) (string, error){
	
	expirationTime := time.Now().Add(time.Duration(expiredTime) * time.Minute)
	claims := &Claims{
		Data: data,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (bool, error){
	
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(jwtKey *jwt.Token)(interface{}, error){
		return jwtKey, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}