package grovepi

// Config holds configuration for the GrovePi
type Config struct {
	address  int
	pins     map[int]string
	commands map[string]int
}

// FromJSON Loads configuration from a JSON file
func FromJSON(path string) Config {
	return Config{}
}
