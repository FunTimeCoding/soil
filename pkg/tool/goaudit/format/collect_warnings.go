package format

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goaudit/scan"
)

func collectWarnings(services []*scan.Service) []string {
	var result []string
	maxName := 0

	for _, s := range services {
		if len(s.Concerns) > 0 && len(s.Name) > maxName {
			maxName = len(s.Name)
		}
	}

	for _, s := range services {
		for _, c := range s.Concerns {
			result = append(
				result,
				fmt.Sprintf("%-*s  %s", maxName, s.Name, c.Text),
			)
		}
	}

	return result
}
