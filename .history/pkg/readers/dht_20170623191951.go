package readers

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// DHT is structure for DHT sensor
type DHT struct {
	sensors.Sensor
	Pin  byte
	conn connections.DHTInput
}

func (o DHT) Read() (sensors.Measurement, error) {
	t, h, err := o.conn.DHTRead(o.pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return sensors.Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}
