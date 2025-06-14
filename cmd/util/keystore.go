package util

import (
	"context"
	"encoding/json"

	"github.com/tpyle/ksv/lib/cfg"
	"github.com/tpyle/ksv/lib/ksverrors"
	"github.com/tpyle/ksv/lib/types"
)

func GetKeystoreFromContext(ctx context.Context) (*types.KSV, error) {
	config := GetConfigFromContext(ctx) // Ensure config is loaded in context
	return GetKeystore(*config)
}

// Loads the keystore from the local storage specified in the config
func GetKeystore(config cfg.Config) (*types.KSV, error) {
	ksv := &types.KSV{}
	data, err := config.LocalStorageConfig.LocalStorage.Load()
	if err != nil {
		if err == ksverrors.ErrEmptyLocalStorage {
			// If the local storage is empty, return an empty keystore
			return ksv, nil
		}
		return ksv, err
	}

	err = json.Unmarshal(data, ksv)
	return ksv, err
}

func SaveKeystoreFromContext(ctx context.Context, ksv *types.KSV) error {
	config := GetConfigFromContext(ctx) // Ensure config is loaded in context
	return SaveKeystore(*config, ksv)
}

func SaveKeystore(config cfg.Config, ksv *types.KSV) error {
	data, err := json.Marshal(ksv)
	if err != nil {
		return err
	}

	return config.LocalStorageConfig.LocalStorage.Save(data)
}
