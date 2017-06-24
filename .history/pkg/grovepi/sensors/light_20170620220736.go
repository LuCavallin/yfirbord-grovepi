package sensors

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

// Light is structure for DHT sensor
type Light struct {
	Sensor
}

func (o Light) Read(g *grovepi.GrovePi) {
	light, err := g.AnalogRead(grovepi.A2)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"light": light,
	}, nil
}
