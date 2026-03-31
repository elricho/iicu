package cmd

import (
	"net/url"

	"github.com/spf13/cobra"
)

var curvesCmd = &cobra.Command{
	Use:   "curves",
	Short: "Performance curves (power, HR, pace)",
}

var curvesPowerCmd = &cobra.Command{
	Use:   "power",
	Short: "Power duration curves",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		sportType, _ := cmd.Flags().GetString("type")
		params := url.Values{}
		if sportType != "" {
			params.Set("type", sportType)
		}
		path := client.AthletePath("/power-curves")
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

var curvesHRCmd = &cobra.Command{
	Use:   "hr",
	Short: "Heart rate curves",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		sportType, _ := cmd.Flags().GetString("type")
		params := url.Values{}
		if sportType != "" {
			params.Set("type", sportType)
		}
		path := client.AthletePath("/hr-curves")
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

var curvesPaceCmd = &cobra.Command{
	Use:   "pace",
	Short: "Pace curves",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		sportType, _ := cmd.Flags().GetString("type")
		gap, _ := cmd.Flags().GetBool("gap")
		params := url.Values{}
		if sportType != "" {
			params.Set("type", sportType)
		}
		if gap {
			params.Set("gap", "true")
		}
		path := client.AthletePath("/pace-curves")
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
	curvesPowerCmd.Flags().String("type", "", "sport type (Ride, Run, etc.)")
	curvesHRCmd.Flags().String("type", "", "sport type (Ride, Run, etc.)")
	curvesPaceCmd.Flags().String("type", "", "sport type (Ride, Run, etc.)")
	curvesPaceCmd.Flags().Bool("gap", false, "use gradient adjusted pace")

	curvesCmd.AddCommand(curvesPowerCmd)
	curvesCmd.AddCommand(curvesHRCmd)
	curvesCmd.AddCommand(curvesPaceCmd)
	rootCmd.AddCommand(curvesCmd)
}
