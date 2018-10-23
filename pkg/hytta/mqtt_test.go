package hytta

import (
	"github.com/icrowley/fake"
	"os"
	"testing"
)

var (
	host = "tcp://mqttbroker.com:12345"
	clientID = fake.UserName()
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
	os.Setenv("MQTT_CLIENT_ID", clientID)
	os.Setenv("MQTT_USERNAME", username)
	os.Setenv("MQTT_PASSWORD", password)
}

func unsetEnvVariables() {
	os.Unsetenv("MQTT_HOST")
	os.Unsetenv("MQTT_CLIENT_ID")
	os.Unsetenv("MQTT_USERNAME")
	os.Unsetenv("MQTT_PASSWORD")
}

func TestNewMqttConfig(t *testing.T) {
	cfg := NewMqttConfig()
	if cfg.Host != host {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.Host, host)
	}
	if cfg.ClientID != clientID {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.ClientID, clientID)
	}
	if cfg.Username != username {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.Username, username)
	}
	if cfg.Password != password {
		t.Errorf("cfg.Host was incorrect, got: %s, want: %s.", cfg.Password, password)
	}
}
