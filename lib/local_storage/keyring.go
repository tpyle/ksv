package localstorage

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/tpyle/ksv/lib/ksverrors"
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

func (fs *KeyringStorage) Load() ([]byte, error) {
	var dataBuilder strings.Builder
	for i := 0; ; i++ {
		val, err := keyring.Get(KeyringService, fmt.Sprintf("ksv_chunk_%d", i))
		if err != nil {
			if err == keyring.ErrNotFound {
				if i == 0 {
					return nil, ksverrors.ErrEmptyLocalStorage
				}
				break
			}
			return nil, fmt.Errorf("error getting keyring value: %w", err)
		}
		dataBuilder.WriteString(val)
	}

	data, err := base64.StdEncoding.DecodeString(dataBuilder.String())
	if err != nil {
		return nil, fmt.Errorf("error decoding base64 data: %w", err)
	}

	return data, nil
}

func (fs *KeyringStorage) Save(data []byte) error {
	base64Data := base64.StdEncoding.EncodeToString(data)

	chunkNum := 0
	for _, chunk := range ChunkString(base64Data, ChunkSize) {
		err := keyring.Set(KeyringService, fmt.Sprintf("ksv_chunk_%d", chunkNum), chunk)
		if err != nil {
			return fmt.Errorf("error setting keyring value: %w", err)
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
