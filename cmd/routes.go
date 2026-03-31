package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var routesCmd = &cobra.Command{
	Use:   "routes",
	Short: "Manage routes",
}

var routesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List routes with activity counts",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/routes"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var routesGetCmd = &cobra.Command{
	Use:   "get <route-id>",
	Short: "Get a route",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath(fmt.Sprintf("/routes/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var routesUpdateCmd = &cobra.Command{
	Use:   "update <route-id>",
	Short: "Update a route",
	Long:  "Update a route. Pass JSON body via stdin.",
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
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/routes/%s", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var routesCompareCmd = &cobra.Command{
	Use:   "compare <route-id> <other-route-id>",
	Short: "Compare route similarity",
	Args:  cobra.ExactArgs(2),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath(fmt.Sprintf("/routes/%s/similarity/%s", args[0], args[1])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	routesCmd.AddCommand(routesListCmd)
	routesCmd.AddCommand(routesGetCmd)
	routesCmd.AddCommand(routesUpdateCmd)
	routesCmd.AddCommand(routesCompareCmd)
	rootCmd.AddCommand(routesCmd)
}
