package connections

// DigitalInput Interface for providing digital input
type DigitalInput interface {
	digitalRead(pin byte)
}

// DigitalOutput Interface for providing digital output
type DigitalOutput interface {
	digitalInput
}

// AnalogInput Interface for providing analog input
type AnalogInput interface {
}

// AnalogOutput Interface for providing analog output
type AnalogOutput interface {
}

// DHTInput Interface for providing DHT input
type DHTInput interface {
}
