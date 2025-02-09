package localstorage

import (
	"io"

	"github.com/tpyle/ksv/lib/errors"
)

type LocalStorage interface {
	LoadConfig(config map[string]interface{}) error
	Load() (io.Reader, error)
	Save(reader io.Reader) error
}

func GetLocalStorage(storageType string, config map[string]interface{}) (LocalStorage, error) {
	switch storageType {
	case FileStorageType:
		var fs FileStorage
		err := fs.LoadConfig(config)
		return &FileStorage{}, err
	case KeyringStorageType:
		var ks KeyringStorage
		err := ks.LoadConfig(config)
		return &KeyringStorage{}, err
	default:
		return nil, errors.ErrUnsupportedLocalStorageType
	}
}
