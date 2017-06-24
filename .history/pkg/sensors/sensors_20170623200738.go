package sensors

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	Mode        string
}

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

func NewSensor(name string, description string, mode string) {
	return Sensor{
		Name:        name,
		Description: description,
		Mode:        mode,
	}, nil
}
