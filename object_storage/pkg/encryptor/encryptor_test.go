package encryptor

import (
	"crypto/aes"
	"encoding/base64"
	"testing"
)

func TestEncryptor(t *testing.T) {
	// Setup
	validKey := "0123456789abcdef0123456789abcdef" // 32 bytes
	invalidKey := "0123456789abcdef"               // 16 bytes (valid AES-128 but wrong)
	encryptor := Encryptor{key: validKey}

	t.Run("basic encryption/decryption", func(t *testing.T) {
		plaintext := "Hello, World!"
		encrypted, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		decrypted, err := encryptor.Decrypt(encrypted)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) != plaintext {
			t.Errorf("Expected %q, got %q", plaintext, decrypted)
		}
	})

	t.Run("empty string", func(t *testing.T) {
		plaintext := ""
		encrypted, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		decrypted, err := encryptor.Decrypt(encrypted)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) != plaintext {
			t.Errorf("Expected empty string, got %q", decrypted)
		}
	})

	t.Run("special characters", func(t *testing.T) {
		plaintext := "!@#$%^&*()_+{}|:\"<>?~`😀"
		encrypted, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		decrypted, err := encryptor.Decrypt(encrypted)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) != plaintext {
			t.Errorf("Expected %q, got %q", plaintext, decrypted)
		}
	})

	t.Run("url-safe encoding", func(t *testing.T) {
		plaintext := "test-url-safe-encoding"
		encrypted, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		// Should only contain URL-safe characters
		for _, c := range encrypted {
			if !(('a' <= c && c <= 'z') ||
				('A' <= c && c <= 'Z') ||
				('0' <= c && c <= '9') ||
				c == '-' || c == '_') {
				t.Errorf("Invalid character in encrypted string: %c", c)
			}
		}
	})

	t.Run("tampered ciphertext", func(t *testing.T) {
		plaintext := "sensitive data"
		encrypted, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		// Tamper with the ciphertext
		tampered := []byte(encrypted)
		if len(tampered) > 10 {
			tampered[10] = 'a' // Introduce error
		}
		_, err = encryptor.Decrypt(string(tampered))
		if err == nil {
			t.Error("Expected error with tampered ciphertext")
		}
	})

	t.Run("wrong key", func(t *testing.T) {
		plaintext := "secret message"
		encrypted, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		// Try decrypting with invalid key
		invalidEncryptor := Encryptor{key: invalidKey}
		_, err = invalidEncryptor.Decrypt(encrypted)
		if err == nil {
			t.Error("Expected error with invalid key")
		}
	})

	t.Run("unique ciphertexts", func(t *testing.T) {
		plaintext := "same input"
		enc1, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("First encryption failed: %v", err)
		}

		enc2, err := encryptor.Encrypt(plaintext)
		if err != nil {
			t.Fatalf("Second encryption failed: %v", err)
		}

		if enc1 == enc2 {
			t.Error("Same plaintext should produce different ciphertexts")
		}
	})

	t.Run("long text", func(t *testing.T) {
		longText := string(make([]byte, 1024*1024)) // 1MB
		encrypted, err := encryptor.Encrypt(longText)
		if err != nil {
			t.Fatalf("Encryption failed: %v", err)
		}

		decrypted, err := encryptor.Decrypt(encrypted)
		if err != nil {
			t.Fatalf("Decryption failed: %v", err)
		}

		if string(decrypted) != longText {
			t.Error("Long text decryption mismatch")
		}
	})

	t.Run("invalid base64", func(t *testing.T) {
		_, err := encryptor.Decrypt("invalid!base64@string")
		if err == nil {
			t.Error("Expected error for invalid base64")
		}
	})

	t.Run("short ciphertext", func(t *testing.T) {
		shortCipher := base64.RawURLEncoding.EncodeToString([]byte("short"))
		_, err := encryptor.Decrypt(shortCipher)
		if err == nil {
			t.Error("Expected error for short ciphertext")
		}
	})

	t.Run("key validation", func(t *testing.T) {
		invalidKeys := []string{
			"16-byte-key",            // Too short
			"24-byte-key-1234567890", // AES-192 (acceptable if supported)
			"",                       // Empty key
		}

		for _, key := range invalidKeys {
			_, err := aes.NewCipher([]byte(key))
			if err == nil {
				t.Errorf("Expected error for invalid key size: %q", key)
			}
		}
	})
}
