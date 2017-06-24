package sensors

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	pin         byte
	mode        string
}
