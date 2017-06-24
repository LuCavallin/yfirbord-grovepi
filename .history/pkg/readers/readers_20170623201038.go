package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor, conn ) (Reader, error) {
	switch sensor.Mode {
	case condition:
		
	}
}
