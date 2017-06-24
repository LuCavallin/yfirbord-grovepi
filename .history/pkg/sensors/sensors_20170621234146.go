package sensors

import (
	"fmt"
	"strings"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	pin         string
	cmd         string
	conn        *grovepi.GrovePi
}

// InputSensor provides an interface for reading from all sensors
type InputSensor interface {
	Measure() (Measurement, error)
}

// SensorFactory returns a function to generate a Sensor
type SensorFactory func(conf map[string]string) (Sensor, error)

var sensorFactories = make(map[string]SensorFactory)

func init() {
	register("dht", newDHT)
	register("light", newLight)
}

func register(name string, factory SensorFactory) {
	if factory == nil {
		// log.Panicf("Sensor factory %s does not exist.", name)
		panic("Sensor factor does not exist.")

	}
	_, registered := sensorFactories[name]
	if registered {
		error("Sensor factory already registered. Ignoring.")
	}
	sensorFactories[name] = factory
}

// CreateSensor from configuration
func CreateSensor(conf map[string]string) (Sensor, error) {

	// Query configuration for datastore defaulting to "memory".
	sensorType := conf.Get("type")

	sensorFactory, found := sensorFactories[sensorType]
	if !found {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableSensors := make([]string, len(sensorFactories))
		for f := range sensorFactories {
			availableSensors = append(availableSensors, f)
		}
		return nil, fmt.Errorf("Invalid Datastore name. Must be one of: %s", strings.Join(sensorFactories, ", "))
	}

	// Run the factory with the configuration.
	return sensorFactory(conf)
}

// newDHT creates a new DHT sensor
func newDHT(conf map[string]string) (Sensor, error) {
	return &DHT{
		Name:        conf.name,
		Description: conf.description,
		pin:         conf.pin,
		cmd:         conf.command,
		conn:        &grovepi.GrovePi,
	}, nil
}

// newDHT creates a new DHT sensor
func newLight(conf map[string]string) (Sensor, error) {
	return &Light{
		Name:        conf.name,
		Description: conf.description,
		pin:         conf.pin,
		cmd:         conf.command,
		conn:        &grovepi.GrovePi,
	}, nil
}
