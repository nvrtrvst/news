package conv

import (
	"strconv"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

// Atoi converts a string to an integer.
// It returns the integer value and an error if the conversion fails.
func Atoi(s string) (int, error) {
	return strconv.Atoi(s)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateSlug(title string) string {
	slug := strings.ToLower(title)
	slug = strings.ReplaceAll(slug, " ", "-")
	return slug
}


