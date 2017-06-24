package sensors

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	name        string
	description string
	pin         string
	command     string
}

type Sensors []Sensor

// InputSensor provides an interface for reading from all sensors
type InputSensor interface {
	Read() (Measurement, error)
}

// NewSensor returns configuration for a new sensor
func NewSensor(name string, description string, pin string, command string) Sensor {
	return Sensor{
		name:        name,
		description: description,
		pin:         pin,
		command:     command,
	}
}
