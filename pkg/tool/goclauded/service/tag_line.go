package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func tagLine(
	statistics *client.Statistics,
	loaded map[int64]bool,
) string {
	var parts []string

	for _, one := range statistics.Tags {
		if one.Identifiers == nil {
			continue
		}

		read := 0

		for _, identifier := range *one.Identifiers {
			if loaded[identifier] {
				read++
			}
		}

		parts = append(
			parts,
			fmt.Sprintf("%s %d/%d", one.Name, read, one.Count),
		)
	}

	if len(parts) == 0 {
		return ""
	}

	return join.Empty("tags     ", join.CommaSpace(parts))
}
