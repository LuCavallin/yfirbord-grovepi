package readers

import (
	"time"
)

// Light is structure for DHT sensor
type Light struct {
	Sensor

	conn *connections.AnalogReader
}

func (o Light) Read() (Measurement, error) {
	light, err := o.conn.AnalogRead(o.pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"light": light,
	}, nil
}
