package main

import (
	"fmt"
	"time"

	"github.com/LuCavallin/yfirbord-grovepi/grovepi"
)

func main() {
	var g grovepi.GrovePi
	g = *grovepi.InitGrovePi(0x04)
	err := g.PinMode(grovepi.D2, "input")
	if err != nil {
		fmt.Println(err)
	}
	for {
		tmp, hum, err := g.ReadDHT(grovepi.D2)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Temperature: %f - Humidity: %f\n", tmp, hum)
		time.Sleep(500 * time.Millisecond)
	}
}
