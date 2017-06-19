package sensors

import (
	"fmt"
	"log"
	"time"

	"github.com/lucavall.in/yfirbord-grovepi/pkg/grovepi"
	"github.com/lucavall.in/yfirbord-grovepi/pkg/sensors"
)

func (o sensors.InputSensor) Read() {
	light, err := g.AnalogRead(grovepi.A2)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("L: %d \n", light)
	time.Sleep(500 * time.Millisecond)
}
