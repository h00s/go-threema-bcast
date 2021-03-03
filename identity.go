package go_threema_bcast

import (
  "encoding/json"
  "time"
)

type Identities struct {
  Links []Link `json:"_links"`
  Identities []Identity `json:"identities"`
}

type Identity struct {
  Links []Link `json:"_links"`
  Id string `json:"id"`
  ThreemaId string `json:"threemaId"`
  ValidUntil time.Time `json:"validUntil"`
  Type string `json:"type"`
  RecipientLimit int `json:"recipientLimit"`
}

func (c *Client) GetIdentities() (Identities, error) {
  var i Identities

  body, err := c.Get("/identities")
  if err != nil {
    return i, err
  }

  if err := json.Unmarshal(body, &i); err != nil {
    return i, err
  }

  return i, nil
}
