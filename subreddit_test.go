package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"testing"
)

func mockResponseFromFile(url string, filepath string) {
	httpmock.Activate()
	response, _ := ioutil.ReadFile(filepath)
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(response)))
}

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
}

func checkSliceSize(slice []*Subreddit, expectedSize int, t *testing.T) {
	if len(slice) != expectedSize {
		t.Error(fmt.Sprintf("Expected %d elements, got ", expectedSize), len(slice))
	}
}

func TestGetDefaultSubreddits(t *testing.T) {
	url := fmt.Sprintf("%s/subreddits/default.json", baseUrl)
	mockResponseFromFile(url, "test_data/subreddit/default_subreddits.json")
	defer httpmock.DeactivateAndReset()

	client := new(Client)
	subreddits, err := client.GetDefaultSubreddits()
	checkError(err, t)
	checkSliceSize(subreddits, 3, t)
}

func TestGetGoldSubreddits(t *testing.T) {
	url := fmt.Sprintf("%s/subreddits/gold.json", baseUrl)
	mockResponseFromFile(url, "test_data/subreddit/gold_subreddits.json")
	defer httpmock.DeactivateAndReset()

	client := new(Client)
	subreddits, err := client.GetGoldSubreddits()
	checkError(err, t)
	checkSliceSize(subreddits, 0, t)
}

func TestGetNewSubreddits(t *testing.T) {
	url := fmt.Sprintf("%s/subreddits/new.json", baseUrl)
	mockResponseFromFile(url, "test_data/subreddit/new_subreddits.json")
	defer httpmock.DeactivateAndReset()

	client := new(Client)
	subreddits, err := client.GetNewSubreddits()
	checkError(err, t)
	checkSliceSize(subreddits, 3, t)
}

func TestGetPopularSubreddits(t *testing.T) {
	url := fmt.Sprintf("%s/subreddits/popular.json", baseUrl)
	mockResponseFromFile(url, "test_data/subreddit/popular_subreddits.json")
	defer httpmock.DeactivateAndReset()

	client := new(Client)
	subreddits, err := client.GetPopularSubreddits()
	checkError(err, t)
	checkSliceSize(subreddits, 3, t)
}
