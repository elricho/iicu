package cmd

import (
	"github.com/spf13/cobra"
)

var athleteCmd = &cobra.Command{
	Use:   "athlete",
	Short: "Athlete profile and fitness data",
}

var athleteProfileCmd = &cobra.Command{
	Use:   "profile",
	Short: "Get athlete profile",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath(""))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var athleteFitnessCmd = &cobra.Command{
	Use:   "fitness",
	Short: "Get fitness summary (CTL, ATL, TSB)",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/athlete-summary"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	athleteCmd.AddCommand(athleteProfileCmd)
	athleteCmd.AddCommand(athleteFitnessCmd)
	rootCmd.AddCommand(athleteCmd)
}
