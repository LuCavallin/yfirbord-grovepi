package grovepi

// Config holds configuration for the GrovePi
type Config struct {
	Address  int
	Pins     map[int]string
	Commands map[string]int
}

// FromJSON Loads configuration from a JSON file
func FromJSON(path string) Config {
	return Config{}
}
