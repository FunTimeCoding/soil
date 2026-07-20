package web

import (
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/board"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
)

func rowLabels(v *board.Entry) []string {
	if v.Widget == constant.NextcloudWidget {
		return []string{constant.FilesLabel, constant.SharesLabel}
	}

	if v.Widget == constant.ArgocdWidget {
		return []string{
			constant.ApplicationsLabel,
			constant.OutOfSyncLabel,
			constant.DegradedLabel,
			constant.MissingLabel,
		}
	}

	var result []string

	for _, row := range v.Rows {
		result = append(result, row.Label)
	}

	return result
}
