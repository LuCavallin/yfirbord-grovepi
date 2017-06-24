package grovepi

// Config holds configuration for the GrovePi
type Config struct {
	pins     map[int]string
	commands map[string]int
}

func ConfigFromJSON(path string) Config {
	return Config{}
}
