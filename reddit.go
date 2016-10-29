// Package reddit provides Reddit API wrapper utilities.
package reddit

import (
	"net/http"
)

const (
	baseAuthURL = "https://oauth.reddit.com"
	baseURL     = "http://reddit.com"
)

// Client is the client for interacting with the Reddit API.
type Client struct {
	http      *http.Client
	userAgent string
}

// NoAuthClient is the unauthenticated client for interacting with the Reddit API.
var NoAuthClient = &Client{
	http: new(http.Client),
}
