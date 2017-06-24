package main

import (
	"io/ioutil"

	"github.com/buger/jsonparser"
	"github.com/davecgh/go-spew/spew"
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

	spew.Dump(configJSON)

	// Init GrovePi on address grovePiAddress
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic("Impossibile to communicate with the GrovePi")
	}
	defer g.Close()

	// Create readers
	// for conf := range configJSON {
	// 	sensor := sensors.NewSensor(configJSON["name"], configJSON["description"], configJSON["pin"], configJSON["mode"])
	// }

	// Create writers
	// for conf := range config["output"] {

	// }
}
