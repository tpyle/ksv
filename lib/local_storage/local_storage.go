package localstorage

import (
	"io"

	"github.com/tpyle/ksv/lib/ksverrors"
)

type LocalStorage interface {
	LoadConfig(config map[string]interface{}) error
	Load() ([]byte, error)
	LoadReader() (io.Reader, error)
	Save(reader []byte) error
	SaveWriter(reader io.Reader) error
}

func GetLocalStorage(storageType string, config map[string]interface{}) (LocalStorage, error) {
	var storage LocalStorage
	var err error
	switch storageType {
	case FileStorageType:
		var fs FileStorage
		err = fs.LoadConfig(config)
		storage = &fs
	case KeyringStorageType:
		var ks KeyringStorage
		err = ks.LoadConfig(config)
		storage = &ks
	default:
		err = ksverrors.ErrUnsupportedLocalStorageType
		storage = nil
	}
	return storage, err
}
