package reddit

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
)

func (c *Client) IsUsernameAvailable(username string) (bool, error) {
	url := fmt.Sprintf("%s/api/username_available.json?user=%s", baseUrl, username)
	resp, err := c.http.Get(url)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	isUsernameAvailable, err := strconv.ParseBool(string(bs))
	if err != nil {
		return false, err
	}

	return isUsernameAvailable, err
}

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
