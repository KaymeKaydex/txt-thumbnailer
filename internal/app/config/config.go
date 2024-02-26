package config

import (
	"context"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/getsentry/sentry-go"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/KaymeKaydex/txt-thumbnailer/internal/app/ds"
)

// Config - структура конфигурации;
// Содержит все конфигурационные данные о сервисе;
// автоподгружается при изменении исходного файла
type Config struct {
	AppInfo ds.AppInfo `yaml:"-"` // AppInfo должно быть обернуто вручную. Это ldflags

	ListenerConfig   ListenerConfig   `yaml:"listener" mapstructure:"listener"`
	MonitoringConfig MonitoringConfig `yaml:"monitoring" mapstructure:"monitoring"`
}

type MonitoringConfig struct {
	Address string `yaml:"address" mapstructure:"address"`
}

type ListenerConfig struct {
	LogLevel string `yaml:"log_level" mapstructure:"log_level"`
	Address  string `yaml:"address" mapstructure:"address"`
}

func (c *Config) WithAppInfo(info ds.AppInfo) *Config {
	c.AppInfo = info

	return c
}

// New cоздаёт новый объект конфигурации, загружая данные из файла конфигурации
func New(ctx context.Context, fPath string) (*Config, error) {
	// Если file path пустой, то возвращаем дефолтный конфиг
	if fPath == "" {
		return nil, fmt.Errorf("no config for no dev version")
	}
	viper.SetConfigType("yaml")

	viper.SetConfigFile(fPath)
	viper.WatchConfig()

	err := viper.ReadInConfig()
	if err != nil {
		sentry.CaptureException(err)

		return nil, err
	}

	cfg := &Config{}
	err = viper.Unmarshal(cfg)
	if err != nil {
		sentry.CaptureException(err)

		return nil, err
	}

	cfg.setLogLevel(cfg.ListenerConfig.LogLevel)

	viper.OnConfigChange(cfg.onConfigChange)

	log.WithContext(ctx).Infof("config parsed from %s", fPath)
	return cfg, nil
}

// GetVersion удовлетворяет интерфейсу стандартного /version обработчика
func (c *Config) GetVersion() string {
	return c.AppInfo.Version
}

// GetLogLevel удовлетворяет интерфейсу стандартного /version обработчика
func (c *Config) GetLogLevel() string {
	return c.ListenerConfig.LogLevel
}

// Запускает обновление данных в объекте конфигурации при изменении исходного файла с данными
func (c *Config) onConfigChange(_ fsnotify.Event) {
	log.Debug("on config changes event")
	err := viper.Unmarshal(c)
	if err != nil {
		log.Error("on config change error")
		sentry.CaptureException(err)

		return
	}

	c.setLogLevel(c.ListenerConfig.LogLevel)
}

// SetLogLevel setup log level for app
func (c *Config) setLogLevel(logLevel string) {
	foundLogLevel, ok := LogLevelMap[logLevel]
	if !ok {
		log.Errorf("incorrect log level %s", logLevel)

		return
	}

	log.SetLevel(foundLogLevel)
}

// LogLevelMap Содержит разрешённые уровни логирования;
// Чем выше уровень, тем больше выводится логов (например, INFO это INFO + WARN + ERROR)
var LogLevelMap = map[string]log.Level{
	"debug": log.DebugLevel,
	"info":  log.InfoLevel,
	"warn":  log.WarnLevel,
	"error": log.ErrorLevel,
}
