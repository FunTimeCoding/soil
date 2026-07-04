package web

import "github.com/funtimecoding/go-library/pkg/tool/godashboardd/board"

func rowLabels(v *board.Entry) []string {
	if v.Widget == board.NextcloudWidget {
		return []string{filesLabel, sharesLabel}
	}

	if v.Widget == board.ArgocdWidget {
		return []string{
			applicationsLabel,
			outOfSyncLabel,
			degradedLabel,
			missingLabel,
		}
	}

	var result []string

	for _, row := range v.Rows {
		result = append(result, row.Label)
	}

	return result
}
