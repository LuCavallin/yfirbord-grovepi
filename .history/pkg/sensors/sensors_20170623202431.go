package sensors

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	Pin         string
	Mode        string
}

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

// NewSensor create a new sensor
func NewSensor(name string, description string, pin btye, mode string) Sensor {
	return Sensor{
		Name:        name,
		Description: description,
		Pin:         pin,
		Mode:        mode,
	}
}
