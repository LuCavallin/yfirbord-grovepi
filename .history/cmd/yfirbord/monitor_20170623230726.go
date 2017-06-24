package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/connections"
)

func main() {
	// Init GrovePi on address 0x04
	g, err := connections.GrovePi.Init(0x04)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// Load sensors
	var config interface{}
	configJSON, err := ioutil.ReadFile("./sensors.config.json")
	if err != nil {
		panic("Sensors configuration file not found. Aborting.")
	}
	err := json.Unmarshal(configJson, &config)

	for {
		// DHT
		t, h, err := g.ReadDHT(grovepi.D2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("T: %f - H: %f \n", t, h)
		time.Sleep(500 * time.Millisecond)

		// Light sensor
		light, err := g.AnalogRead(grovepi.A2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("L: %d \n", light)
		time.Sleep(500 * time.Millisecond)

	}
}
