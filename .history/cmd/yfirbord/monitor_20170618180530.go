package main

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/LuCavallin/yfirbord-grovepi/pkg/grovepi"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	g, err := grovepi.Init(0x04)
	if err != nil {
		log.Fatal(err)
	}
	defer g.Close()

	err = g.PinMode(grovepi.D4, grovepi.OutputPin)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// DHT
		t, h, err := g.ReadDHT(grovepi.D2)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("T: %f - H: %f \n", t, h)
		time.Sleep(500 * time.Millisecond)

		// g.DigitalWrite(grovepi.D4, 1)
		rawLight, err := g.AnalogRead(grovepi.A2)
		if err != nil {
			log.Fatal(err)
		}

		spew.Dump(rawLight)
		lightOutput := math.Float32frombits(lightBits)
		fmt.Printf("L: %f \n", lightOutput)
		time.Sleep(500 * time.Millisecond)
		// g.DigitalWrite(grovepi.D4, 0)
		// time.Sleep(500 * time.Millisecond)
	}
}
