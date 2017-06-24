package sensors

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

// DHT is structure for DHT sensor
type DHT struct {
	sensor     Sensor
	connection *grovepi.GrovePi
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
