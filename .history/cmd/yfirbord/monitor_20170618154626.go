package main

import (
	"encoding/binary"
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

	err = g.PinMode(grovepi.D2, grovepi.OutputPin)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// g.DigitalWrite(grovepi.D4, 1)
		l, err := g.DigitalRead(grovepi.D2)
		spew.Dump(l)
		cac := binary.LittleEndian.Uint32(l[:])
		// read sensors concurrently
		if err != nil {
			log.Fatal(err)
		}

		// chan to communicate between reding sensor and sending api

		fmt.Printf("L: %f \n", math.Float32frombits(cac))
		time.Sleep(500 * time.Millisecond)
		// g.DigitalWrite(grovepi.D4, 0)
		// time.Sleep(500 * time.Millisecond)
	}
}
