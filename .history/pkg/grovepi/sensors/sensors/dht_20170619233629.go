package sensors

import (
	"time"
)

func (o Sensor) Read() (Measurement, error) {
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
