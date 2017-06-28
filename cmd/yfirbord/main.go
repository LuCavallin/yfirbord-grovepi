package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/io"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors/parsers"
)

const (
	grovePiAddress = 0x04
	secondsWait    = 360
)

// config should have pinNumber -> sensors, in a way that sensors are mapped to pins
var loadedSensors map[string][]sensors.Sensor

func main() {
	// Load sensors
	config, err := ioutil.ReadFile("./sensors.json")
	if err != nil {
		fmt.Printf("Sensors configuration file not found.\nERROR: %s \n", err)
		os.Exit(1)
	}
	if json.Unmarshal(config, &loadedSensors); err != nil {
		fmt.Printf("Sensors configuration is invalid.\nERROR: %s \n", err)
		os.Exit(1)
	}

	// Init GrovePi on address
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		fmt.Printf("Couldn't create connection with device.\nERROR: %s \n", err)
		os.Exit(1)
	}
	defer g.Close()

	// Create readers
	reader := io.NewReader(parsers.GetParsers(), g)

	// Create channel
	c := make(chan parsers.Measurement)

	// Start reading!
	for _, sensor := range loadedSensors["input"] {
		// Go read the sensor every 5 minutes
		go func(r *io.Reader, s sensors.Sensor, c chan<- parsers.Measurement) {
			for {
				measurement, err := r.Read(s)
				if err != nil {
					log.Panicln(err)
				} else {
					c <- measurement
				}
				// Read every secondsWait minutes
				time.Sleep(time.Second * secondsWait)
			}
		}(reader, sensor, c)
	}

	// Create API manager
	// Print out
	go func(c chan parsers.Measurement) {
		measurement := <-c
		for item, measure := range measurement {
			fmt.Printf("%s: %v\n", item, measure)
		}
	}(c)

	time.Sleep(time.Minute)
}
