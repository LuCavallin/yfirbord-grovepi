package main

import (
	"fmt"
	"time"

	"github.com/LuCavallin/yfirbord-grovepi/grovepi"
)

func main() {
	var g grovepi.GrovePi
	g = *grovepi.InitGrovePi(0x04)
	err := g.PinMode(grovepi.D2, "output")
	if err != nil {
		fmt.Println(err)
	}
	for {
		g.ReadDHT(grovepi.D2)
		time.Sleep(500 * time.Millisecond)
	}
}
