package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const DefaultBaseURL = "https://intervals.icu/api/v1"

type APIError struct {
	StatusCode int
	Message    string
	Body       string
}

func (e *APIError) Error() string {
	if e.Message != "" {
		return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Message)
	}
	return fmt.Sprintf("API error %d: %s", e.StatusCode, e.Body)
}

type Client struct {
	BaseURL    string
	APIKey     string
	AthleteID  string
	HTTPClient *http.Client
}

func NewClient(baseURL, apiKey, athleteID string) *Client {
	return &Client{
		BaseURL:    baseURL,
		APIKey:     apiKey,
		AthleteID:  athleteID,
		HTTPClient: &http.Client{},
	}
}

func (c *Client) AthletePath(suffix string) string {
	return fmt.Sprintf("/athlete/%s%s", c.AthleteID, suffix)
}

func (c *Client) doRequest(method, path string, body io.Reader) ([]byte, error) {
	url := c.BaseURL + path

	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, fmt.Errorf("creating request: %w", err)
	}

	req.SetBasicAuth("API_KEY", c.APIKey)
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("executing request: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("reading response: %w", err)
	}

	if resp.StatusCode >= 400 {
		apiErr := &APIError{
			StatusCode: resp.StatusCode,
			Body:       string(respBody),
		}
		var errResp struct {
			Message string `json:"message"`
		}
		if json.Unmarshal(respBody, &errResp) == nil && errResp.Message != "" {
			apiErr.Message = errResp.Message
		}
		return nil, apiErr
	}

	return respBody, nil
}

func (c *Client) Get(path string) ([]byte, error) {
	return c.doRequest(http.MethodGet, path, nil)
}

func (c *Client) Post(path string, payload interface{}) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshaling request body: %w", err)
	}
	return c.doRequest(http.MethodPost, path, bytes.NewReader(data))
}

func (c *Client) Put(path string, payload interface{}) ([]byte, error) {
	data, err := json.Marshal(payload)
	if err != nil {
		return nil, fmt.Errorf("marshaling request body: %w", err)
	}
	return c.doRequest(http.MethodPut, path, bytes.NewReader(data))
}

func (c *Client) Delete(path string) ([]byte, error) {
	return c.doRequest(http.MethodDelete, path, nil)
}

func (c *Client) PutRaw(path string, body []byte) ([]byte, error) {
	return c.doRequest(http.MethodPut, path, bytes.NewReader(body))
}

func (c *Client) PostRaw(path string, body []byte) ([]byte, error) {
	return c.doRequest(http.MethodPost, path, bytes.NewReader(body))
}
