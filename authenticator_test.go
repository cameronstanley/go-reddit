package reddit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetAuthenticationUrl(t *testing.T) {
	authenticator := NewAuthenticator("client_id", "client_secret", "123456789abcdef", "identity")
	authenticationUrl := authenticator.GetAuthenticationUrl()
	assert.Equal(t, authenticationUrl, "https://www.reddit.com/api/v1/authorize?client_id=client_id&response_type=code&scope=identity&state=123456789abcdef")
}
