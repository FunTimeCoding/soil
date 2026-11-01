package layout

import "github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"

func (e *Entry) RowLabels() []string {
	if e.Widget == constant.NextcloudWidget {
		return []string{constant.FilesLabel, constant.SharesLabel}
	}

	if e.Widget == constant.ArgocdWidget {
		return []string{
			constant.ApplicationsLabel,
			constant.OutOfSyncLabel,
			constant.DegradedLabel,
			constant.MissingLabel,
		}
	}

	var result []string

	for _, v := range e.Rows {
		result = append(result, v.Label)
	}

	return result
}
