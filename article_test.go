package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGetHotArticles(t *testing.T) {
	url := fmt.Sprintf("%s/r/news/hot.json", baseUrl)
	mockResponseFromFile(url, "test_data/article/hot_articles.json")
	defer httpmock.DeactivateAndReset()

	client := new(Client)
	client.GetHotArticles("news")
}
