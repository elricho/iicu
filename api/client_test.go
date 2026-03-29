package api

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClient_Get_SetsBasicAuth(t *testing.T) {
	var gotUser, gotPass string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotUser, gotPass, _ = r.BasicAuth()
		w.WriteHeader(200)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "my-api-key", "i123")
	_, err := c.Get("/test")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotUser != "API_KEY" {
		t.Errorf("expected user 'API_KEY', got %q", gotUser)
	}
	if gotPass != "my-api-key" {
		t.Errorf("expected pass 'my-api-key', got %q", gotPass)
	}
}

func TestClient_Get_ReturnsBody(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"name": "Test"})
	}))
	defer server.Close()

	c := NewClient(server.URL, "key", "i123")
	body, err := c.Get("/athlete/i123")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	var result map[string]string
	if err := json.Unmarshal(body, &result); err != nil {
		t.Fatalf("failed to unmarshal: %v", err)
	}
	if result["name"] != "Test" {
		t.Errorf("expected name 'Test', got %q", result["name"])
	}
}

func TestClient_Get_ReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
		w.Write([]byte(`{"message":"not found"}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "key", "i123")
	_, err := c.Get("/nope")
	if err == nil {
		t.Fatal("expected error")
	}
	apiErr, ok := err.(*APIError)
	if !ok {
		t.Fatalf("expected *APIError, got %T", err)
	}
	if apiErr.StatusCode != 404 {
		t.Errorf("expected status 404, got %d", apiErr.StatusCode)
	}
}

func TestClient_Post_SendsJSON(t *testing.T) {
	var gotContentType string
	var gotBody map[string]string
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotContentType = r.Header.Get("Content-Type")
		json.NewDecoder(r.Body).Decode(&gotBody)
		w.WriteHeader(200)
		w.Write([]byte(`{}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "key", "i123")
	_, err := c.Post("/test", map[string]string{"name": "Workout"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotContentType != "application/json" {
		t.Errorf("expected content-type 'application/json', got %q", gotContentType)
	}
	if gotBody["name"] != "Workout" {
		t.Errorf("expected name 'Workout', got %q", gotBody["name"])
	}
}

func TestClient_AthletePath(t *testing.T) {
	c := NewClient("https://intervals.icu/api/v1", "key", "i99999")
	path := c.AthletePath("/activities")
	if path != "/athlete/i99999/activities" {
		t.Errorf("expected '/athlete/i99999/activities', got %q", path)
	}
}
