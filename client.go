package go_threema_bcast

import "net/http"

// BaseURL is the default base URL for the Threema Broadcast API
const BaseURL = "https://broadcast.threema.ch/api/v1"

type Client struct {
	HTTPClient *http.Client
	BaseURL    string
}

// NewClient constructs a client using http.DefaultClient and the default
// base URL. The returned client is ready for use.
func NewClient() *Client {
	return &Client{
		HTTPClient: http.DefaultClient,
		BaseURL:    BaseURL,
	}
}
