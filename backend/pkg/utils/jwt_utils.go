package utils

import (
	"fmt"
	"time"

	"github.com/blastertwist/flag-dash/internal/dao"
	"github.com/golang-jwt/jwt/v4"
)


type Claims struct {
	Email	string
	ID		string
	jwt.StandardClaims
}

type ProjectClaims struct {
	ID	string
	jwt.StandardClaims
}


func GenerateJWT(user *dao.User, secretKey string, expiredTime uint8) (string, error){
	
	expirationTime := time.Now().Add(time.Duration(expiredTime) * time.Minute)
	claims := &Claims{
		Email: user.Email,
		ID: fmt.Sprint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func GenerateJWTProject(project *dao.Project, secretKey string) (string, error) {
	claims := &ProjectClaims{
		ID: fmt.Sprint(project.ID),
		StandardClaims: jwt.StandardClaims{},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string, tokenSecret string) (bool, *Claims, error){
	
	claims := &Claims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {

        return []byte(tokenSecret), nil
    })

	if err != nil {
		return false, nil, err
	}

	if !token.Valid {
		return false, nil, nil
	}

	return true, claims, nil
}

func VerifyJWTProject(tokenString string, tokenSecret string) (bool, *ProjectClaims, error) {
	claims := &ProjectClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error){
		return []byte(tokenSecret), nil
	})

	if err != nil {
		return false, nil, err
	}

	if !token.Valid {
		return false, nil, nil
	}

	return true, claims, nil
}