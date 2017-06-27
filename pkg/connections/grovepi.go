package connections

import (
	"time"

	"github.com/mrmorphic/hwio"
)

// Commands format
const (
	CommandDigitalRead  = 1
	CommandDigitalWrite = 2
	CommandAnalogRead   = 3
	CommandAnalogWrite  = 4
	CommandDHTRead      = 40
)

// put check to see if available pin is in range

// GrovePi struct is used for handling the connection with board
type GrovePi struct {
	i2cmodule hwio.I2CModule
	i2cDevice hwio.I2CDevice
}

// NewGrovePi initializes the GrovePi
func NewGrovePi(address int) (*GrovePi, error) {
	g := new(GrovePi)
	m, err := hwio.GetModule("i2c")
	if err != nil {
		return nil, err
	}

	g.i2cmodule = m.(hwio.I2CModule)
	err = g.i2cmodule.Enable()
	if err != nil {
		return nil, err
	}

	g.i2cDevice = g.i2cmodule.GetDevice(address)
	return g, nil
}

// Close closes the connection with the GrovePi
func (g *GrovePi) Close() {
	g.i2cmodule.Disable()
	hwio.CloseAll()
}

// Read from [pin] in [mode] bytes [size]
func (g *GrovePi) Read(pin byte, mode string, size int) ([]byte, error) {
	var raw []byte
	var err error

	switch mode {
	case "digital":
		raw, err = g.digitalRead(pin, size)
		break
	case "analog":
		raw, err = g.analogRead(pin, size)
		break
	case "dht":
		raw, err = g.dhtRead(pin, size)
		break
	default:
		break
	}

	return raw, err
}

// analogRead reads analogically to the GrovePi
func (g *GrovePi) analogRead(pin byte, size int) ([]byte, error) {
	b := []byte{CommandAnalogRead, pin, 0, 0}
	err := g.i2cDevice.Write(1, b)
	if err != nil {
		return nil, err
	}

	time.Sleep(100 * time.Millisecond)

	g.i2cDevice.ReadByte(1)

	raw, err := g.i2cDevice.Read(1, size)
	if err != nil {
		return nil, err
	}

	return raw, err
}

// digitalRead reads digitally to the GrovePi
func (g *GrovePi) digitalRead(pin byte, size int) ([]byte, error) {
	b := []byte{CommandDigitalRead, pin, 0, 0}
	err := g.i2cDevice.Write(1, b)
	if err != nil {
		return nil, err
	}
	time.Sleep(100 * time.Millisecond)

	// TODO set size via parameter, it's better
	return g.i2cDevice.Read(1, size)
}

// dhtRead returns temperature and humidity from DHT sensor
func (g *GrovePi) dhtRead(pin byte, size int) ([]byte, error) {
	cmd := []byte{CommandDHTRead, pin, 0, 0}

	// prepare and read raw data
	err := g.i2cDevice.Write(1, cmd)
	if err != nil {
		return nil, err
	}
	time.Sleep(600 * time.Millisecond)
	g.i2cDevice.ReadByte(1)
	time.Sleep(100 * time.Millisecond)

	rawdata, err := g.i2cDevice.Read(1, size)
	if err != nil {
		return nil, err
	}

	return rawdata, nil
}

// digitalWrite writes digitally to the GrovePi
func (g *GrovePi) digitalWrite(pin byte, val byte) error {
	b := []byte{CommandDigitalWrite, pin, val, 0}
	err := g.i2cDevice.Write(1, b)
	time.Sleep(100 * time.Millisecond)
	return err
}
