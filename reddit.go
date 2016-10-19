package reddit

import (
	"net/http"
)

type Client struct {
	http *http.Client
}

const baseUrl = "http://reddit.com"

var NoAuthClient = &Client{
	http: new(http.Client),
}
