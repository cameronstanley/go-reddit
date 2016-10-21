package reddit

import (
	"golang.org/x/oauth2"
)

func AuthenticationUrl(clientID string, clientSecret string, state string, scopes ...string) string {
	conf := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.reddit.com/api/v1/authorize",
			TokenURL: "https://www.reddit.com/api/v1/access_token",
		},
	}

	return conf.AuthCodeURL(state)
}
