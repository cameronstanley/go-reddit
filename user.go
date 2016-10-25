package reddit

func (c *Client) GetUserInfo(username string) (Account, error) {
	url := fmt.Sprintf("%s/user/%s/about.json", baseUrl, username)
	resp, err := c.http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo Account
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
