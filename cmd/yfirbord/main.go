package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
	"github.com/lucavallin/yfirbord-grovepi/pkg/readers"
	"github.com/lucavallin/yfirbord-grovepi/pkg/sensors"
	"github.com/lucavallin/yfirbord-grovepi/pkg/writers"
)

const (
	grovePiAddress = 0x04
	minutesWait    = 5
)

// config should have pinNumber -> sensors, in a way that sensors are mapped to pins
var sensorConfig map[string][]sensors.Sensor

func main() {
	// Load sensors
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		fmt.Printf("Sensors configuration file not found.\nERROR: %s \n", err)
	}
	err = json.Unmarshal(configJSON, &sensorConfig)
	if err != nil {
		fmt.Printf("Sensors configuration is invalid.\nERROR: %s \n", err)
	}

	// Init GrovePi on address
	g, err := connections.NewGrovePi(grovePiAddress)
	if err != nil {
		fmt.Printf("Couldn't create connection with device.\nERROR: %s \n", err)
	}
	defer g.Close()

	// Create readers
	var readersList []readers.Reader
	for _, sensor := range sensorConfig["input"] {
		reader, err := readers.NewReader(sensor, g)
		if err != nil {
			fmt.Printf("Couldn't establish reading connection with device.\nERROR: %s \n", err)
		}
		readersList = append(readersList, reader)
	}

	// Create writers
	var writersList []writers.Writer
	for _, sensor := range sensorConfig["output"] {
		writer, err := writers.NewWriter(sensor, g)
		if err != nil {
			fmt.Printf("Couldn't establish writing connection with device.\nERROR: %s \n", err)
		}
		writersList = append(writersList, writer)
	}

	// Create API manager

	// run
	spew.Dump(readersList)
	c := make(chan sensors.Measurement)
	for _, reader := range readersList {
		// Go read the sensor every 5 minutes
		go func(r readers.Reader, c chan sensors.Measurement) {
			r.Read(c)
			time.Sleep(time.Minute * minutesWait)
		}(reader, c)
	}
}
