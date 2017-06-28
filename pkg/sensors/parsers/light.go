package parsers

// Light is structure for DHT sensor
type Light struct{}

// ToSI converts raw output to SI units
func (p Light) ToSI(raw []byte) Measurement {
	light := ((int(raw[1]) << 8) | int(raw[2]))

	return Measurement{
		"light": light,
	}
}
