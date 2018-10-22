package messages

import "time"

// Reading is a struct containing sensor data for a new reading
type Reading struct {
	Name string `json:"name"`
	Value interface{} `json:"value"`
	Time time.Time `json:"time"`
}

// NewReading creates a new reading message for transmission of sensors data
func NewReading(name string, value interface{}, timestamp time.Time) *Reading {
	return &Reading{name, value, timestamp}
}