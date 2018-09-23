// Copyright © 2018 Luca Cavallin <me@lucavall.in>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/lucavallin/hytta/pkg/message"
	"github.com/lucavallin/hytta/pkg/messages"
	"github.com/spf13/cobra"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/aio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/mqtt"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

// senseCmd represents the sense command
var senseCmd = &cobra.Command{
	Use:   "sense",
	Short: "DEVICE: start gathering data from sensors",
	Run: func(cmd *cobra.Command, args []string) {
		sense()
	},
}

var interval string

func init() {
	senseCmd.Flags().StringVarP(&interval, "interval", "i", "5000ms", "Polling time interval")
	rootCmd.AddCommand(senseCmd)
}

func sense() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)

	// Configuration management needs to be improved
	mqttAdaptor := mqtt.NewAdaptorWithAuth("xxx", "hytta", "xxx", "xxx")
	sound := aio.NewGroveSoundSensorDriver(gp, "A0")
	temperature := aio.NewGroveTemperatureSensorDriver(gp, "A1")
	light := aio.NewGroveLightSensorDriver(gp, "A2")

	timeInterval, err := time.ParseDuration(interval)
	if err != nil {
		panic(err)
	}

	work := func() {
		gobot.Every(timeInterval, func() {

			// This could well be abstracted to remove code duplication, but since I have no plans to
			// have more sensors, this is good enough for now
			lightVal, _ := light.Read()
			l, e := json.Marshal(messages.NewReading(light.Name(), lightVal))
			fmt.Println(lightVal, l, e)
			mqttAdaptor.Publish("from", l)

			soundVal, _ := sound.Read()
			s, e := json.Marshal(messages.NewReading(sound.Name(), soundVal))
			fmt.Println(soundVal, s, e)
			mqttAdaptor.Publish("from", s)

			temperatureVal, _ := temperature.Read()
			t, e := json.Marshal(messages.NewReading(temperature.Name()	, temperatureVal))
			fmt.Println(temperatureVal, t, e)
			mqttAdaptor.Publish("from", t)
		})
	}

	robot := gobot.NewRobot("hytta",
		[]gobot.Connection{r, mqttAdaptor},
		[]gobot.Device{gp, light, sound, temperature},
		work,
	)

	robot.Start()
}
