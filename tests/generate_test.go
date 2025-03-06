package password_test

import (
	"regexp"
	"testing"

	"github.com/mechezarreta3/password-generator/pkg/password" // adjust the import path as needed
)

func TestGenerateLength(t *testing.T) {
	length := 20
	pw, err := password.Generate(length, false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(pw) != length {
		t.Errorf("expected length %d, got %d", length, len(pw))
	}
}

func TestGenerateContainsAlpha(t *testing.T) {
	length := 20
	pw, err := password.Generate(length, false, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check for at least one alphabetical character
	alphaRegex := regexp.MustCompile("[a-zA-Z]")
	if !alphaRegex.MatchString(pw) {
		t.Errorf("expected at least one alphabetical character in the password, got %q", pw)
	}
}

func TestGenerateContainsDigit(t *testing.T) {
	length := 20
	pw, err := password.Generate(length, true, false)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check for at least one digit.
	digitRegex := regexp.MustCompile("[0-9]")
	if !digitRegex.MatchString(pw) {
		t.Errorf("expected at least one digit in the password, got %q", pw)
	}
}

func TestGenerateContainsSpecial(t *testing.T) {
	length := 20
	pw, err := password.Generate(length, false, true)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Check for at least one special character.
	// Adjust the regex if your special character set changes.
	specialRegex := regexp.MustCompile(`[!@#$%^&*()\-\_=+\[\]{}|;:,.<>/?]`)
	if !specialRegex.MatchString(pw) {
		t.Errorf("expected at least one special character in the password, got %q", pw)
	}
}

func TestGenerateLengthTooShort(t *testing.T) {
	// When both includeDigits and includeSpecial are true, the minimum acceptable length is 2.
	_, err := password.Generate(1, true, true)
	if err == nil {
		t.Error("expected an error for insufficient length, but got none")
	}
}
