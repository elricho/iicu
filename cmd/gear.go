package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var gearCmd = &cobra.Command{
	Use:   "gear",
	Short: "Manage gear and components",
}

var gearListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all gear",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/gear"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var gearCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Add new gear",
	Long:  "Create gear. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/gear"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var gearUpdateCmd = &cobra.Command{
	Use:   "update <gear-id>",
	Short: "Update gear",
	Long:  "Update gear. Pass JSON body via stdin.",
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
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/gear/%s", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var gearDeleteCmd = &cobra.Command{
	Use:   "delete <gear-id>",
	Short: "Delete gear",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(client.AthletePath(fmt.Sprintf("/gear/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var gearAddReminderCmd = &cobra.Command{
	Use:   "add-reminder <gear-id>",
	Short: "Add a maintenance reminder",
	Long:  "Add reminder to gear. Pass JSON body via stdin.",
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
		data, err := client.PostRaw(client.AthletePath(fmt.Sprintf("/gear/%s/reminder", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var gearUpdateReminderCmd = &cobra.Command{
	Use:   "update-reminder <gear-id> <reminder-id>",
	Short: "Update a gear reminder",
	Long:  "Update reminder. Pass JSON body via stdin.",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/gear/%s/reminder/%s", args[0], args[1])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	gearCmd.AddCommand(gearListCmd)
	gearCmd.AddCommand(gearCreateCmd)
	gearCmd.AddCommand(gearUpdateCmd)
	gearCmd.AddCommand(gearDeleteCmd)
	gearCmd.AddCommand(gearAddReminderCmd)
	gearCmd.AddCommand(gearUpdateReminderCmd)
	rootCmd.AddCommand(gearCmd)
}
