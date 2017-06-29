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

	err = json.Unmarshal(config, &loadedSensors)
	if err != nil {
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
	ch := make(chan parsers.Measurement)

	// Start reading!
	for _, sensor := range loadedSensors["input"] {
		go func(reader *io.Reader, sensor sensors.Sensor, c chan<- parsers.Measurement) {
			for {
				measurement, err := reader.Read(sensor)

				if err != nil {
					log.Panicln(err)
				} else {
					c <- measurement
				}

				time.Sleep(time.Second)
			}
		}(reader, sensor, ch)
	}

	// Create API manager
	go func(c <-chan parsers.Measurement) {
		for {
			measurement := <-c
			for item, measure := range measurement {
				fmt.Printf("%s: %v\n", item, measure)
			}
		}
	}(ch)

	// How do I quit when the api is done?
	time.Sleep(time.Minute)
	close(ch)
}
