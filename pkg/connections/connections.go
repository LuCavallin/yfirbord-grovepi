package connections

// ReadConnection is for connections that can read
type ReadConnection interface {
	Read(byte, string, int) ([]byte, error)
}
