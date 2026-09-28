package mock_client

import "github.com/funtimecoding/soil/pkg/jellyfin/item"

func (c *Client) Episodes(seriesIdentifier string) ([]*item.Item, error) {
	var result []*item.Item

	for _, i := range c.items {
		if i.SeriesName != "" && i.Type == "Episode" {
			result = append(result, i)
		}
	}

	return result, nil
}
