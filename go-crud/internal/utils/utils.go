// internal/utils/utils.go
package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword hashes a plain text password using bcrypt.
func HashPassword(password string) (string, error) {
	// bcrypt.GenerateFromPassword hashes the password with a cost of bcrypt.DefaultCost
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Println("Error hashing password:", err)
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPasswordHash compares a hashed password with a plain text password
func CheckPasswordHash(password, hashedPassword string) bool {
	// bcrypt.CompareHashAndPassword compares the hashed password with the password provided
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil // returns true if passwords match
}
