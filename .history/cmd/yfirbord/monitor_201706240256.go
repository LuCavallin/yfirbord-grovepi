package main

import (
	"io/ioutil"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
	"github.com/stretchr/objx"
	"github.com/buger/jsonparser"
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
	config, err := jsonparser.Get(configJSON, "input")
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
	// for conf := config. {
	// 	sensor := sensors.NewSensor(conf["name"], conf["description"], conf["pin"], conf["mode"])
	// }

	// Create writers
	// for conf := range config["output"] {

	// }
}
