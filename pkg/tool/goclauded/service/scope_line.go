package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func scopeLine(
	statistics *client.Statistics,
	loaded int,
) string {
	var parts []string

	for _, one := range statistics.Scopes {
		if one.Name == "" {
			parts = append(
				parts,
				fmt.Sprintf("%d/%d loaded", loaded, one.Count),
			)

			continue
		}

		parts = append(parts, fmt.Sprintf("%s %d", one.Name, one.Count))
	}

	return join.Empty("memory   ", join.CommaSpace(parts))
}
