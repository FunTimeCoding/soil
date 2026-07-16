package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
	"github.com/gpustack/gguf-parser-go/util/ptr"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) GroupProjects(identifier ...int64) ([]*project.Project, error) {
	var result []*gitlab.Project

	for _, i := range identifier {
		number := int64(1)

		for {
			page, _, e := c.client.Groups.ListGroupProjects(
				i,
				&gitlab.ListGroupProjectsOptions{
					ListOptions: gitlab.ListOptions{
						PerPage: constant.PerPage100,
						Page:    number,
					},
					Archived: ptr.Bool(false),
				},
			)

			if e != nil {
				return nil, e
			}

			result = append(result, page...)

			if int64(len(page)) < constant.PerPage100 {
				break
			}

			number++
		}
	}

	return project.NewSlice(result), nil
}
