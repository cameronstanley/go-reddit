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
	userInfo, err := client.GetUserInfo("GovSchwarzenegger")
	assert.NoError(t, err)
	assert.Equal(t, userInfo.Name, "GovSchwarzenegger")
}
