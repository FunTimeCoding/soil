package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) SearchBlob(query string) ([]*gitlab.Blob, error) {
	var result []*gitlab.Blob
	projects, e := c.Projects()

	if e != nil {
		return nil, e
	}

	for _, p := range projects {
		page, _, f := c.client.Search.BlobsByProject(
			p.Identifier,
			query,
			&gitlab.SearchOptions{},
		)

		if f != nil {
			return nil, wrapError(f)
		}

		result = append(result, page...)
	}

	return result, nil
}
