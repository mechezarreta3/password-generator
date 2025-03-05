package password

import (
	"crypto/rand"
	"errors"
	"math/big"
)

// GeneratePassword creates a random password of the specified length.
// It always uses letters (both uppercase and lowercase).
// If includeDigits is true, it forces at least one digit to be present.
// If includeSpecial is true, it forces at least one special character to be present.
func GeneratePassword(length int, includeDigits bool, includeSpecial bool) (string, error) {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits := "0123456789"
	specialChars := "!@#$%^&*()-_=+[]{}|;:,.<>/?"

	// Build the full charset for password generation according to the flags.
	charset := letters
	if includeDigits {
		charset += digits
	}
	if includeSpecial {
		charset += specialChars
	}

	// Prepare a slice to hold the forced characters.
	minRequiredChars := []byte{}

	// Require one digit if flag is 'true'.
	if includeDigits {
		i, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", err
		}
		minRequiredChars = append(minRequiredChars, digits[i.Int64()])
	}

	// Require one special character if flag is 'true'.
	if includeSpecial {
		i, err := rand.Int(rand.Reader, big.NewInt(int64(len(specialChars))))
		if err != nil {
			return "", err
		}
		minRequiredChars = append(minRequiredChars, specialChars[i.Int64()])
	}

	// Ensure the total length can accommodate the forced characters.
	if length <= len(minRequiredChars) {
		return "", errors.New("password length is too short for the required options")
	}

	// Calculate the number of remaining characters to generate.
	remaining := length - len(minRequiredChars)
	passwordChars := make([]byte, 0, length)

	// Generate random characters for the remaining length from the full charset.
	for i := 0; i < remaining; i++ {
		i, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		passwordChars = append(passwordChars, charset[i.Int64()])
	}

	// Append the minimally required characters.
	passwordChars = append(passwordChars, minRequiredChars...)

	// Shuffle the resulting slice using the Fisher-Yates algorithm.
	for i := len(passwordChars) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return "", err
		}
		passwordChars[i], passwordChars[j.Int64()] = passwordChars[j.Int64()], passwordChars[i]
	}

	return string(passwordChars), nil
}
