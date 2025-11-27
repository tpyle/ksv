package enc

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

var ErrShortCiphertext = errors.New("ciphertext too short")

// Encrypts using AES-256-GCM
// Key does not need to be 32 bytes, it will be hashed to 32 bytes internally
func Encrypt(input io.Reader, key []byte, output io.Writer) error {
	// Hash the key to ensure it's exactly 32 bytes for AES-256
	hash := sha256.Sum256(key)

	// Create AES cipher block
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return err
	}

	// Create GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	// Read all data from input
	plaintext, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)

	// Write nonce followed by ciphertext to output
	if _, err := output.Write(nonce); err != nil {
		return err
	}

	if _, err := output.Write(ciphertext); err != nil {
		return err
	}

	return nil
}

// Decrypts using AES-256-GCM
// Key does not need to be 32 bytes, it will be hashed to 32 bytes internally
func Decrypt(input io.Reader, key []byte, output io.Writer) error {
	// Hash the key to ensure it's exactly 32 bytes for AES-256
	hash := sha256.Sum256(key)

	// Create AES cipher block
	block, err := aes.NewCipher(hash[:])
	if err != nil {
		return err
	}

	// Create GCM cipher mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	// Read all data from input
	data, err := io.ReadAll(input)
	if err != nil {
		return err
	}

	// Check minimum size (nonce + at least some ciphertext)
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		return ErrShortCiphertext
	}

	// Extract nonce and ciphertext
	nonce := data[:nonceSize]
	ciphertext := data[nonceSize:]

	// Decrypt the data
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// Write decrypted data to output
	if _, err := output.Write(plaintext); err != nil {
		return err
	}

	return nil
}
