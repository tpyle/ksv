package localstorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"

	"errors"

	"github.com/mitchellh/mapstructure"
	"github.com/tpyle/ksv/lib/ksverrors"
)

const (
	FileStorageType = "file"
)

type FileStorage struct {
	Config *FileStorageConfig
}

type FileStorageConfig struct {
	FilePath string `mapstructure:"file_path"`
}

func (fs *FileStorage) LoadConfig(config map[string]interface{}) error {
	var fileStorageConfig FileStorageConfig
	err := mapstructure.Decode(config, &fileStorageConfig)
	if err != nil {
		return fmt.Errorf("error decoding file storage config: %w", err)
	}

	if fileStorageConfig.FilePath == "" {
		return fmt.Errorf("file_path is required for file storage")
	}

	fs.Config = &fileStorageConfig

	return nil
}

func (fs *FileStorage) Load() ([]byte, error) {
	file, err := os.Open(fs.Config.FilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, ksverrors.ErrEmptyLocalStorage
		}
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return io.ReadAll(file)
}

func (fs *FileStorage) Save(data []byte) error {
	var file *os.File

	fileDir := filepath.Dir(fs.Config.FilePath)
	if _, err := os.Stat(fileDir); os.IsNotExist(err) {
		if err := os.MkdirAll(fileDir, 0700); err != nil {
			return fmt.Errorf("error creating directory: %w", err)
		}
	}

	file, err := os.Create(fs.Config.FilePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}

	if _, err := file.Write(data); err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
