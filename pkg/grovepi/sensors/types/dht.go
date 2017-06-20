package types

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi/sensors"
)

// DHT is structure for DHT sensor
type DHT struct {
	sensors.Sensor
}

func (o DHT) Read(g *grovepi.GrovePi) {
	t, h, err := g.ReadDHT(o.Pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}
