package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
	"github.com/stretchr/objx"
)

const (
	grovePiAddress = 0x04
)

func main() {
	// Load sensors
	var config map[string]interface{}
	config, err := objx.FromURLQuery("./sensors.config.json")
	if err != nil {
		panic("Sensors configuration file not found. Aborting.")
	}

	spew.Dump(config)

	// Init GrovePi on address grovePiAddress
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic("Impossibile to communicate with the GrovePi")
	}
	defer g.Close()

	// Create readers
	for conf := range config["input"] {
		sensor := sensors.NewSensor(conf["name"], conf["description"], conf["pin"], conf["mode"])
	}

	// Create writers
	// for conf := range config["output"] {

	// }
}
