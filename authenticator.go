package reddit

import (
	"errors"
	"golang.org/x/oauth2"
)

type Authenticator struct {
	config *oauth2.Config
	state  string
}

const (
	authUrl  = "https://www.reddit.com/api/v1/authorize"
	tokenUrl = "https://www.reddit.com/api/v1/access_token"

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

func NewAuthenticator(clientID string, clientSecret string, redirectUrl string, state string, scopes ...string) *Authenticator {
	config := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Scopes:       scopes,
		Endpoint: oauth2.Endpoint{
			AuthURL:  authUrl,
			TokenURL: tokenUrl,
		},
		RedirectURL: redirectUrl,
	}

	return &Authenticator{config: config, state: state}
}

func (a *Authenticator) GetAuthenticationUrl() string {
	return a.config.AuthCodeURL(a.state)
}

func (a *Authenticator) GetToken(state string, code string) (*oauth2.Token, error) {
	if state != a.state {
		return nil, errors.New("Invalid state")
	}

	return a.config.Exchange(oauth2.NoContext, code)
}

func (a *Authenticator) GetAuthClient(token *oauth2.Token) *Client {
	return &Client{
		http: a.config.Client(oauth2.NoContext, token),
	}
}
