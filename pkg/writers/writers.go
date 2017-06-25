package writers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// Writer interface for input sensors
type Writer interface {
	Write(val byte) error
}

// NewWriter Creates a new reader given sensor and connection
func NewWriter(sensor sensors.Sensor, conn *connections.GrovePi) (Writer, error) {
	var writer Writer

	switch sensor.Mode {
	default:
		return nil, nil
	}

	return writer, nil
}
