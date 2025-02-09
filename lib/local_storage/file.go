package localstorage

import (
	"fmt"
	"io"
	"os"

	"github.com/mitchellh/mapstructure"
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

func (fs *FileStorage) Load() (io.Reader, error) {
	file, err := os.Open(fs.Config.FilePath)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %w", err)
	}

	return file, nil
}

func (fs *FileStorage) Save(reader io.Reader) error {
	var file *os.File
	if _, err := os.Stat(fs.Config.FilePath); os.IsNotExist(err) {
		file, err = os.Create(fs.Config.FilePath)
		if err != nil {
			return fmt.Errorf("error creating file: %w", err)
		}
	} else {
		file, err = os.Open(fs.Config.FilePath)
		if err != nil {
			return fmt.Errorf("error opening file: %w", err)
		}
	}

	if _, err := file.Seek(0, 0); err != nil {
		return fmt.Errorf("error seeking to beginning of file: %w", err)
	}

	if err := file.Truncate(0); err != nil {
		return fmt.Errorf("error truncating file: %w", err)
	}

	_, err := io.Copy(file, reader)
	if err != nil {
		return fmt.Errorf("error writing to file: %w", err)
	}

	return nil
}
