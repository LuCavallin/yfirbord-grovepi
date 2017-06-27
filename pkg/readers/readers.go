package readers

import (
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/readers/parsers"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

// Reader knows how to read from the connection and parse the content
type Reader struct {
	parsers map[string]parsers.Parser
	conn    connections.ReadConnection
}

//NewReader creates a new reader
func NewReader(conn connections.ReadConnection) *Reader {
	return &Reader{
		parsers: getParsers(),
		conn:    conn,
	}
}

// Read, reads and parsers from sensor
// @TODO not sure about error handling here
func (r *Reader) Read(s sensors.Sensor, c chan<- parsers.Measurement) {
	raw, err := r.conn.Read(s.Pin, s.Mode, s.Size)
	si := r.parsers[s.Type].ToSI(raw)

	if err != nil {
		// log
		panic(err)
	}

	c <- si
}

// TODO Factory logic
func getParsers() map[string]parsers.Parser {
	return map[string]parsers.Parser{
		"dht":   new(parsers.DHT),
		"light": new(parsers.Light),
	}
}
