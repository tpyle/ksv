package localstorage

import (
	"github.com/tpyle/ksv/lib/ksverrors"
)

type LocalStorage interface {
	LoadConfig(config map[string]interface{}) error
	Load() ([]byte, error)
	Save(reader []byte) error
}

func GetLocalStorage(storageType string, config map[string]interface{}) (LocalStorage, error) {
	switch storageType {
	case FileStorageType:
		var fs FileStorage
		err := fs.LoadConfig(config)
		return &fs, err
	case KeyringStorageType:
		var ks KeyringStorage
		err := ks.LoadConfig(config)
		return &ks, err
	default:
		return nil, ksverrors.ErrUnsupportedLocalStorageType
	}
}
