package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoadConfig_CreatesDefaultWhenMissing(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")

	cfg, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.DefaultProfile != "" {
		t.Errorf("expected empty default profile, got %q", cfg.DefaultProfile)
	}
	if cfg.Profiles == nil {
		t.Error("expected non-nil profiles map")
	}
}

func TestLoadConfig_ReadsExistingFile(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")

	content := []byte(`default_profile: me
profiles:
  me:
    api_key: "key123"
    athlete_id: "i99999"
`)
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatal(err)
	}

	cfg, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if cfg.DefaultProfile != "me" {
		t.Errorf("expected default_profile 'me', got %q", cfg.DefaultProfile)
	}
	p, ok := cfg.Profiles["me"]
	if !ok {
		t.Fatal("expected profile 'me' to exist")
	}
	if p.APIKey != "key123" {
		t.Errorf("expected api_key 'key123', got %q", p.APIKey)
	}
	if p.AthleteID != "i99999" {
		t.Errorf("expected athlete_id 'i99999', got %q", p.AthleteID)
	}
}

func TestAddProfile(t *testing.T) {
	cfg := &Config{
		Profiles: make(map[string]Profile),
	}
	cfg.AddProfile("alice", Profile{APIKey: "k1", AthleteID: "i1"})

	p, ok := cfg.Profiles["alice"]
	if !ok {
		t.Fatal("expected profile 'alice'")
	}
	if p.APIKey != "k1" {
		t.Errorf("expected api_key 'k1', got %q", p.APIKey)
	}
}

func TestRemoveProfile(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "alice",
		Profiles: map[string]Profile{
			"alice": {APIKey: "k1", AthleteID: "i1"},
			"bob":   {APIKey: "k2", AthleteID: "i2"},
		},
	}
	err := cfg.RemoveProfile("alice")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if _, ok := cfg.Profiles["alice"]; ok {
		t.Error("expected 'alice' to be removed")
	}
	if cfg.DefaultProfile != "" {
		t.Errorf("expected default profile cleared, got %q", cfg.DefaultProfile)
	}
}

func TestRemoveProfile_NotFound(t *testing.T) {
	cfg := &Config{Profiles: make(map[string]Profile)}
	err := cfg.RemoveProfile("nope")
	if err == nil {
		t.Error("expected error for missing profile")
	}
}

func TestResolveCredentials_FlagOverride(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "me",
		Profiles: map[string]Profile{
			"me": {APIKey: "config-key", AthleteID: "config-id"},
		},
	}
	apiKey, athleteID, err := cfg.ResolveCredentials("", "flag-key", "flag-id")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if apiKey != "flag-key" {
		t.Errorf("expected 'flag-key', got %q", apiKey)
	}
	if athleteID != "flag-id" {
		t.Errorf("expected 'flag-id', got %q", athleteID)
	}
}

func TestResolveCredentials_EnvOverride(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "me",
		Profiles: map[string]Profile{
			"me": {APIKey: "config-key", AthleteID: "config-id"},
		},
	}
	t.Setenv("IICU_API_KEY", "env-key")
	t.Setenv("IICU_ATHLETE_ID", "env-id")

	apiKey, athleteID, err := cfg.ResolveCredentials("", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if apiKey != "env-key" {
		t.Errorf("expected 'env-key', got %q", apiKey)
	}
	if athleteID != "env-id" {
		t.Errorf("expected 'env-id', got %q", athleteID)
	}
}

func TestResolveCredentials_ProfileSelection(t *testing.T) {
	cfg := &Config{
		DefaultProfile: "me",
		Profiles: map[string]Profile{
			"me":   {APIKey: "my-key", AthleteID: "my-id"},
			"jane": {APIKey: "jane-key", AthleteID: "jane-id"},
		},
	}
	apiKey, athleteID, err := cfg.ResolveCredentials("jane", "", "")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if apiKey != "jane-key" {
		t.Errorf("expected 'jane-key', got %q", apiKey)
	}
	if athleteID != "jane-id" {
		t.Errorf("expected 'jane-id', got %q", athleteID)
	}
}

func TestResolveCredentials_NoConfig(t *testing.T) {
	cfg := &Config{Profiles: make(map[string]Profile)}
	_, _, err := cfg.ResolveCredentials("", "", "")
	if err == nil {
		t.Error("expected error when no credentials available")
	}
}

func TestSave(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "config.yaml")

	cfg := &Config{
		path:           path,
		DefaultProfile: "me",
		Profiles: map[string]Profile{
			"me": {APIKey: "k1", AthleteID: "i1"},
		},
	}
	if err := cfg.Save(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	loaded, err := LoadFromPath(path)
	if err != nil {
		t.Fatalf("unexpected error loading: %v", err)
	}
	if loaded.DefaultProfile != "me" {
		t.Errorf("expected 'me', got %q", loaded.DefaultProfile)
	}
	p := loaded.Profiles["me"]
	if p.APIKey != "k1" || p.AthleteID != "i1" {
		t.Errorf("unexpected profile values: %+v", p)
	}
}
