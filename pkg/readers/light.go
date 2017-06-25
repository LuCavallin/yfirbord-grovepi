package readers

import (
	"fmt"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// Light is structure for DHT sensor
type Light struct {
	sensors.Sensor
	conn connections.AnalogInput
}

func (o Light) Read(c chan<- sensors.Measurement) error {
	light, err := o.conn.AnalogRead(o.Pin)
	if err != nil {
		fmt.Printf("Couldn't read from Light.\nERROR: %s \n", err)
	}

	c <- sensors.Measurement{
		"light": light,
	}
}
