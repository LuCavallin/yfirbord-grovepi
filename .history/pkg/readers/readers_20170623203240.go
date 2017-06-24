package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor) (Reader, error) {
	switch sensor.Mode {
	case "dht":
		return new(DHT{sensor, nil})
	case "light":

	}

}
