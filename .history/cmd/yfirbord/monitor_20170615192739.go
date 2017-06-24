package main

import (
	"fmt"
	"log"
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
		// g.DigitalWrite(grovepi.D4, 1)
		t, h, err := g.ReadDHT(grovepi.D3)
		spew.Dump(t, h)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("T: %f - H: %f \n", t, h)
		time.Sleep(500 * time.Millisecond)
		// g.DigitalWrite(grovepi.D4, 0)
		// time.Sleep(500 * time.Millisecond)
	}
}
