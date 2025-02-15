package types

import (
	"encoding/base64"
	"encoding/json"
)

type KSV struct {
	Namespaces map[string][]byte `json:"namespaces"`

	// Default Namespace
	DefaultNamespace Namespace `json:"default_namespace"`
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
