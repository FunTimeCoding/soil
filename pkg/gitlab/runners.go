package gitlab

import (
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) Runners(all bool) ([]*runner.Runner, error) {
	var result []*gitlab.Runner
	number := int64(1)

	for {
		var page []*gitlab.Runner
		var e error
		options := &gitlab.ListRunnersOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: constant.PerPage100,
				Page:    number,
			},
		}

		if all {
			page, _, e = c.allRunners(options)
		} else {
			page, _, e = c.ownerRunners(options)
		}

		if e != nil {
			return nil, wrapError(e)
		}

		result = append(result, page...)

		if int64(len(page)) < constant.PerPage100 {
			break
		}

		number++
	}

	return c.enrichRunners(runner.NewSlice(runner.Deduplicate(result)))
}
