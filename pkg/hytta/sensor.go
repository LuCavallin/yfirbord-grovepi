package hytta

// ReadableSensor interface for readable sensors
type ReadableSensor interface {
	Read() (int, error)
}