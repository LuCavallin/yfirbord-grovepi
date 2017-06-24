package sensors

import (
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

// DHT is structure for DHT sensor
type DHT struct {
	Sensor
}

func (o DHT) Read(g *grovepi.GrovePi) {
	t, h, err := o.conn.ReadDHT(pin byte)(o.Pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}
