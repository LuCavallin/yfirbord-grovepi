package connections

// AnalogInput Interface for providing analog input
type AnalogInput interface {
	analogRead(pin byte) (int, error)
}

// DigitalInput Interface for providing digital input
type DigitalInput interface {
	digitalRead(pin byte) ([]byte, error)
}

// DigitalOutput Interface for providing digital output
type DigitalOutput interface {
	digitalWrite(pin byte)
}

// DHTInput Interface for providing DHT input
type DHTInput interface {
	dhtRead(pin byte)
}
