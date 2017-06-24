package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor) (Reader, error) {
	conn := connections.GrovePi
	switch sensor.Mode {
	case "dht":
		return DHT{sensor, conn: conn}, nil
	case "light":

	}

}
