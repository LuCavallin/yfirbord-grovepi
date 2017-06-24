package main

import (
	"io/ioutil"

	"github.com/buger/jsonparser"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
)

const (
	grovePiAddress = 0x04
)

func main() {
	// Load sensors
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		panic("Sensors configuration file not found. Aborting.")
	}
	config, _, _, err := jsonparser.Get(configJSON, "input")
	if err != nil {
		panic("Sensors configuration invalid format. Aborting.")
	}

	// Init GrovePi on address grovePiAddress
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic("Impossibile to communicate with the GrovePi")
	}
	defer g.Close()

	// Create readers
	for conf := config. {
		sensor := sensors.NewSensor(conf["name"], conf["description"], conf["pin"], conf["mode"])
	}

	// Create writers
	// for conf := range config["output"] {

	// }
}
