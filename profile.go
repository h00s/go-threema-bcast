package go_threema_bcast

import (
  "encoding/json"
)

type Profile struct {
  Links []Link `json:"_links"`
}

func (c *Client) GetProfile() (Profile, error) {
  var p Profile

  body, err := c.Get("/")
  if err != nil {
    return p, err
  }

  if err := json.Unmarshal(body, &p); err != nil {
    return p, err
  }

  return p, nil
}
