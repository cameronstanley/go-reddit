package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserInfo(t *testing.T) {
	url := fmt.Sprintf("%s/user/GovSchwarzenegger/about.json", baseUrl)
	mockResponseFromFile(url, "test_data/user/user_info.json")
	defer httpmock.DeactivateAndReset()

	client := NoAuthClient
	user, err := client.GetUserInfo()
	assert.NoError(t, err)
  assert.Equal(t, user.Name, "GovSchwarzenegger")
}
