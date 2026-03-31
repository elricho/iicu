package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sharedEventsCmd = &cobra.Command{
	Use:   "shared-events",
	Short: "Shared/public events (races, group workouts)",
}

var sharedEventsGetCmd = &cobra.Command{
	Use:   "get <id>",
	Short: "Get a shared event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/shared-event/%s", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sharedEventsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a shared event",
	Long:  "Create a shared event. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw("/shared-event", body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sharedEventsUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update a shared event",
	Long:  "Update a shared event. Pass JSON body via stdin.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(fmt.Sprintf("/shared-event/%s", args[0]), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sharedEventsDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete a shared event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(fmt.Sprintf("/shared-event/%s", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	sharedEventsCmd.AddCommand(sharedEventsGetCmd)
	sharedEventsCmd.AddCommand(sharedEventsCreateCmd)
	sharedEventsCmd.AddCommand(sharedEventsUpdateCmd)
	sharedEventsCmd.AddCommand(sharedEventsDeleteCmd)
	rootCmd.AddCommand(sharedEventsCmd)
}
