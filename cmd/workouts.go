package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var workoutsCmd = &cobra.Command{
	Use:   "workouts",
	Short: "Workout library management",
}

var workoutsFoldersCmd = &cobra.Command{
	Use:   "folders",
	Short: "List workout library folders and plans",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/folders"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsListCmd = &cobra.Command{
	Use:   "list",
	Short: "List workouts in library",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/workouts"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a workout",
	Long:  "Create a workout in the library. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/workouts"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsUpdateCmd = &cobra.Command{
	Use:   "update <workout-id>",
	Short: "Update a workout",
	Long:  "Update a workout. Pass JSON body via stdin.",
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
		data, err := client.PutRaw(client.AthletePath(fmt.Sprintf("/workouts/%s", args[0])), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsDeleteCmd = &cobra.Command{
	Use:   "delete <workout-id>",
	Short: "Delete a workout",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Delete(client.AthletePath(fmt.Sprintf("/workouts/%s", args[0])))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsBulkCreateCmd = &cobra.Command{
	Use:   "bulk-create",
	Short: "Create multiple workouts",
	Long:  "Create multiple workouts. Pass JSON array via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/workouts/bulk"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsDuplicateCmd = &cobra.Command{
	Use:   "duplicate",
	Short: "Duplicate workouts on a plan",
	Long:  "Duplicate workouts. Pass JSON body via stdin.",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		body, err := readStdin()
		if err != nil {
			return err
		}
		data, err := client.PostRaw(client.AthletePath("/duplicate-workouts"), body)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsTagsCmd = &cobra.Command{
	Use:   "tags",
	Short: "List all workout tags",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Get(client.AthletePath("/workout-tags"))
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

var workoutsApplyPlanCmd = &cobra.Command{
	Use:   "apply-plan",
	Short: "Apply plan changes to calendar",
	RunE: func(cmd *cobra.Command, args []string) error {
		client, err := newClient()
		if err != nil {
			return err
		}
		data, err := client.Put(client.AthletePath("/apply-plan-changes"), nil)
		if err != nil {
			return err
		}
		outputJSON(data)
		return nil
	},
}

func init() {
	workoutsCmd.AddCommand(workoutsFoldersCmd)
	workoutsCmd.AddCommand(workoutsListCmd)
	workoutsCmd.AddCommand(workoutsCreateCmd)
	workoutsCmd.AddCommand(workoutsUpdateCmd)
	workoutsCmd.AddCommand(workoutsDeleteCmd)
	workoutsCmd.AddCommand(workoutsBulkCreateCmd)
	workoutsCmd.AddCommand(workoutsDuplicateCmd)
	workoutsCmd.AddCommand(workoutsTagsCmd)
	workoutsCmd.AddCommand(workoutsApplyPlanCmd)
	rootCmd.AddCommand(workoutsCmd)
}
