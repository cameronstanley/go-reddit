package reddit

import (
  "fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHotArticles(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/hot.json", baseUrl)
	mockResponseFromFile(url, "test_data/article/hot_articles.json")
	defer httpmock.DeactivateAndReset()

	client := new(Client)
  articles, err := client.GetHotArticles("news")
  assert.NoError(t, err)
  assert.Equal(t, len(articles), 3)
}
