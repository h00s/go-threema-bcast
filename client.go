package go_threema_bcast

import (
  "fmt"
  "io"
  "net/http"
)

// BaseURL is the default base URL for the Threema Broadcast API
const BaseURL = "https://broadcast.threema.ch/api/v1"

type Client struct {
  HTTPClient *http.Client
  BaseURL    string
  APIKey     string
}

type Links struct {
  Links []Link `json:"_links"`
}

type Link struct {
  Ref  string `json:"ref"`
  Link string `json:"link"`
}

// NewClient constructs a client using http.DefaultClient and the default
// base URL. The returned client is ready for use.
func NewClient(APIKey string) *Client {
  return &Client{
    HTTPClient: http.DefaultClient,
    BaseURL:    BaseURL,
    APIKey:     APIKey,
  }
}

func (c *Client) Get(path string) ([]byte, error) {
  req, err := http.NewRequest("GET", c.BaseURL+path, nil)
  if err != nil {
    return nil, fmt.Errorf("failed to build request: %v", err)
  }
  req.Header.Set("X-Api-Key", c.APIKey)
  req.Header.Set("Content-Type", "application/json")

  resp, err := c.HTTPClient.Do(req)
  if err != nil {
    return nil, fmt.Errorf("failed to do request: %v", err)
  }
  defer resp.Body.Close()

  if resp.StatusCode != 200 {
    return nil, fmt.Errorf("returned status code: %v", resp.StatusCode)
  }

  body, err := io.ReadAll(resp.Body)
  if err != nil {
    return nil, fmt.Errorf("failed to read body: %v", err)
  }

  return body, nil
}
