package utils

import (
	"regexp"
	"strings"
)

func NormalizeEmail(email string) string {
	return strings.ToLower(strings.TrimSpace(email))
}

func IsValidEmail(email string) bool {
	regex := `^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`
	match, _ := regexp.MatchString(regex, email)
	return match
}

func IsValidPhone(phone string) bool {
	// Assumes E.164 or local formats
	return strings.HasPrefix(phone, "+") || strings.HasPrefix(phone, "07")
}
