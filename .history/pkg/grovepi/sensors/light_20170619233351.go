package sensors

import (
	"fmt"
	"log"
	"time"

	"github.com/lucavallin/yfirbord-grovepi/pkg/grovepi"
)

func (o Sensor) Read() {
	light, err := g.AnalogRead(grovepi.A2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("L: %d \n", light)
	time.Sleep(500 * time.Millisecond)
}
