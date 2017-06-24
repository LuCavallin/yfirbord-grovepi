package readers

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// Light is structure for DHT sensor
type Light struct {
	sensors.Sensor
	Pin  byte
	conn connections.AnalogInput
}

func (o Light) Read() (Measurement, error) {
	light, err := o.conn.AnalogRead(o.Pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"light": light,
	}, nil
}
