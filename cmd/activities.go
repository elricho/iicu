package cmd

import (
	"fmt"
	"net/url"

	"github.com/elricho/iicu/api"
	"github.com/spf13/cobra"
)

var activitiesCmd = &cobra.Command{
	Use:   "activities",
	Short: "Manage activities",
}

var activitiesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List recent activities",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		oldest, _ := cmd.Flags().GetString("oldest")
		newest, _ := cmd.Flags().GetString("newest")

		params := url.Values{}
		if oldest != "" {
			d, err := api.ParseDate(oldest)
			if err != nil {
				return err
			}
			params.Set("oldest", d)
		}
		if newest != "" {
			d, err := api.ParseDate(newest)
			if err != nil {
				return err
			}
			params.Set("newest", d)
		}

		path := client.AthletePath("/activities")
		if len(params) > 0 {
			path += "?" + params.Encode()
		}
		data, err := client.Get(path)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesGetCmd = &cobra.Command{
	Use:   "get <activity-id>",
	Short: "Get activity details",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/activity/%s", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesSearchCmd = &cobra.Command{
	Use:   "search <query>",
	Short: "Search activities by name or tag",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		params := url.Values{"q": {args[0]}}
		data, err := client.Get(client.AthletePath("/activities/search") + "?" + params.Encode())
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesUpdateCmd = &cobra.Command{
	Use:   "update <activity-id>",
	Short: "Update activity fields",
	Long:  "Update activity fields. Pass JSON body via stdin.",
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
		data, err := client.PutRaw(fmt.Sprintf("/activity/%s", args[0]), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesDeleteCmd = &cobra.Command{
	Use:   "delete <activity-id>",
	Short: "Delete an activity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(fmt.Sprintf("/activity/%s", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesStreamsCmd = &cobra.Command{
	Use:   "streams <activity-id>",
	Short: "Get activity data streams (power, HR, cadence, etc.)",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		types, _ := cmd.Flags().GetString("types")
		path := fmt.Sprintf("/activity/%s/streams", args[0])
		if types != "" {
			path += "?" + url.Values{"types": {types}}.Encode()
		}
		data, err := client.Get(path)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesIntervalsCmd = &cobra.Command{
	Use:   "intervals <activity-id>",
	Short: "Get activity intervals/laps",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(fmt.Sprintf("/activity/%s/intervals", args[0]))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var activitiesEffortsCmd = &cobra.Command{
	Use:   "efforts <activity-id>",
	Short: "Get best efforts for an activity",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		stream, _ := cmd.Flags().GetString("stream")
		params := url.Values{}
		if stream != "" {
			params.Set("stream", stream)
		}
		path := fmt.Sprintf("/activity/%s/best-efforts", args[0])
		if len(params) > 0 {
			path += "?" + params.Encode()
		}
		data, err := client.Get(path)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	activitiesListCmd.Flags().String("oldest", "", "oldest date (YYYY-MM-DD, 'today', '-7d', etc.)")
	activitiesListCmd.Flags().String("newest", "", "newest date")

	activitiesStreamsCmd.Flags().String("types", "", "stream types (comma-separated: watts,heartrate,cadence,etc.)")

	activitiesEffortsCmd.Flags().String("stream", "watts", "stream to find best efforts for")

	activitiesCmd.AddCommand(activitiesListCmd)
	activitiesCmd.AddCommand(activitiesGetCmd)
	activitiesCmd.AddCommand(activitiesSearchCmd)
	activitiesCmd.AddCommand(activitiesUpdateCmd)
	activitiesCmd.AddCommand(activitiesDeleteCmd)
	activitiesCmd.AddCommand(activitiesStreamsCmd)
	activitiesCmd.AddCommand(activitiesIntervalsCmd)
	activitiesCmd.AddCommand(activitiesEffortsCmd)
	rootCmd.AddCommand(activitiesCmd)
}
