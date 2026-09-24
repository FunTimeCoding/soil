package coverage

import (
	"github.com/funtimecoding/soil/pkg/crap/constant"
	library "github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func Parse(output string) map[string]float64 {
	result := map[string]float64{}

	for _, line := range split.NewLine(output) {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, constant.TotalPrefix) {
			continue
		}

		fields := strings.Fields(line)

		if len(fields) < 3 {
			continue
		}

		key := strings.TrimSuffix(fields[0], ":")
		result[key] = library.MustToFloat(
			strings.TrimSuffix(fields[len(fields)-1], constant.Percent),
		)
	}

	return result
}
