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

	client := Client{}
  articles, err := client.GetHotArticles("news")
  assert.NoError(t, err)
  assert.Equal(t, len(articles), 3)
}

func TestGetNewArticles(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/new.json", baseUrl)
	mockResponseFromFile(url, "test_data/article/new_articles.json")
	defer httpmock.DeactivateAndReset()

	client := Client{}
  articles, err := client.GetNewArticles("news")
  assert.NoError(t, err)
  assert.Equal(t, len(articles), 3)
}

func TestGetTopArticles(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/top.json", baseUrl)
	mockResponseFromFile(url, "test_data/article/top_articles.json")
	defer httpmock.DeactivateAndReset()

	client := Client{}
  articles, err := client.GetTopArticles("news")
  assert.NoError(t, err)
  assert.Equal(t, len(articles), 3)
}
