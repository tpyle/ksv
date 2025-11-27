package types

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"

	"github.com/tpyle/ksv/lib/enc"
	"github.com/tpyle/ksv/lib/ksverrors"
)

type KSV struct {
	Namespaces map[string][]byte `json:"namespaces"`

	// Default Namespace
	DefaultNamespace Namespace `json:"default_namespace"`
}

func DecryptKSV(input io.Reader, key []byte) (*KSV, error) {
	var ksv KSV
	err := enc.DecryptAndUnmarshal(input, key, &ksv)
	if err != nil {
		return nil, err
	}
	return &ksv, nil
}

func (k *KSV) Encrypt(output io.Writer, key []byte) error {
	return enc.MarshalAndEncrypt(k, key, output)
}

func (k *KSV) GetNamespace(name string) ([]byte, bool) {
	namespace, ok := k.Namespaces[name]
	return namespace, ok
}

func (k *KSV) HasNamespace(name string) bool {
	_, ok := k.Namespaces[name]
	return ok
}

func (k *KSV) AddNamespace(name string, namespace []byte) {
	if k.Namespaces == nil {
		k.Namespaces = make(map[string][]byte)
	}
	k.Namespaces[name] = namespace
}

func (k *KSV) RemoveNamespace(name string) {
	delete(k.Namespaces, name)
}

func (k *KSV) ReplaceNamespace(name string, namespace []byte) error {
	if !k.HasNamespace(name) {
		return fmt.Errorf("%w: %s", ksverrors.ErrNoSuchNamespace, name)
	}
	k.Namespaces[name] = namespace
	return nil
}

func (k *KSV) UnmarshalJSON(data []byte) error {

	aux := &struct {
		Namespaces       map[string]string `json:"namespaces"`
		DefaultNamespace Namespace         `json:"default_namespace"`
	}{
		Namespaces: make(map[string]string),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	k.Namespaces = make(map[string][]byte)
	for key, val := range aux.Namespaces {
		namespaceBytes, err := base64.StdEncoding.DecodeString(val)
		if err != nil {
			return err
		}
		k.Namespaces[key] = namespaceBytes
	}

	k.DefaultNamespace = aux.DefaultNamespace

	return nil
}

func (k *KSV) MarshalJSON() ([]byte, error) {
	aux := &struct {
		Namespaces       map[string]string `json:"namespaces"`
		DefaultNamespace Namespace         `json:"default_namespace"`
	}{
		Namespaces:       make(map[string]string),
		DefaultNamespace: k.DefaultNamespace,
	}

	for key, val := range k.Namespaces {
		aux.Namespaces[key] = base64.StdEncoding.EncodeToString(val)
	}

	return json.Marshal(aux)
}

func (k *KSV) Validate() []error {
	var errs []error
	errs = append(errs, k.DefaultNamespace.Validate()...)

	return errs
}
