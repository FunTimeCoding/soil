package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/tree"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Tree(
	project int64,
	path string,
	reference string,
	recursive bool,
	limit int64,
) ([]*tree.Node, error) {
	if limit == 0 {
		limit = 100
	}

	o := &gitlab.ListTreeOptions{
		ListOptions: gitlab.ListOptions{PerPage: limit},
	}

	if path != "" {
		o.Path = &path
	}

	if reference != "" {
		o.Ref = &reference
	}

	if recursive {
		o.Recursive = &recursive
	}

	result, r, e := c.client.Repositories.ListTree(project, o)

	if r != nil && r.StatusCode == 404 {
		// Do not panic
		return []*tree.Node{}, nil
	}

	if e != nil {
		return nil, wrapError(e)
	}

	return tree.NewSlice(result), nil
}
