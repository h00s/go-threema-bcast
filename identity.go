package go_threema_bcast

import "time"

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
