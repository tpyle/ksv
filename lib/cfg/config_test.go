package cfg_test

import (
	"path/filepath"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/tpyle/ksv/lib/cfg"
)

func TestLoadConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")

	// Test loading non-existent config file
	config, err := cfg.LoadConfig(configPath)
	assert.NoError(t, err)
	assert.NotNil(t, config)

	// Test loading existing config file
	viper.SetConfigFile(configPath)
	viper.Set("local_storage.type", "file")
	viper.Set("local_storage.config.file_path", "ksv.dat")
	assert.NoError(t, viper.WriteConfig())

	config, err = cfg.LoadConfig(configPath)
	assert.NoError(t, err)
	assert.NotNil(t, config)
	assert.Equal(t, "file", config.LocalStorageConfig.Type)
	assert.Equal(t, "ksv.dat", config.LocalStorageConfig.Config["file_path"])
}

func TestWriteDefaultConfig(t *testing.T) {
	tempDir := t.TempDir()
	configPath := filepath.Join(tempDir, "config.yaml")

	err := cfg.WriteDefaultConfig(configPath)
	assert.NoError(t, err)

	viper.SetConfigFile(configPath)
	assert.NoError(t, viper.ReadInConfig())

	assert.Equal(t, "file", viper.GetString("local_storage.type"))
	assert.Equal(t, "ksv.dat", viper.GetString("local_storage.config.file_path"))
}
