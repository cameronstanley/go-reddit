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
