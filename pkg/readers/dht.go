package readers

import (
	"fmt"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// DHT is structure for DHT sensor
type DHT struct {
	sensors.Sensor
	conn connections.DHTInput
}

func (o DHT) Read(c chan<- sensors.Measurement) error {
	t, h, err := o.conn.DHTRead(o.Pin)
	if err != nil {
		fmt.Printf("Couldn't read from DHT.\nERROR: %s \n", err)
	}

	c <- sensors.Measurement{
		"temperature": t,
		"humidity":    h,
	}
}
