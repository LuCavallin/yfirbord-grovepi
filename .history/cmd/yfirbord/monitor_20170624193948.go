package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
)

const (
	grovePiAddress = 0x04
)

var config []sensors.Sensor

func main() {
	// Load sensors
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configJSON, &config)
	if err != nil {
		panic(err)
	}
	spew.Dump(config)

	// Init GrovePi on address
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic("Impossibile to communicate with the GrovePi")
	}
	defer g.Close()

	// Create readers
	// readersConfig := config.Get("input")
	// for conf := objx.ArrayEach() {
	// 	sensor := sensors.NewSensor(conf["name"], conf["description"], conf["pin"], conf["mode"])
	// }

	// Create writers
	// for conf := range config["output"] {

	// }
}
