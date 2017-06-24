package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
)

type Reader interface {
	Read() (sensors.Measurement, error)
}

func NewReader(sensor sensors.Sensor, conn connections.) (Reader, error) {
	switch sensor.Mode {
	case "dht":
		return DHT{sensor, conn: conn}
	case "light":

		
}


}
