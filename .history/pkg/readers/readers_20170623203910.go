package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor, conn connections.GrovePi) (Reader, error) {
	reader Reader

	switch sensor.Mode {
	case "dht":
		return DHT{sensor, conn}, nil
		break
	case "light":
		return Light{sensor, conn}, nil
		break
	default:

	}
}
