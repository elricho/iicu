package cmd

import (
	"fmt"
	"net/url"

	"github.com/elricho/iicu/api"
	"github.com/spf13/cobra"
)

var wellnessCmd = &cobra.Command{
	Use:   "wellness",
	Short: "Wellness and health data",
}

var wellnessGetCmd = &cobra.Command{
	Use:   "get",
	Short: "Get wellness data for a date or date range",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		date, _ := cmd.Flags().GetString("date")
		oldest, _ := cmd.Flags().GetString("oldest")
		newest, _ := cmd.Flags().GetString("newest")

		if date != "" {
			d, err := api.ParseDate(date)
			if err != nil {
				return err
			}
			data, err := client.Get(client.AthletePath(fmt.Sprintf("/wellness/%s", d)))
			if err != nil {
				return err
			}
			outputJSON(data)
			return nil
		}

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

		path := client.AthletePath("/wellness")
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

var wellnessUpdateCmd = &cobra.Command{
	Use:   "update <date>",
	Short: "Update wellness entry for a date",
	Long:  "Update wellness entry. Pass JSON body via stdin. Date format: YYYY-MM-DD.",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		d, err := api.ParseDate(args[0])
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/wellness/%s", d)), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	wellnessGetCmd.Flags().String("date", "", "specific date (YYYY-MM-DD, 'today', etc.)")
	wellnessGetCmd.Flags().String("oldest", "", "oldest date for range")
	wellnessGetCmd.Flags().String("newest", "", "newest date for range")

	wellnessCmd.AddCommand(wellnessGetCmd)
	wellnessCmd.AddCommand(wellnessUpdateCmd)
	rootCmd.AddCommand(wellnessCmd)
}
