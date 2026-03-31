package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/elricho/iicu/config"
	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage iicu configuration",
}

var configInitCmd = &cobra.Command{
	Use:   "init",
	Short: "Set up iicu configuration",
	RunE: func(cmd *cobra.Command, args []string) error {
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("Profile name [default]: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = "default"
		}

		fmt.Print("API key (from https://intervals.icu/settings): ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		if apiKey == "" {
			return fmt.Errorf("API key is required")
		}

		fmt.Print("Athlete ID (e.g. i12345, from your profile URL): ")
		athleteID, _ := reader.ReadString('\n')
		athleteID = strings.TrimSpace(athleteID)
		if athleteID == "" {
			return fmt.Errorf("athlete ID is required")
		}

		cfg.AddProfile(name, config.Profile{
			APIKey:    apiKey,
			AthleteID: athleteID,
		})
		if cfg.DefaultProfile == "" {
			cfg.DefaultProfile = name
		}

		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}

		fmt.Printf("Profile %q saved as default. Config at %s\n", name, config.DefaultPath())
		return nil
	},
}

var configProfilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "List configured profiles",
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(cfg.Profiles) == 0 {
			fmt.Println("No profiles configured. Run 'iicu config init' to set up.")
			return nil
		}
		for name, p := range cfg.Profiles {
			marker := "  "
			if name == cfg.DefaultProfile {
				marker = "* "
			}
			fmt.Printf("%s%s (athlete: %s)\n", marker, name, p.AthleteID)
		}
		return nil
	},
}

var configUseCmd = &cobra.Command{
	Use:   "use <profile>",
	Short: "Switch default profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if _, ok := cfg.Profiles[name]; !ok {
			return fmt.Errorf("profile %q not found", name)
		}
		cfg.DefaultProfile = name
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}
		fmt.Printf("Default profile set to %q\n", name)
		return nil
	},
}

var configAddCmd = &cobra.Command{
	Use:   "add <name>",
	Short: "Add a new profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		reader := bufio.NewReader(os.Stdin)

		fmt.Print("API key: ")
		apiKey, _ := reader.ReadString('\n')
		apiKey = strings.TrimSpace(apiKey)
		if apiKey == "" {
			return fmt.Errorf("API key is required")
		}

		fmt.Print("Athlete ID: ")
		athleteID, _ := reader.ReadString('\n')
		athleteID = strings.TrimSpace(athleteID)
		if athleteID == "" {
			return fmt.Errorf("athlete ID is required")
		}

		cfg.AddProfile(name, config.Profile{
			APIKey:    apiKey,
			AthleteID: athleteID,
		})
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}
		fmt.Printf("Profile %q added\n", name)
		return nil
	},
}

var configRemoveCmd = &cobra.Command{
	Use:   "remove <name>",
	Short: "Remove a profile",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		name := args[0]
		if err := cfg.RemoveProfile(name); err != nil {
			return err
		}
		if err := cfg.Save(); err != nil {
			return fmt.Errorf("saving config: %w", err)
		}
		fmt.Printf("Profile %q removed\n", name)
		return nil
	},
}

func init() {
	configCmd.AddCommand(configInitCmd)
	configCmd.AddCommand(configProfilesCmd)
	configCmd.AddCommand(configUseCmd)
	configCmd.AddCommand(configAddCmd)
	configCmd.AddCommand(configRemoveCmd)
	rootCmd.AddCommand(configCmd)
}
