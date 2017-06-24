package grovepi

import (
	"encoding/binary"
	"math"
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi/sensors"
	"github.com/mrmorphic/hwio"
	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

// Pins
const (
	A0 = 0
	A1 = 1
	A2 = 2

	D2 = 2
	D3 = 3
	D4 = 4
	D5 = 5
	D6 = 6
	D7 = 7
	D8 = 8
)

// Commands format
const (
	CommandDigitalRead  = 1
	CommandDigitalWrite = 2
	CommandAnalogRead   = 3
	CommandAnalogWrite  = 4
	CommandPinMode      = 5
	CommandDHTRead      = 40
)

// GrovePi struct is used for handling the connection with board
type GrovePi struct {
	i2cmodule hwio.I2CModule
	i2cDevice hwio.I2CDevice
}

// Init initializes the GrovePi
func Init(config Config) (*GrovePi, error) {
	grovePi := new(GrovePi)
	m, err := hwio.GetModule("i2c")
	if err != nil {
		return nil, err
	}

	grovePi.i2cmodule = m.(hwio.I2CModule)
	err = grovePi.i2cmodule.Enable()
	if err != nil {
		return nil, err
	}

	grovePi.i2cDevice = grovePi.i2cmodule.GetDevice(config.Address)
	return grovePi, nil
}

// Close closes the connection with the GrovePi
func (grovePi *GrovePi) Close() {
	grovePi.i2cmodule.Disable()
	hwio.CloseAll()
}

// AnalogRead reads analogically to the GrovePi
func (grovePi *GrovePi) analogRead(pin byte) (int, error) {
	b := []byte{CommandAnalogRead, pin, 0, 0}
	err := grovePi.i2cDevice.Write(1, b)
	if err != nil {
		return 0, err
	}
	time.Sleep(100 * time.Millisecond)

	grovePi.i2cDevice.ReadByte(1)
	val, err := grovePi.i2cDevice.Read(1, 4)
	if err != nil {
		return 0, err
	}

	return ((int(val[1]) << 8) | int(val[2])), nil
}

// DigitalRead reads digitally to the GrovePi
func (grovePi *GrovePi) digitalRead(pin byte) ([]byte, error) {
	b := []byte{CommandDigitalRead, pin, 0, 0}
	err := grovePi.i2cDevice.Write(1, b)
	if err != nil {
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)

	// TODO set size via parameter, it's better
	return grovePi.i2cDevice.Read(1, 1)
}

// DigitalWrite writes digitally to the GrovePi
func (grovePi *GrovePi) digitalWrite(pin byte, val byte) error {
	b := []byte{CommandDigitalWrite, pin, val, 0}
	err := grovePi.i2cDevice.Write(1, b)
	time.Sleep(100 * time.Millisecond)
	return err
}

// ReadDHT returns temperature and humidity from DHT sensor
func (grovePi *GrovePi) readDHT(pin byte) (float32, float32, error) {
	cmd := []byte{CommandDHTRead, pin, 0, 0}

	// prepare and read raw data
	err := grovePi.i2cDevice.Write(1, cmd)
	if err != nil {
		return 0, 0, err
	}
	time.Sleep(600 * time.Millisecond)
	grovePi.i2cDevice.ReadByte(1)
	time.Sleep(100 * time.Millisecond)
	rawdata, err := grovePi.i2cDevice.Read(1, 9)
	if err != nil {
		return 0, 0, err
	}

	temperatureData := binary.LittleEndian.Uint32(rawdata[1:5])
	t := math.Float32frombits(temperatureData)

	humidityData := binary.LittleEndian.Uint32(rawdata[5:9])
	h := math.Float32frombits(humidityData)

	return t, h, nil
}

// ReadFromSensor reads from a given sensor
func (grovePi *GrovePi) ReadFromSensor(s sensors.Sensor) (sensors.Measurement, error) {
	var m sensors.Measurement

	switch s.Mode {
		case "analog":
			m = grovePi.analogRead(s.Pin)
			break
		case "dht":
			m = grovePi.readDHT(s.Pin)
			break
		default
			panic("Unsupported sensor mode: %s", s.Mode)
	}
	
	return m
}
