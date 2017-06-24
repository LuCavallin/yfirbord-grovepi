package connections

import (
	"encoding/binary"
	"math"
	"time"

	"github.com/mrmorphic/hwio"
)

// Pins, so we can handle connections in a easier way
var pins = map[string][]*Pin{
	"analog": {
		*Pin{0, false},
		Pin{1, false},
		Pin{2, false},
	},
	"digital": {
		Pin{2, false},
		Pin{3, false},
		Pin{4, false},
		Pin{5, false},
		Pin{6, false},
		Pin{7, false},
		Pin{8, false},
	},
}

// Commands format
const (
	CommandDigitalRead  = 1
	CommandDigitalWrite = 2
	CommandAnalogRead   = 3
	CommandAnalogWrite  = 4
	CommandDHTRead      = 40
)

// GrovePi struct is used for handling the connection with board
type GrovePi struct {
	i2cmodule hwio.I2CModule
	i2cDevice hwio.I2CDevice
}

// Config holds configuration for the GrovePi
type Config struct {
	Address int
	Pins    map[int]string
}

// FromJSON Loads configuration from a JSON file
// TODO: we could abstract the config provider
func FromJSON(path string) Config {
	return Config{}
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
func (grovePi *GrovePi) AnalogRead(pin byte) (int, error) {
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
func (grovePi *GrovePi) DigitalRead(pin byte) ([]byte, error) {
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
func (grovePi *GrovePi) DigitalWrite(pin byte, val byte) error {
	b := []byte{CommandDigitalWrite, pin, val, 0}
	err := grovePi.i2cDevice.Write(1, b)
	time.Sleep(100 * time.Millisecond)
	return err
}

// ReadDHT returns temperature and humidity from DHT sensor
func (grovePi *GrovePi) ReadDHT(pin byte) (float32, float32, error) {
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
