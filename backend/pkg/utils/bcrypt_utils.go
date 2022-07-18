package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error){
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	
	if err != nil {
		return "", err
	}
	hashedPassword := string(hashBytes)

	return hashedPassword, nil

}

func ValidatePassword(password string, hashedPassword string) (bool, error) {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false, err
	}
	
	return true, nil
}