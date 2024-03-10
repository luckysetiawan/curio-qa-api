// Package util provides utility functions to support the server.
package util

import "golang.org/x/crypto/bcrypt"

// HashPassword generate a bcrypt password from the input password.
func HashPassword(password string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return ""
	}

	return string(hashedPassword)
}

// ComparePassword compares the stored hashed password with the input password.
func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}

	return nil
}
