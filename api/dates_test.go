package api

import (
	"testing"
	"time"
)

func TestParseDate_ISO(t *testing.T) {
	result, err := ParseDate("2024-06-15")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != "2024-06-15" {
		t.Errorf("expected '2024-06-15', got %q", result)
	}
}

func TestParseDate_Today(t *testing.T) {
	result, err := ParseDate("today")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := time.Now().Format("2006-01-02")
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestParseDate_Yesterday(t *testing.T) {
	result, err := ParseDate("yesterday")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := time.Now().AddDate(0, 0, -1).Format("2006-01-02")
	if result != expected {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestParseDate_RelativeDays(t *testing.T) {
	tests := []struct {
		input   string
		daysAgo int
	}{
		{"-7d", 7},
		{"-30d", 30},
		{"-1d", 1},
	}
	for _, tt := range tests {
		result, err := ParseDate(tt.input)
		if err != nil {
			t.Fatalf("ParseDate(%q): unexpected error: %v", tt.input, err)
		}
		expected := time.Now().AddDate(0, 0, -tt.daysAgo).Format("2006-01-02")
		if result != expected {
			t.Errorf("ParseDate(%q) = %q, want %q", tt.input, result, expected)
		}
	}
}

func TestParseDate_Invalid(t *testing.T) {
	_, err := ParseDate("not-a-date")
	if err == nil {
		t.Error("expected error for invalid date")
	}
}
