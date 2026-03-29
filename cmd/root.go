// cmd/root.go
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/elricho/iicu/api"
	"github.com/elricho/iicu/config"
	"github.com/spf13/cobra"
)

var (
	flagProfile   string
	flagAPIKey    string
	flagAthleteID string
	flagHuman     bool

	cfg *config.Config
)

var rootCmd = &cobra.Command{
	Use:   "iicu",
	Short: "CLI for intervals.icu",
	Long:  "Command line tools for the intervals.icu training platform API.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&flagProfile, "profile", "", "config profile to use")
	rootCmd.PersistentFlags().StringVar(&flagAPIKey, "api-key", "", "API key (overrides config)")
	rootCmd.PersistentFlags().StringVar(&flagAthleteID, "athlete-id", "", "athlete ID (overrides config)")
	rootCmd.PersistentFlags().BoolVar(&flagHuman, "human", false, "human-readable output")
}

func initConfig() {
	var err error
	cfg, err = config.Load()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Warning: could not load config: %v\n", err)
		cfg = &config.Config{Profiles: make(map[string]config.Profile)}
	}
}

func newClient() (*api.Client, error) {
	apiKey, athleteID, err := cfg.ResolveCredentials(flagProfile, flagAPIKey, flagAthleteID)
	if err != nil {
		return nil, err
	}
	return api.NewClient(api.DefaultBaseURL, apiKey, athleteID), nil
}

func outputJSON(data []byte) {
	if flagHuman {
		var indented bytes.Buffer
		if err := json.Indent(&indented, data, "", "  "); err != nil {
			os.Stdout.Write(data)
			fmt.Fprintln(os.Stdout)
			return
		}
		fmt.Fprintln(os.Stdout, indented.String())
		return
	}
	os.Stdout.Write(data)
	fmt.Fprintln(os.Stdout)
}

func readStdin() ([]byte, error) {
	return io.ReadAll(os.Stdin)
}
