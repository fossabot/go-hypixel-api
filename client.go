package hypixel

import (
	"net/http"
)

type Client struct {
	BaseURL string
	APIKey  string
	HTTP    *http.Client
}

// NewClient creates a new hypixel client
// key is your hypixel api key
//
// https://api.hypixel.net/
func NewClient(key string) *Client {
	return &Client{
		BaseURL: "https://api.hypixel.net/",
		APIKey:  key,
		HTTP:    http.DefaultClient,
	}
}

func (c *Client) GetBaseURL() string {
	return c.BaseURL
}

func (c *Client) GetAPIKey() string {
	return c.APIKey
}

func (c *Client) GetHTTPClient() *http.Client {
	return c.HTTP
}

func (c *Client) SetBaseURL(url string) {
	c.BaseURL = url
}

func (c *Client) SetHTTPClient(client *http.Client) {
	c.HTTP = client
}

func (c *Client) SetAPIKey(key string) {
	c.APIKey = key
}
