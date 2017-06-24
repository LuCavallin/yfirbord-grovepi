package grovepi

// Config holds configuration for the GrovePi
type Config struct {
	Address  int
	Pins     map[int]string
	Commands map[string]int
	Sensors  Sensors
}

// FromJSON Loads configuration from a JSON file
// TODO: we could abstract the config provider
func FromJSON(path string) Config {
	return Config{}
}
