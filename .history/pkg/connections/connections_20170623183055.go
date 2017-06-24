package connections

// DigitalInput Interface for providing digital input
type DigitalInput interface {
	digitalRead(pin byte)
}

// DigitalOutput Interface for providing digital output
type DigitalOutput interface {
	digitalWrite(pin byte)
}

// AnalogInput Interface for providing analog input
type AnalogInput interface {
	analogInput(pin byte)
}

// AnalogOutput Interface for providing analog output
type AnalogOutput interface {
	digitalInput(pin byte)
}

// DHTInput Interface for providing DHT input
type DHTInput interface {
	digitalInput(pin byte)
}
