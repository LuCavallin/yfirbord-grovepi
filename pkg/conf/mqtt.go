package conf

import (
	"github.com/caarlos0/env"
	"github.com/prometheus/common/log"
)

// Config contains configuration for connecting to a MQTT broker
type Mqtt struct {
	Host     string `env:"MQTT_HOST"`
	ClientId string `env:"MQTT_CLIENT_ID"`
	Username string `env:"MQTT_USERNAME"`
	Password string `env:"MQTT_PASSWORD"`
}

// NewMqttConfig creates a new configuration for the MQTT broker from env variables
func NewMqttConfig() (*Mqtt) {
	cfg := Mqtt{}
	err := env.Parse(&cfg)
	if err != nil {
		log.Fatal("Could not read MQTT configuration variables.")
	}

	return &cfg
}