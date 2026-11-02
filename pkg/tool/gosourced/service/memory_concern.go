package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/lint/concern"
	"github.com/funtimecoding/soil/pkg/lint/types/retarget_change"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/types/retarget_report"
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
)

func memoryConcern(
	report *retarget_report.Report,
	c *retarget_change.Change,
) *concern.Concern {
	memory := fmt.Sprintf("memory #%d %s", report.Identifier, report.Name)

	if c.Reason != "" {
		return concern.NewFile(
			constant.ConcernReferenceLeft,
			fmt.Sprintf("%s: %s: %s", memory, c.Before, c.Reason),
			"",
			true,
		)
	}

	return concern.NewFile(
		constant.ConcernRetargeted,
		fmt.Sprintf("%s: %s → %s", memory, c.Before, c.After),
		"",
		true,
	)
}
