package readers

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// DHT is structure for DHT sensor
type DHT struct {
	sensors.Sensor
	//
	conn connections.DHTInput
}

func (o DHT) Read() (Measurement, error) {
	t, h, err := o.conn.DHTRead(o.Pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}
