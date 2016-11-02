package reddit

import (
	"fmt"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeleteComment(t *testing.T) {
	url := fmt.Sprintf("%s/api/del", baseAuthURL)
	httpmock.Activate()
	httpmock.RegisterResponder("POST", url, httpmock.NewStringResponder(200, "{}"))
	defer httpmock.DeactivateAndReset()

	client := NoAuthClient
	err := client.DeleteComment("t1_d9hthja")
	assert.NoError(t, err)
}
