package gitlab

import (
	"context"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) SearchBlob(
	x context.Context,
	query string,
) ([]*gitlab.Blob, error) {
	var result []*gitlab.Blob
	projects, e := c.Projects(x)

	if e != nil {
		return nil, e
	}

	for _, p := range projects {
		page, _, f := c.client.Search.BlobsByProject(
			p.Identifier,
			query,
			&gitlab.SearchOptions{},
			gitlab.WithContext(x),
		)

		if f != nil {
			return nil, wrapError(f)
		}

		result = append(result, page...)
	}

	return result, nil
}
