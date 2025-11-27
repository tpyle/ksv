package enc

import (
	"bytes"
	"testing"
)

func TestLongData(t *testing.T) {
	// Prepare long data
	plainData := bytes.Repeat([]byte("A"), 10*1024*1024) // 10 MB of 'A's
	key := []byte("longdatatestkey")

	// Encrypt
	var encrypted bytes.Buffer
	err := Encrypt(bytes.NewReader(plainData), key, &encrypted)
	if err != nil {
		t.Fatalf("Encryption failed: %v", err)
	}

	// Decrypt
	var decrypted bytes.Buffer
	err = Decrypt(&encrypted, key, &decrypted)
	if err != nil {
		t.Fatalf("Decryption failed: %v", err)
	}

	// Verify
	if !bytes.Equal(plainData, decrypted.Bytes()) {
		t.Fatal("Decrypted data does not match original long data")
	}
}

func TestTwoEncryptionsDifferent(t *testing.T) {
	plainData := []byte("Sample data for encryption")
	key := []byte("testkey123")

	// First encryption
	var encrypted1 bytes.Buffer
	err := Encrypt(bytes.NewReader(plainData), key, &encrypted1)
	if err != nil {
		t.Fatalf("First encryption failed: %v", err)
	}

	// Second encryption
	var encrypted2 bytes.Buffer
	err = Encrypt(bytes.NewReader(plainData), key, &encrypted2)
	if err != nil {
		t.Fatalf("Second encryption failed: %v", err)
	}

	// Verify that the two ciphertexts are different
	if bytes.Equal(encrypted1.Bytes(), encrypted2.Bytes()) {
		t.Fatal("Two encryptions produced the same ciphertext, expected different due to randomness")
	}
}
