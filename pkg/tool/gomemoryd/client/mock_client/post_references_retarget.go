package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
)

func (c *Client) PostReferencesRetarget(
	_ context.Context,
	moves *move_set.Set,
	_ bool,
) ([]*retarget_report.Report, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	c.Retargets = append(c.Retargets, moves)

	return []*retarget_report.Report{}, nil
}
