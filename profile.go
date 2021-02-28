package go_threema_bcast

import (
  "encoding/json"
)

type Profile struct {
  links Links
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
