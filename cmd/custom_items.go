package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var customItemsCmd = &cobra.Command{
	Use:   "custom-items",
	Short: "Custom fields and charts",
}

var customItemsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List custom items",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/custom-item"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var customItemsGetCmd = &cobra.Command{
	Use:   "get <item-id>",
	Short: "Get a custom item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath(fmt.Sprintf("/custom-item/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var customItemsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a custom item",
	Long:  "Create a custom item. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/custom-item"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var customItemsUpdateCmd = &cobra.Command{
	Use:   "update <item-id>",
	Short: "Update a custom item",
	Long:  "Update a custom item. Pass JSON body via stdin.",
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
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/custom-item/%s", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var customItemsDeleteCmd = &cobra.Command{
	Use:   "delete <item-id>",
	Short: "Delete a custom item",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(client.AthletePath(fmt.Sprintf("/custom-item/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var customItemsReorderCmd = &cobra.Command{
	Use:   "reorder",
	Short: "Re-order custom items",
	Long:  "Re-order custom items. Pass JSON array via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(client.AthletePath("/custom-item-indexes"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	customItemsCmd.AddCommand(customItemsListCmd)
	customItemsCmd.AddCommand(customItemsGetCmd)
	customItemsCmd.AddCommand(customItemsCreateCmd)
	customItemsCmd.AddCommand(customItemsUpdateCmd)
	customItemsCmd.AddCommand(customItemsDeleteCmd)
	customItemsCmd.AddCommand(customItemsReorderCmd)
	rootCmd.AddCommand(customItemsCmd)
}
