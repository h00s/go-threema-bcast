package go_threema_bcast

import (
	"encoding/json"
)

type Profile struct {
	Links []Link `json:"_links"`
}

func (c *Client) GetProfile() (*Profile, error) {
	p := &Profile{}

	body, err := c.Get("/")
	if err != nil {
		return p, err
	}

	return p.Unmarshal(body)
}

func (p *Profile) Unmarshal(body []byte) (*Profile, error) {
	if err := json.Unmarshal(body, &p); err != nil {
		return p, err
	}

	return p, nil
}
