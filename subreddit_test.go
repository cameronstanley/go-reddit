package reddit

import (
  "testing"
)

func TestGetDefaultSubreddits(t *testing.T) {
  client := new(Client)
  client.GetDefaultSubreddits()
}
