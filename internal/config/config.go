package config

import (
	"sync"

	"github.com/spf13/viper"
)

var onceConfig sync.Once
var config *Config

type Config struct {
	// Prometheus configuration
	PrometheusHost string
	PrometheusPort int

	// Application configuration
	Debug              bool
	LogLevel           string
	ReconciliationLoop int
	// Services configuration
	CommandNames []string
	CommandPaths []string
}

func Get() *Config {
	onceConfig.Do(func() {
		viper.SetDefault("prometheus.host", "localhost")
		viper.SetDefault("prometheus.port", 9090)
		viper.SetDefault("debug", false)
		viper.SetDefault("logLevel", "INFO")
		viper.SetDefault("reconciliationLoop", 5)

		viper.SetConfigName("config")
		viper.SetConfigType("yaml")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.deucalion")
		viper.AddConfigPath("/etc/deucalion/")

		if err := viper.ReadInConfig(); err != nil {
			if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
				panic(err)
			}
		}

		config = &Config{
			PrometheusHost:     viper.GetString("prometheus.host"),
			PrometheusPort:     viper.GetInt("prometheus.port"),
			Debug:              viper.GetBool("debug"),
			LogLevel:           viper.GetString("logLevel"),
			ReconciliationLoop: viper.GetInt("reconciliationLoop"),

			CommandNames: viper.GetStringSlice("Commands"),
			CommandPaths: viper.GetStringSlice("CommandPaths"),
		}
	})

	return config
}
