package reddit

import (
	"encoding/json"
	"fmt"
)

func (c *Client) GetUserInfo(username string) (*Account, error) {
	url := fmt.Sprintf("%s/user/%s/about.json", baseUrl, username)
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result struct {
		Kind string  `json:"kind"`
		Data Account `json:data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return &result.Data, nil
}
