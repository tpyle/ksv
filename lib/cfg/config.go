package cfg

import (
	"fmt"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	localstorage "github.com/tpyle/ksv/lib/local_storage"
)

type Config struct {
	LocalStorageConfig *LocalStorageConfig `mapstructure:"local_storage"`
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
			panic(err)
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

	return &config, nil
}

// WriteDefaultConfig writes the default configuration to a file.
func WriteDefaultConfig(configPath string) error {
	viper.SetConfigFile(configPath)

	viper.SetDefault("local_storage.type", "file")
	viper.SetDefault("local_storage.config.file_path", "ksv.dat")

	if err := viper.WriteConfigAs(configPath); err != nil {
		return fmt.Errorf("error writing default config file: %w", err)
	}

	return nil
}
