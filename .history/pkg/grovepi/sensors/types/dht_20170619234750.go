package sensors

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi/sensors"
)

type DHT struct {
	sensors.Sensor
}

func (o DHT) Read() {
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
