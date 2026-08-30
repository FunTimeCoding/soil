package web

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/types/board_entry"
)

func summary(entries []*board_entry.Entry) []string {
	projects := map[string]bool{}
	failed := 0

	for _, entry := range entries {
		projects[entry.Project] = true

		if entry.Status == constant.JobFail {
			failed++
		}
	}

	result := []string{fmt.Sprintf("%d projects", len(projects))}

	if failed > 0 {
		result = append(result, fmt.Sprintf("%d failed", failed))
	}

	return result
}
