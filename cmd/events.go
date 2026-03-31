package cmd

import (
	"fmt"
	"net/url"

	"github.com/elricho/iicu/api"
	"github.com/spf13/cobra"
)

var eventsCmd = &cobra.Command{
	Use:   "events",
	Short: "Manage calendar events and planned workouts",
}

var eventsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List calendar events",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		oldest, _ := cmd.Flags().GetString("oldest")
		newest, _ := cmd.Flags().GetString("newest")
		category, _ := cmd.Flags().GetStringSlice("category")

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
		for _, c := range category {
			params.Add("category", c)
		}

		path := client.AthletePath("/events")
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

var eventsGetCmd = &cobra.Command{
	Use:   "get <event-id>",
	Short: "Get an event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath(fmt.Sprintf("/events/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var eventsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a calendar event",
	Long:  "Create a calendar event. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/events"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var eventsUpdateCmd = &cobra.Command{
	Use:   "update <event-id>",
	Short: "Update an event",
	Long:  "Update an event. Pass JSON body via stdin.",
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
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/events/%s", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var eventsDeleteCmd = &cobra.Command{
	Use:   "delete <event-id>",
	Short: "Delete an event",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(client.AthletePath(fmt.Sprintf("/events/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var eventsBulkCreateCmd = &cobra.Command{
	Use:   "bulk-create",
	Short: "Create multiple events",
	Long:  "Create multiple events. Pass JSON array via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/events/bulk"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var eventsBulkDeleteCmd = &cobra.Command{
	Use:   "bulk-delete",
	Short: "Delete multiple events by ID",
	Long:  "Delete events by ID. Pass JSON array of {id} or {external_id} via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(client.AthletePath("/events/bulk-delete"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var eventsDuplicateCmd = &cobra.Command{
	Use:   "duplicate",
	Short: "Duplicate events on calendar",
	Long:  "Duplicate events. Pass JSON body with event IDs and target dates via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/duplicate-events"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	eventsListCmd.Flags().String("oldest", "", "oldest date")
	eventsListCmd.Flags().String("newest", "", "newest date")
	eventsListCmd.Flags().StringSlice("category", nil, "filter by category (WORKOUT, RACE_A, NOTE, etc.)")

	eventsCmd.AddCommand(eventsListCmd)
	eventsCmd.AddCommand(eventsGetCmd)
	eventsCmd.AddCommand(eventsCreateCmd)
	eventsCmd.AddCommand(eventsUpdateCmd)
	eventsCmd.AddCommand(eventsDeleteCmd)
	eventsCmd.AddCommand(eventsBulkCreateCmd)
	eventsCmd.AddCommand(eventsBulkDeleteCmd)
	eventsCmd.AddCommand(eventsDuplicateCmd)
	rootCmd.AddCommand(eventsCmd)
}
