package sensors

import (
	"time"
)

// DHT is structure for DHT sensor
type DHT struct {
	Sensor
	//
	conn *connections.grovepi
}

func (o DHT) Read() (Measurement, error) {
	t, h, err := o.conn.ReadDHT(o.pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}

// InputSensor provides an interface for reading from all sensors
type InputSensor interface {
	Read() (Measurement, error)
}
