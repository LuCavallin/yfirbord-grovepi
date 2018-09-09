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
	"github.com/spf13/cobra"
	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/drivers/i2c"
	"gobot.io/x/gobot/platforms/raspi"
	"time"
)

// senseCmd represents the sense command
var senseCmd = &cobra.Command{
	Use:   "sense",
	Short: "Start gathering data from sensors",
	Run: func(cmd *cobra.Command, args []string) {
		sense()
	},
}

var interval int

func init() {
	rootCmd.AddCommand(senseCmd)
	rootCmd.Flags().IntVarP(&interval, "interval", "i", 5000, "Polling interval in milliseconds")
}

func sense() {
	r := raspi.NewAdaptor()
	gp := i2c.NewGrovePiDriver(r)
	//mqttAdaptor := mqtt.NewAdaptorWithAuth();
	//mqttClient := mqtt.NewDriver(mqttAdaptor, "to")

	led := gpio.NewLedDriver(gp, "D3")

	work := func() {
		gobot.Every(1*time.Second, func() {
			led.Toggle()
		})
	}

	robot := gobot.NewRobot("hytta",
		[]gobot.Connection{r},
		[]gobot.Device{gp, led},
		work,
	)

	robot.Start()
}
