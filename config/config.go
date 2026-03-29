package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Profile struct {
	APIKey    string `yaml:"api_key"`
	AthleteID string `yaml:"athlete_id"`
}

type Config struct {
	path           string
	DefaultProfile string             `yaml:"default_profile"`
	Profiles       map[string]Profile `yaml:"profiles"`
}

func DefaultPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".config", "iicu", "config.yaml")
}

func Load() (*Config, error) {
	return LoadFromPath(DefaultPath())
}

func LoadFromPath(path string) (*Config, error) {
	cfg := &Config{
		path:     path,
		Profiles: make(map[string]Profile),
	}

	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return cfg, nil
		}
		return nil, fmt.Errorf("reading config: %w", err)
	}

	if err := yaml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("parsing config: %w", err)
	}
	cfg.path = path
	if cfg.Profiles == nil {
		cfg.Profiles = make(map[string]Profile)
	}
	return cfg, nil
}

func (c *Config) Save() error {
	dir := filepath.Dir(c.path)
	if err := os.MkdirAll(dir, 0700); err != nil {
		return fmt.Errorf("creating config directory: %w", err)
	}

	data, err := yaml.Marshal(c)
	if err != nil {
		return fmt.Errorf("marshaling config: %w", err)
	}

	if err := os.WriteFile(c.path, data, 0600); err != nil {
		return fmt.Errorf("writing config: %w", err)
	}
	return nil
}

func (c *Config) AddProfile(name string, profile Profile) {
	c.Profiles[name] = profile
}

func (c *Config) RemoveProfile(name string) error {
	if _, ok := c.Profiles[name]; !ok {
		return fmt.Errorf("profile %q not found", name)
	}
	delete(c.Profiles, name)
	if c.DefaultProfile == name {
		c.DefaultProfile = ""
	}
	return nil
}

func (c *Config) ResolveCredentials(profileName, flagAPIKey, flagAthleteID string) (apiKey, athleteID string, err error) {
	apiKey = flagAPIKey
	athleteID = flagAthleteID

	if apiKey == "" {
		apiKey = os.Getenv("IICU_API_KEY")
	}
	if athleteID == "" {
		athleteID = os.Getenv("IICU_ATHLETE_ID")
	}

	if apiKey == "" || athleteID == "" {
		pName := profileName
		if pName == "" {
			pName = c.DefaultProfile
		}
		if pName != "" {
			if p, ok := c.Profiles[pName]; ok {
				if apiKey == "" {
					apiKey = p.APIKey
				}
				if athleteID == "" {
					athleteID = p.AthleteID
				}
			}
		}
	}

	if apiKey == "" || athleteID == "" {
		return "", "", fmt.Errorf("missing credentials: run 'iicu config init' to set up, or provide --api-key and --athlete-id")
	}
	return apiKey, athleteID, nil
}
