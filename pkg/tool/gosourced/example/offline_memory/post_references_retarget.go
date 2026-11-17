package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/lint/types/move_set"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
)

func (c *Client) PostReferencesRetarget(
	context.Context,
	*move_set.Set,
	bool,
) ([]*retarget_report.Report, error) {
	return nil, nil
}
