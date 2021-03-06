package go_threema_bcast

import (
	"encoding/json"
)

type Identities struct {
	Links      []Link     `json:"_links"`
	Identities []Identity `json:"identities"`
}

type Identity struct {
	Links          []Link   `json:"_links"`
	Id             string   `json:"id"`
	ThreemaId      string   `json:"threemaId"`
	ValidUntil     DateTime `json:"validUntil"`
	Type           string   `json:"type"`
	RecipientLimit int      `json:"recipientLimit"`
}

func (c *Client) GetIdentities() (*Identities, error) {
	i := &Identities{}

	body, err := c.Get("/identities")
	if err != nil {
		return i, err
	}

	return i.Unmarshal(body)
}

func (p *Identities) Unmarshal(body []byte) (*Identities, error) {
	if err := json.Unmarshal(body, &p); err != nil {
		return p, err
	}

	return p, nil
}
