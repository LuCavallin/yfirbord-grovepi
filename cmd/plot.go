package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var sensor string

func init() {
	plotCmd.Flags().StringVarP(&sensor, "sensor", "s", "", "Sensor to display data for")
	rootCmd.AddCommand(plotCmd)
}

var plotCmd = &cobra.Command{
	Use:   "plot",
	Short: "Display data as graph",
	Run: func(cmd *cobra.Command, args []string) {
		plot()
	},
}

func plot() {
	fmt.Print("Displaying data...")
}
