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
	"github.com/lucavallin/hytta/pkg/hytta"
	"github.com/lucavallin/hytta/pkg/messages"
	"github.com/prometheus/common/log"
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

func init() {
	senseCmd.Flags().StringVarP(&interval, "interval", "i", "5000ms", "Polling time interval")
	rootCmd.AddCommand(senseCmd)
}

var (
	interval string
	// This could be improved with an abstract factory or something similar.
	// Pin configuration should be moved to a e.g. YAML file
	r = raspi.NewAdaptor()
	gp = i2c.NewGrovePiDriver(r)
	sound = aio.NewGroveSoundSensorDriver(gp, "A0")
	temperature = aio.NewGroveTemperatureSensorDriver(gp, "A1")
	light = aio.NewGroveLightSensorDriver(gp, "A2")
)


func sense() {
	timeInterval, err := time.ParseDuration(interval)
	if err != nil {
		log.Fatal(err)
	}

	// MQTT configuration would be better in e.g. a YAML file
	mqttConfig := hytta.NewMqttConfig()
	mqttAdaptor := mqtt.NewAdaptorWithAuth(mqttConfig.Host, mqttConfig.ClientID, mqttConfig.Username, mqttConfig.Password)


	sensors := map[string]hytta.ReadableSensor{
		"sound": sound,
		"temperature": temperature,
		"light": light,
	}
	work := func() {
		gobot.Every(timeInterval, func() {
			for name, sensor := range sensors {
				val, _ := sensor.Read()
				message, _ := json.Marshal(messages.NewReading(name, val, time.Now()))
				mqttAdaptor.Publish("from", message)
			}
		})
	}

	robot := gobot.NewRobot("hytta",
		[]gobot.Connection{r, mqttAdaptor},
		[]gobot.Device{gp, sound, temperature, light},
		work,
	)

	robot.Start()
}
