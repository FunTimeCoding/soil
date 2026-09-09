package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/pipeline"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) Pipelines(
	project int64,
	reference string,
	status string,
	limit int64,
) ([]*pipeline.Pipeline, error) {
	if limit == 0 {
		limit = constant.PerPage1000
	}

	o := &gitlab.ListProjectPipelinesOptions{
		ListOptions: gitlab.ListOptions{PerPage: limit},
	}

	if reference != "" {
		o.Ref = &reference
	}

	if status != "" {
		o.Status = new(gitlab.BuildStateValue(status))
	}

	result, _, e := c.client.Pipelines.ListProjectPipelines(project, o)

	if e != nil {
		return nil, wrapError(e)
	}

	return pipeline.NewSlice(result), nil
}
