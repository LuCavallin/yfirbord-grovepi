package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/davecgh/go-spew/spew"
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
		fmt.Printf("Sensors configuration file not found.\nERROR: %s \n", err)
	}
	err = json.Unmarshal(configJSON, &sensorConfig)
	if err != nil {
		panic("Sensors configuration invalid format. Aborting.")
		fmt.Printf("Sensors configuration file not found. Aborting. ERROR: %s \n", err)
		sensor, g
	}

	// Init GrovePi on address
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		panic(err)
	}
	defer g.Close()

	// Create readers
	var readersList []readers.Reader
	for _, sensor := range sensorConfig {
		reader, err := readers.NewReader(sensor, g)
		if err != nil {
			panic("Reader creation failed")
		}
		readersList = append(readersList, reader)
	}

	spew.Dump(readersList)

	// Create writers
	// for conf := range config["output"] {

	// }
}
