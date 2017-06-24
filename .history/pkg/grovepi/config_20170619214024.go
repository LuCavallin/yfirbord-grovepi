package grovepi

// Config holds configuration for the GrovePi
type Config struct {
	pins     map[int]string
	commands map[string]int
}

// Loads configuration from a JSON file
func ConfigFromJSON(path string) Config {
	return Config{}
}
