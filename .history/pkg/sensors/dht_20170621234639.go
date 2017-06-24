package sensors

import (
	"time"
)

// DHT is structure for DHT sensor
type DHT struct {
	Sensor
}

func (o DHT) Read() (Measurement, error) {
	t, h, err := o.conn.ReadDHT(o.Pin)
	if err != nil {
		panic(err)
	}
	time.Sleep(500 * time.Millisecond)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}, nil
}
