package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/types/retarget_change"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func retargetConcern(
	path string,
	c *retarget_change.Change,
) *concern.Concern {
	if c.Reason != "" {
		return concern.NewLine(
			constant.ConcernReferenceLeft,
			fmt.Sprintf("%s: %s", c.Before, c.Reason),
			path,
			c.Line,
			"",
			true,
		)
	}

	return concern.NewLine(
		constant.ConcernRetargeted,
		fmt.Sprintf("%s → %s", c.Before, c.After),
		path,
		c.Line,
		"",
		true,
	)
}
