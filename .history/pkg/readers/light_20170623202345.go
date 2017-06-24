package readers

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// Light is structure for DHT sensor
type Light struct {
	sensors.Sensor
	conn connections.AnalogInput
}

func (o Light) Read() (sensors.Measurement, error) {
	light, err := o.conn.AnalogRead(o.pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return sensors.Measurement{
		"light": light,
	}, nil
}
