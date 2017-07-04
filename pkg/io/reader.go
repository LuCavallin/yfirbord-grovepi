package io

import (
	"github.com/LuCavallin/yfirbord-grovepi/pkg/connections"
	"github.com/LuCavallin/yfirbord-grovepi/pkg/sensors"
	"github.com/LuCavallin/yfirbord-grovepi/pkg/sensors/parsers"
)

// Reader knows how to read from the connection and parse the content
type Reader struct {
	parsers map[string]parsers.Parser
	conn    connections.ReadConnection
}

//NewReader creates a new reader
func NewReader(parsers map[string]parsers.Parser, conn connections.ReadConnection) *Reader {
	return &Reader{
		parsers: parsers,
		conn:    conn,
	}
}

// Read, reads and parsers from sensor
// @TODO not sure about error handling here
func (r *Reader) Read(s sensors.Sensor) (parsers.Measurement, error) {
	raw, err := r.conn.Read(s.Pin, s.Mode, s.Size)
	if err != nil {
		return nil, err
	}

	return r.parsers[s.Parser].ToSI(raw), nil
}
