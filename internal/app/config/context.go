package config

import (
	"context"

	log "github.com/sirupsen/logrus"
)

type configContextKeyType struct{}

var configContextKey = configContextKeyType{}

// FromContext - Возвращает конфиг из корневого контекста
func FromContext(ctx context.Context) *Config {
	cfgRaw := ctx.Value(configContextKey)
	cfg, ok := cfgRaw.(*Config)
	if ok {
		return cfg
	}

	log.Error("config FromContext executed, but no config in context")

	return nil
}

// WrapContext - Обогащает контекст конфигом
func WrapContext(ctx context.Context, cfg *Config) context.Context {
	return context.WithValue(ctx, configContextKey, cfg)
}
