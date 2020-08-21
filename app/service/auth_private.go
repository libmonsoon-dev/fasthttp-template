package service

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

const (
	argon2time    = 1
	argon2memory  = 64 * 1024
	argon2threads = 4
	argon2keyLen  = 32
)

// generatePasswordHash is used to generate a new password hash
func (as AuthService) generatePasswordHash(password []byte) string {

	// Generate a Salt
	salt := make([]byte, 16)

	for {
		_, err := rand.Read(salt)
		if err == nil {
			break
		}
		as.logger.Printf("AuthService.generatePasswordHash(): rand.Read(salt): ", err)
	}
	hash := argon2.IDKey(password, salt, argon2time, argon2memory, argon2threads, argon2keyLen)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	format := "$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s"
	return fmt.Sprintf(format, argon2.Version, argon2memory, argon2time, argon2threads, b64Salt, b64Hash)
}

// comparePassword is used to compare a user-inputted password to a hash to see
// if the password matches or not.
func (as AuthService) comparePassword(password []byte, hash string) bool {
	var (
		time    uint32
		memory  uint32
		threads uint8
		keyLen  uint32
	)

	parts := strings.Split(hash, "$")

	if len(parts) != 6 {
		as.logger.Printf("AuthService.comparePassword(): invalid hash len(parts) = %v", len(parts))
		return false
	}

	_, err := fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &time, &threads)
	if err != nil {
		as.logger.Printf("AuthService.comparePassword(): Sscanf(parts[3]): %s", err)
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		as.logger.Printf("AuthService.comparePassword(): base64.RawStdEncoding.DecodeString(parts[4]): %s", err)
		return false
	}

	decodedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		as.logger.Printf("AuthService.comparePassword(): base64.RawStdEncoding.DecodeString(parts[5]): %s", err)
		return false
	}
	keyLen = uint32(len(decodedHash))

	comparisonHash := argon2.IDKey(password, salt, time, memory, threads, keyLen)
	return subtle.ConstantTimeCompare(decodedHash, comparisonHash) == 1
}
