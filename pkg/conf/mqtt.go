package conf

import (
	"github.com/caarlos0/env"
)

// Config contains configuration for connecting to a MQTT broker
type Mqtt struct {
	Host     string `env:"MQTT_HOST"`
	ClientId string `env:"MQTT_CLIENT_ID"`
	Username string `env:"MQTT_USERNAME"`
	Password string `env:"MQTT_PASSWORD"`
}

// NewMqttConfig creates a new configuration for the MQTT broker from env variables
func NewMqttConfig() (*Mqtt, error) {
	cfg := Mqtt{}
	err := env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}