package sensors

import (
	"errors"
	"fmt"
	"strings"
)

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	Pin         string
	Command     string
	Connection  *Connection
}

// Sensors is to easily handle a list of sensors
type Sensors []Sensor

// InputSensor provides an interface for reading from all sensors
type InputSensor interface {
	Measure() (Measurement, error)
}

// func sensorSFROMCONFIG() Sensors
// 	var sensors Sensors
// 	freach config
// 		sensors[] = factoryFromType(config.type) // light dht sound
// 	=
// 	=

// 	return sensors

// load() {
//  registerformat("dht", func(Sensor) return DHT{Sensor})
//  registerformat("light", func(Sensor) return Light{Sensor})
//  sensorsFromConfig()
//  }

func LoadFromConfig(configFile string) ([]Sensors, error) {
	// parse json and return
}

func CreateSensor(conf map[string]string) (Sensor, error) {

	// Query configuration for datastore defaulting to "memory".
	sensorType := conf.Get("sensorType")

	sensorFactory, ok := datastoreFactories[sensorType]
	if !ok {
		// Factory has not been registered.
		// Make a list of all available datastore factories for logging.
		availableDatastores := make([]string, len(datastoreFactories))
		for k, _ := range datastoreFactories {
			availableDatastores = append(availableDatastores, k)
		}
		return nil, errors.New(fmt.Sprintf("Invalid Datastore name. Must be one of: %s", strings.Join(availableDatastores, ", ")))
	}

	// Run the factory with the configuration.
	return engineFactory(conf)
}

type SensorFactory func(conf map[string]string) (Sensor, error)

var sensorFactories = make(map[string]SensorFactory)

func register(name string, factory SensorFactory) {
	if factory == nil {
		log.Panicf("Sensor factory %s does not exist.", name)
	}
	_, registered := sensorFactories[name]
	if registered {
		log.Errorf("Sensor factory %s already registered. Ignoring.", name)
	}
	sensorFactories[name] = factory
}

func init() {
	Register("postgres", NewPostgreSQLDataStore)
	Register("memory", NewMemoryDataStore)
}
