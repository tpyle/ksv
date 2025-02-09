package remotestorage

import "io"

type RemoteStorage interface {
	Load(storageConfig map[string]interface{}) (io.Reader, error)
	Save(storageConfig map[string]interface{}, writer io.Writer) error
}
