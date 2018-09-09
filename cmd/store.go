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
	"fmt"

	"github.com/spf13/cobra"
)

// storeCmd represents the store command
var storeCmd = &cobra.Command{
	Use:   "store",
	Short: "Connects to MQTT and stores incoming data [SERVER-SIDE ONLY!]",
	Long: `This command is supposed to be used on the server side to store the data coming into the MQTT broker`,
	Run: func(cmd *cobra.Command, args []string) {
		store()
	},
}

func init() {
	rootCmd.AddCommand(storeCmd)
}

func store() {
	fmt.Println("Storing data...")
}
