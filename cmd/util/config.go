package util

import (
	"context"

	"github.com/tpyle/ksv/lib/cfg"
)

type KSVConfigKey string

const (
	// ConfigKey is the key used to store the config in the context
	ConfigKey = KSVConfigKey("config")
)

func AttachConfigToContext(config *cfg.Config, ctx context.Context) context.Context {
	return context.WithValue(context.Background(), ConfigKey, config)
}

func GetConfigFromContext(ctx context.Context) *cfg.Config {
	config, ok := ctx.Value(ConfigKey).(*cfg.Config)
	if !ok {
		// This should never be called
		panic("config not found in context")
	}

	return config
}
