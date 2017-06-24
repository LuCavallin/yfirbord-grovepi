package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/readers"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

const (
	grovePiAddress = 0x04
)

var sensorConfig []sensors.Sensor

func main() {
	// Load sensors
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		panic("Sensors configuration file not found. Aborting.")
	}
	err = json.Unmarshal(configJSON, &sensorConfig)
	if err != nil {
		panic("Sensors configuration invalid format. Aborting.")
	}

	// Init GrovePi on address
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic("Impossibile to communicate with the GrovePi")
	}
	defer g.Close()

	// Create readers
	var readersList []readers.Reader
	for _, sensor := range sensorConfig {
		readersList = append(readersList, readers.NewReader(sensor, g))
	}

	// Create writers
	// for conf := range config["output"] {

	// }
}
