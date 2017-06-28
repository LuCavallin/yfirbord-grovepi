package connections

// ReadConnection is for connections that can read
type ReadConnection interface {
	Read(pin byte, mode string, size int) ([]byte, error)
}
