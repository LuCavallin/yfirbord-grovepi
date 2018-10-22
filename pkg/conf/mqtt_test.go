package conf

import (
	"github.com/icrowley/fake"
	"os"
	"testing"
)

var (
	host = "tcp://mqttbroker.com:12345"
	clientId = fake.UserName()
	username = fake.UserName()
	password = fake.Password(10, 10, false, false, false)
)

func TestMain(m *testing.M) {
	setEnvVariables()
	retCode := m.Run()
	unsetEnvVariables()
	os.Exit(retCode)
}

func setEnvVariables() {
	os.Setenv("MQTT_HOST", host)
	os.Setenv("MQTT_CLIENT_ID", clientId)
	os.Setenv("MQTT_USERNAME", username)
	os.Setenv("MQTT_PASSWORD", password)
}

func unsetEnvVariables() {
	os.Setenv("MQTT_HOST", "")
	os.Setenv("MQTT_CLIENT_ID", "")
	os.Setenv("MQTT_USERNAME", "")
	os.Setenv("MQTT_PASSWORD", "")
}

func TestNewMqttConfig(t *testing.T) {
	cfg := NewMqttConfig()
	if cfg.Host != host {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.Host, host)
	}
	if cfg.ClientId != clientId {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.ClientId, clientId)
	}
	if cfg.Username != username {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.Username, username)
	}
	if cfg.Password != password {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.Password, password)
	}
}
