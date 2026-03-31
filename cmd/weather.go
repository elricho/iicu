package cmd

import (
	"github.com/spf13/cobra"
)

var weatherCmd = &cobra.Command{
	Use:   "weather",
	Short: "Weather forecast data",
}

var weatherConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Get weather forecast configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/weather-config"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var weatherUpdateConfigCmd = &cobra.Command{
	Use:   "update-config",
	Short: "Update weather forecast configuration",
	Long:  "Update weather config. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(client.AthletePath("/weather-config"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var weatherForecastCmd = &cobra.Command{
	Use:   "forecast",
	Short: "Get weather forecast",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/weather-forecast"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	weatherCmd.AddCommand(weatherConfigCmd)
	weatherCmd.AddCommand(weatherUpdateConfigCmd)
	weatherCmd.AddCommand(weatherForecastCmd)
	rootCmd.AddCommand(weatherCmd)
}
