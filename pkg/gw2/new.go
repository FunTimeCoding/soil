package gw2

import "github.com/MrGunflame/gw2api"

// Reference: https://wiki.guildwars2.com/wiki/API:Main
func New(token string) *Client {
	client := gw2api.New()

	if token != "" {
		client = client.WithAccessToken(token)
	}

	return &Client{client: client}
}
