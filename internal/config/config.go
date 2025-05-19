package config

import (
	"github.com/caarlos0/env/v11"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	ConnectionString string `env:"CONNECTION_STRING"`
	WeatherApiKey    string `env:"WEATHER_API_KEY"`
	SenderConfig
}

type SenderConfig struct {
	SenderEmail string `env:"SENDER_EMAIL"`
	SenderPass  string `env:"SENDER_PASS"`
}

var once sync.Once

var configInstance *Config

func GetConfig() *Config {
	if configInstance == nil {
		once.Do(func() {
			log.Println("Creating config instance now.")

			var cfg Config
			var senderCfg SenderConfig

			if err := env.Parse(&cfg); err != nil {
				log.Fatal(err)
			}

			if err := env.Parse(&senderCfg); err != nil {
				log.Fatal(err)
			}

			cfg.SenderConfig = senderCfg
			configInstance = &cfg
		})

	}

	return configInstance
}
