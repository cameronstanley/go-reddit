package reddit

import (
	"golang.org/x/oauth2"
)

type Authenticator struct {
	config *oauth2.Config
	state  string
}

const (
	ScopeIdentity        = "identity"
	ScopeEdit            = "edit"
	ScopeFlair           = "flair"
	ScopeHistory         = "history"
	ScopeModConfig       = "modconfig"
	ScopeModFlair        = "modflair"
	ScopeModLog          = "modlog"
	ScopeModWiki         = "modwiki"
	ScopeMySubreddits    = "mysubreddits"
	ScopePrivateMessages = "privatemessages"
	ScopeRead            = "read"
	ScopeReport          = "report"
	ScopeSave            = "save"
	ScopeSubmit          = "submit"
	ScopeSubscribe       = "subscribe"
	ScopeVote            = "vote"
	ScopeWikiEdit        = "wikiedit"
	ScopeWikiRead        = "wikiread"
)

func NewAuthenticator(clientID string, clientSecret string, state string, scopes ...string) *Authenticator {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://www.reddit.com/api/v1/authorize",
			TokenURL: "https://www.reddit.com/api/v1/access_token",
		},
	}

	return &Authenticator{config: config, state: state}
}

func (a *Authenticator) GetAuthenticationUrl() string {
	return a.config.AuthCodeURL(a.state)
}
