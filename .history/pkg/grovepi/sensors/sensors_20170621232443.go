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
	Connection  *Connection
}

// InputSensor provides an interface for reading from all sensors
type InputSensor interface {
	Measure() (Measurement, error)
}

// SensorFactory returns a function to generate a Sensor
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
	register("dht", DHT)
	register("light", Light)
}
