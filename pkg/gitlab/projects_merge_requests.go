package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/merge_request"

func (c *Client) ProjectsMergeRequests() []*merge_request.Request {
	var result []*merge_request.Request

	for _, identifier := range c.projects {
		for _, e := range c.MustProjectMergeRequests(identifier, false) {
			if e.Done() {
				continue
			}

			result = append(result, e)
		}
	}

	return result
}
