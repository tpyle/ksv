package util

import (
	"context"

	"github.com/tpyle/ksv/lib/cfg"
)

const (
	// ConfigKey is the key used to store the config in the context
	ConfigKey = "config"
)

func AttachConfigToContext(config *cfg.Config, ctx context.Context) context.Context {
	return context.WithValue(context.Background(), ConfigKey, config)
}

func GetConfigFromContext(ctx context.Context) *cfg.Config {
	config, ok := ctx.Value(ConfigKey).(*cfg.Config)
	if !ok {
		panic("config not found in context")
	}

	return config
}
