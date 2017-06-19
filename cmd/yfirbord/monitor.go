package main

import (
	"fmt"
	"log"
	"time"

	"github.com/LuCavallin/yfirbord-grovepi/pkg/grovepi"
)

func main() {
	// Load GrovePIconfiguration first
	grovePiConfig := grovepi.Config{Address: 0x04, Pins: nil, Commands: nil}
	g, err := grovepi.Init(grovePiConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	// TODO Load sensors

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
