package service

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/finding"
)

func unownedFinding(
	kind string,
	noun string,
	callsigns []string,
	rows int64,
) *finding.Finding {
	if rows == 0 {
		return nil
	}

	named := make([]string, 0, len(callsigns))

	for _, i := range callsigns {
		if i == "" {
			named = append(named, "(no callsign)")

			continue
		}

		named = append(named, i)
	}

	return finding.New(
		kind,
		"",
		fmt.Sprintf(
			"%d pending %s for %d callsigns no session holds: %s",
			rows,
			noun,
			len(named),
			join.CommaSpace(named),
		),
		rows,
	)
}
