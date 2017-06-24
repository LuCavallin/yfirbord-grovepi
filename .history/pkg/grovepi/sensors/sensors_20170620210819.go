package sensors

// Measurement is a map with the sensor name as index and the measurement as vaue
// e.g. map[temperature] = 32 or map[humidity] = 64
type Measurement map[string]interface{}

// Sensor contains pin and pinMode for the sensor
type Sensor struct {
	Name        string
	Description string
	Pin         string
	Command     string
}

// Sensors is to easily handle a list of sensors
type Sensors []Sensor

// InputSensor provides an interface for reading from all sensors
type InputSensor interface {
	Read() (Measurement, error)
}

var datastoreFactories = make(map[string]DataStoreFactory)

func Register(name string, factory DataStoreFactory) {
	if factory == nil {
		log.Panicf("Datastore factory %s does not exist.", name)
	}
	_, registered := datastoreFactories[name]
	if registered {
		log.Errorf("Datastore factory %s already registered. Ignoring.", name)
	}
	datastoreFactories[name] = factory
}

func init() {
	Register("postgres", NewPostgreSQLDataStore)
	Register("memory", NewMemoryDataStore)
}
