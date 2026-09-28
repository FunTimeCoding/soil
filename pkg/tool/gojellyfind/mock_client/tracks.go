package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/item"

func (c *Client) Tracks(albumIdentifier string) ([]*item.Item, error) {
	var result []*item.Item

	for _, i := range c.items {
		if i.Type == "Audio" {
			result = append(result, i)
		}
	}

	return result, nil
}
