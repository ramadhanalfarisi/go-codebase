package helpers

import (
	"bytes"
	"crypto/rand"
	"crypto/sha512"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)
type ValidatePassword struct {
	PasswordHashed string
	PasswordInput string
}

// HashPassword creates an Argon2id hash with salt and parameters encoded.
func HashPassword(password string) (string, error) {
	// Parameters — you can tune these
	var (
		memory      uint32 = 64 * 1024 // 64 MB
		iterations  uint32 = 3
		parallelism uint8  = 2
		saltLength  uint32 = 16
		keyLength   uint32 = 32
	)

	// Generate random salt
	salt := make([]byte, saltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	// Derive the key
	hash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, keyLength)

	// Encode as a single string for storage
	encoded := fmt.Sprintf("$argon2id$v=19$m=%d,t=%d,p=%d$%s$%s",
		memory, iterations, parallelism,
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash))

	return encoded, nil
}

// VerifyPassword compares a plaintext password with an Argon2id hash.
func VerifyPassword(dataPass ValidatePassword) (bool, error) {
	parts := strings.Split(dataPass.PasswordHashed, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	var memory uint32
	var iterations uint32
	var parallelism uint8
	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, err
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	hash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	keyLen := uint32(len(hash))
	computed := argon2.IDKey([]byte(dataPass.PasswordInput), salt, iterations, memory, parallelism, keyLen)

	// Constant-time compare
	if bytes.Equal(hash, computed) {
		return true, nil
	}
	return false, nil
}

func SHA512(text string) string {
	hasher := sha512.New()

	// Write the input data to the hash object.
	// The Write method expects a byte slice, so convert the string.
	hasher.Write([]byte(text))

	// Calculate the final hash sum.
	// The Sum method takes an optional byte slice to append the hash to.
	// Passing nil returns a new byte slice containing the hash.
	hashedBytes := hasher.Sum(nil)

	// Convert the byte slice hash to a hexadecimal string for display.
	hexHash := hex.EncodeToString(hashedBytes)

	return hexHash
}