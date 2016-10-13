package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"io/ioutil"
	"testing"
)

func TestGetDefaultSubreddits(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	response, _ := ioutil.ReadFile("test_data/default_subreddits.json")

	url := fmt.Sprintf("%s/subreddits/default.json", baseUrl)
	httpmock.RegisterResponder("GET", url, httpmock.NewStringResponder(200, string(response)))

	client := new(Client)
	client.GetDefaultSubreddits()
}
