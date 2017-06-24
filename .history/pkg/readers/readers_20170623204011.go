package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor, conn *connections.GrovePi) (Reader, error) {
	var reader Reader

	switch sensor.Mode {
	case "dht":
		reader = DHT{sensor, conn}
		break
	case "light":
		reader = Light{sensor, conn}
		break
	default:
		return nil, nil
	}

	return reader, nil
}
