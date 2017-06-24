package sensors

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi/sensors"
)

// DHT is structure for DHT sensor
type DHT struct {
	sensors.Sensor
}

func (o DHT) Read(*grovepi.GrovePi) {
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
