package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sync"

	"github.com/lucavallin/hytta-grovepi/pkg/connections"
	"github.com/lucavallin/hytta-grovepi/pkg/io"
	"github.com/lucavallin/hytta-grovepi/pkg/sensors"
	"github.com/lucavallin/hytta-grovepi/pkg/sensors/parsers"
)

const (
	grovePiAddress = 0x04
)

// config should have pinNumber -> sensors, in a way that sensors are mapped to pins
var loadedSensors map[string][]sensors.Sensor

func main() {
	// Init concurrency stuff
	m := &sync.Mutex{}
	var wg sync.WaitGroup

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
	wg.Add(len(loadedSensors["input"]))
	for _, sensor := range loadedSensors["input"] {
		go func(reader *io.Reader, sensor sensors.Sensor, c chan<- parsers.Measurement) {
			defer wg.Done()
			m.Lock()
			measurement, err := reader.Read(sensor)
			m.Unlock()

			if err != nil {
				log.Panicln(err)
			} else {
				c <- measurement
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

	// Wait then quit
	wg.Wait()
	close(ch)
}
