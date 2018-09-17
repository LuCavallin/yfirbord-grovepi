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
	"github.com/lucavallin/hytta/pkg/message"
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

var interval int

func init() {
	senseCmd.Flags().IntVarP(&interval, "interval", "i", 5000, "Polling interval in milliseconds")
	rootCmd.AddCommand(senseCmd)
}

func sense() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)

	// Find way to handle configuration
	mqttAdaptor := mqtt.NewAdaptorWithAuth("xxx", "hytta", "xxx", "xxx")

	// Find way to generate sensors in a abstract manner
	sound := aio.NewGroveSoundSensorDriver(gp, "A0")
	temperature := aio.NewGroveTemperatureSensorDriver(gp, "A1")
	light := aio.NewGroveLightSensorDriver(gp, "A2")

	work := func() {
		gobot.Every(10 * time.Second, func() {
			lightVal, _ := light.Read()
			l, _ := json.Marshal(message.NewReading(light.Name(), lightVal))
			mqttAdaptor.Publish("from", l)

			soundVal, _ := sound.Read()
			s, _ := json.Marshal(message.NewReading(sound.Name(), soundVal))
			mqttAdaptor.Publish("from", s)

			temperatureVal, _ := temperature.Read()
			t, _ := json.Marshal(message.NewReading(temperature.Name(), temperatureVal))
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
