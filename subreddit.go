package reddit

import (
  "fmt"
  "io/ioutil"
  "net/http"
)

type Subreddit struct {
  
}

func (c *Client) GetDefaultSubreddits() {
  resp, _ := http.Get("http://reddit.com/subreddits/default.json")
  defer resp.Body.Close()
  body, _ := ioutil.ReadAll(resp.Body)
  fmt.Print(string(body))
}
