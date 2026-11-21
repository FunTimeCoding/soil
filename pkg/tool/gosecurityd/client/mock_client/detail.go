package mock_client

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) detail(name string) *client.SystemDetail {
	if d, okay := c.details[name]; okay {
		return d
	}

	s, okay := c.snapshots[name]

	if !okay {
		return nil
	}

	return &client.SystemDetail{
		Summary: client.SystemSummary{
			Name:     name,
			Release:  s.Release,
			Coverage: constant.CoverageUntracked,
			Updated:  c.updated[name],
		},
		Findings: []*client.SystemFinding{},
	}
}
