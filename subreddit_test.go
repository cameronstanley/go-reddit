package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

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
