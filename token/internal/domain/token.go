package domain

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base32"
	"errors"
	"time"
)

type Token struct {
	Plaintext string
	Hash      []byte
	UserID    int64
	Expiry    time.Time
}

func GenerateToken(userID int64, ttl time.Duration) (*Token, error) {
	// Create a Token instance containing the user ID, expiry, and scope information.
	token := &Token{
		UserID: userID,
		Expiry: time.Now().Add(ttl),
	}

	// Initialize a byte slice with a length of 16 bytes.
	randomBytes := make([]byte, 16)

	// Fill bytes
	_, err := rand.Read(randomBytes)
	if err != nil {
		return nil, err
	}

	// Encode to base-32-encoded string
	token.Plaintext = base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(randomBytes)

	// Generate a SHA-256 hash of to be stored in the database table
	hash := sha256.Sum256([]byte(token.Plaintext))
	token.Hash = hash[:]

	return token, nil
}

func ValidateTokenPlaintext(tokenPlaintext string) error {
	if tokenPlaintext != "" {
		return errors.New("must be provided")
	}
	if len(tokenPlaintext) != 26 {
		return errors.New("must be 26 bytes long")
	}
	return nil
}

func GetTokenHash(tokenPlaintext string) []byte {
	hash := sha256.Sum256([]byte(tokenPlaintext))
	return hash[:]
}
