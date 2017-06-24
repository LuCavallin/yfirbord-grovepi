package connections

// AnalogInput Interface for providing analog input
type AnalogInput interface {
	AnalogRead(pin byte) (int, error)
}

// DigitalInput Interface for providing digital input
type DigitalInput interface {
	DigitalRead(pin byte) ([]byte, error)
}

// DigitalOutput Interface for providing digital output
type DigitalOutput interface {
	DigitalWrite(pin byte) error
}

// DHTInput Interface for providing DHT input
// @TODO, this should disappear
type DHTInput interface {
	DhtRead(pin byte) (float32, float32, error)
}
