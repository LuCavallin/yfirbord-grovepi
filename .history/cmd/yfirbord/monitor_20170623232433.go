package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
)

const (
	grovePiAddress = 0x04
)

func main() {
	// Load sensors
	var config interface{}
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		panic("Sensors configuration file not found. Aborting.")
	}
	if err := json.Unmarshal(configJSON, &config); err != nil {
		panic("Sensors configuration file is not a valid JSON.")
	}

	spew.Dump(config)

	// Init GrovePi on address grovePiAddress
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic("Impossibile to communicate with the GrovePo")
	}
	defer g.Close()

	// Create readers
	// for sensorConf := range config {
	// 	sensor := sensors.NewSensor()
	// }
}
