package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var sportsCmd = &cobra.Command{
	Use:   "sports",
	Short: "Sport settings (zones, thresholds)",
}

var sportsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List sport settings",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/sport-settings"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sportsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create sport settings",
	Long:  "Create sport settings with defaults. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/sport-settings"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sportsUpdateCmd = &cobra.Command{
	Use:   "update <id>",
	Short: "Update sport settings",
	Long:  "Update sport settings. Pass JSON body via stdin. ID can be numeric ID or sport type (e.g. 'Run').",
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
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/sport-settings/%s", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sportsDeleteCmd = &cobra.Command{
	Use:   "delete <id>",
	Short: "Delete sport settings",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(client.AthletePath(fmt.Sprintf("/sport-settings/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var sportsApplyCmd = &cobra.Command{
	Use:   "apply <id>",
	Short: "Apply sport settings to matching activities",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Put(client.AthletePath(fmt.Sprintf("/sport-settings/%s/apply", args[0])), nil)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	sportsCmd.AddCommand(sportsListCmd)
	sportsCmd.AddCommand(sportsCreateCmd)
	sportsCmd.AddCommand(sportsUpdateCmd)
	sportsCmd.AddCommand(sportsDeleteCmd)
	sportsCmd.AddCommand(sportsApplyCmd)
	rootCmd.AddCommand(sportsCmd)
}
