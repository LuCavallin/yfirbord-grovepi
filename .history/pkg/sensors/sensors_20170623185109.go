package sensors

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	Pin         byte
	mode        string
}
