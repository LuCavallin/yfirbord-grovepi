package parsers

import (
	"encoding/binary"
	"math"
)

// DHT is interface for DHT sensor
type DHT struct{}

const (
	startTemperature = 1
	endTemperature   = 5
	startHumidity    = 5
	endHumidity      = 9
)

// ToSI converts raw output to SI units
func (p DHT) ToSI(raw []byte) Measurement {
	temperatureData := binary.LittleEndian.Uint32(raw[startTemperature:endTemperature])
	humidityData := binary.LittleEndian.Uint32(raw[startHumidity:endHumidity])

	t := math.Float32frombits(temperatureData)
	h := math.Float32frombits(humidityData)

	return Measurement{
		"temperature": t,
		"humidity":    h,
	}
}
