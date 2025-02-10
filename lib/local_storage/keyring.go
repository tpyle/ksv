package localstorage

import (
	"encoding/base64"
	"fmt"
	"io"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/tpyle/ksv/lib/errors"
	"github.com/zalando/go-keyring"
)

const (
	KeyringStorageType = "keyring"
	KeyringService     = "ksv"
	ChunkSize          = 2048 // 2k for mac & windows, 64k for unix/linux
)

type KeyringStorage struct {
	Config *KeyringStorageConfig
}

type KeyringStorageConfig struct {
	StorageSize int `mapstructure:"storage_size"`
}

func (fs *KeyringStorage) LoadConfig(config map[string]interface{}) error {
	var storageConfig KeyringStorageConfig
	err := mapstructure.Decode(config, &storageConfig)
	if err != nil {
		return fmt.Errorf("error decoding file storage config: %w", err)
	}

	if storageConfig.StorageSize == 0 {
		logrus.Debugf("storage_size is unset, assuming unlimited storage")
	}

	fs.Config = &storageConfig

	return nil
}

func (fs *KeyringStorage) Load() (io.Reader, error) {
	var dataBuilder strings.Builder
	for i := 0; ; i++ {
		val, err := keyring.Get(KeyringService, fmt.Sprintf("ksv_chunk_%d", i))
		if err != nil {
			if err == keyring.ErrNotFound {
				if i == 0 {
					return nil, errors.ErrEmptyLocalStorage
				}
				break
			}
			return nil, fmt.Errorf("error getting keyring value: %w", err)
		}
		dataBuilder.WriteString(val)
	}

	return base64.NewDecoder(base64.StdEncoding, strings.NewReader(dataBuilder.String())), nil
}

func (fs *KeyringStorage) Save(reader io.Reader) error {
	// Create a pipe
	pr, pw := io.Pipe()

	// Create a base64 encoder that writes to the pipe writer
	base64Encoder := base64.NewEncoder(base64.StdEncoding, pw)

	// Channel to capture any errors from the goroutines
	errChan := make(chan error, 1)

	// Goroutine to read from the reader and write to the base64 encoder
	go func() {
		defer pw.Close()
		_, err := io.Copy(base64Encoder, reader)
		if err != nil {
			errChan <- fmt.Errorf("error copying data to base64 encoder: %w", err)
			return
		}
		err = base64Encoder.Close()
		if err != nil {
			errChan <- fmt.Errorf("error closing base64 encoder: %w", err)
			return
		}
		errChan <- nil
	}()

	// Goroutine to read from the pipe reader and write chunks to the keyring
	go func() {
		defer pr.Close()

		chunk := make([]byte, ChunkSize)
		chunkNum := 0
		for {
			n, err := io.ReadFull(pr, chunk)
			if err != nil {
				if err == io.EOF || err == io.ErrUnexpectedEOF {
					// Only write the actual bytes read
					err = keyring.Set(KeyringService, fmt.Sprintf("ksv_chunk_%d", chunkNum), string(chunk[:n]))
					if err != nil {
						errChan <- fmt.Errorf("error setting keyring value: %w", err)
						return
					}
					chunkNum++
					break
				}
				errChan <- fmt.Errorf("error reading from pipe: %w", err)
				return
			}

			// Only write the actual bytes read
			err = keyring.Set(KeyringService, fmt.Sprintf("ksv_chunk_%d", chunkNum), string(chunk[:n]))
			if err != nil {
				errChan <- fmt.Errorf("error setting keyring value: %w", err)
				return
			}
			chunkNum++
		}

		// Delete any remaining old chunks
		for i := chunkNum; ; i++ {
			err := keyring.Delete(KeyringService, fmt.Sprintf("ksv_chunk_%d", i))
			if err != nil {
				if err == keyring.ErrNotFound {
					break
				}
				logrus.WithError(err).Error("error deleting old keyring value")
			}
		}
		errChan <- nil
	}()

	// Wait for both goroutines to finish and check for errors
	for i := 0; i < 2; i++ {
		if err := <-errChan; err != nil {
			return err
		}
	}

	return nil
}

func ChunkString(s string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(s); i += chunkSize {
		end := i + chunkSize
		if end > len(s) {
			end = len(s)
		}
		chunks = append(chunks, s[i:end])
	}
	return chunks
}
