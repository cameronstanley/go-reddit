package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHotLinks(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/hot.json", baseURL)
	mockResponseFromFile(url, "test_data/link/hot_links.json")
	defer httpmock.DeactivateAndReset()

	client := NoAuthClient
	links, err := client.GetHotLinks("news")
	assert.NoError(t, err)
	assert.Equal(t, len(links), 3)
}

func TestGetNewLinks(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/new.json", baseURL)
	mockResponseFromFile(url, "test_data/link/new_links.json")
	defer httpmock.DeactivateAndReset()

	client := NoAuthClient
	links, err := client.GetNewLinks("news")
	assert.NoError(t, err)
	assert.Equal(t, len(links), 3)
}

func TestGetTopLinks(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/top.json", baseURL)
	mockResponseFromFile(url, "test_data/link/top_links.json")
	defer httpmock.DeactivateAndReset()

	client := NoAuthClient
	links, err := client.GetTopLinks("news")
	assert.NoError(t, err)
	assert.Equal(t, len(links), 3)
}
