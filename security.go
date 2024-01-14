package main

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return false
	} else {
		return true
	}
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err == nil {
		return string(hashedPassword), nil
	} else {
		return "", err
	}

}
