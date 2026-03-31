package cmd

import (
	"github.com/spf13/cobra"
)

var trainingPlanCmd = &cobra.Command{
	Use:   "training-plan",
	Short: "Athlete training plan assignment",
}

var trainingPlanGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get athlete's training plan",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/training-plan"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var trainingPlanSetCmd = &cobra.Command{
	Use:   "set",
	Short: "Change athlete's training plan",
	Long:  "Change training plan. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(client.AthletePath("/training-plan"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	trainingPlanCmd.AddCommand(trainingPlanGetCmd)
	trainingPlanCmd.AddCommand(trainingPlanSetCmd)
	rootCmd.AddCommand(trainingPlanCmd)
}
