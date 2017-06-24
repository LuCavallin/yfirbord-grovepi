package sensors

import (
	"time"
)

// DHT is all that is needed for a new DHT sensor
type DHT struct {
	config sensors.Sensor
}

// NewDHT Inits a new DHT sensor
func NewDHT(sensor Sensors) DHT {
	return DHT{
		sensor: sensor
	}
}

func (o DHT) Read() (Measurement, error) {
	t, h, err := g.ReadDHT(o.config.pin)
	if err != nil {
		return nil, err
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}
