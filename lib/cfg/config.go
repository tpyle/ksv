package cfg

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	localstorage "github.com/tpyle/ksv/lib/local_storage"
)

type Config struct {
	LocalStorageConfig LocalStorageConfig `mapstructure:"local_storage"`
}

func (c *Config) Validate() error {
	if c.LocalStorageConfig.Type == "" {
		return fmt.Errorf("local_storage.type is required")
	}

	_, err := localstorage.GetLocalStorage(c.LocalStorageConfig.Type, c.LocalStorageConfig.Config)
	if err != nil {
		return err
	}

	return nil
}

type LocalStorageConfig struct {
	Type         string                    `mapstructure:"type"`
	Config       map[string]interface{}    `mapstructure:"config"`
	LocalStorage localstorage.LocalStorage `mapstructure:"-"`
}

// LoadConfig loads the configuration from a file.
func LoadConfig(configPath string) (*Config, error) {
	fi, err := os.Stat(configPath)
	if os.IsNotExist(err) || fi.IsDir() {
		logrus.Info("Config file does not exist, creating default config")
		if err := WriteDefaultConfig(configPath); err != nil {
			return nil, fmt.Errorf("error writing default config: %w", err)
		}
	}

	viper.SetConfigFile(configPath)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.SetEnvPrefix("KSV")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		return nil, fmt.Errorf("error unmarshaling config: %w", err)
	}

	config.LocalStorageConfig.LocalStorage, err = localstorage.GetLocalStorage(config.LocalStorageConfig.Type, config.LocalStorageConfig.Config)
	if err != nil {
		return nil, fmt.Errorf("error getting local storage: %w", err)
	}

	return &config, nil
}

// getDefaultFilePath returns the default file path based on the operating system.
func getDefaultFilePath() string {
	switch runtime.GOOS {
	case "windows":
		return filepath.Join(os.Getenv("AppData"), "ksv", "ksv.dat")
	case "darwin":
		return filepath.Join(os.Getenv("HOME"), "Library", "Application Support", "ksv", "ksv.dat")
	default: // "linux" and other Unix-like systems
		return filepath.Join(os.Getenv("HOME"), ".local", "share", "ksv", "ksv.dat")
	}
}

// WriteDefaultConfig writes the default configuration to a file.
func WriteDefaultConfig(configPath string) error {
	viper.SetConfigFile(configPath)

	viper.SetDefault("local_storage.type", "file")
	viper.SetDefault("local_storage.config.file_path", getDefaultFilePath())

	if err := os.MkdirAll(filepath.Dir(configPath), 0700); err != nil {
		return fmt.Errorf("error creating config directory: %w", err)
	}

	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	return nil
}
