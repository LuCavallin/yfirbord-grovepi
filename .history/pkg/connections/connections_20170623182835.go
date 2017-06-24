package connections

// DigitalInput Interface for providing digital input
type DigitalInput interface {
	digitalRead(pin byte)
}

// Interface for providing digital input
type DigitalOutput interface {
}

// Interface for providing digital input
type AnalogInput interface {
}

// Interface for providing digital input
type AnalogOutput interface {
}

// Interface for providing digital input
type DHTInput interface {
}
