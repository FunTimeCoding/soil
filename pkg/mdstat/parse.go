package mdstat

import (
	"github.com/funtimecoding/soil/pkg/mdstat/array"
	"github.com/funtimecoding/soil/pkg/mdstat/constant"
	soilStrings "github.com/funtimecoding/soil/pkg/strings"
	"github.com/funtimecoding/soil/pkg/strings/split"
	"strings"
)

func Parse(content string) []*array.Array {
	var result []*array.Array
	var current *array.Array

	for _, line := range split.NewLine(content) {
		if m := constant.Header.FindStringSubmatch(line); m != nil {
			current = array.New(m[1], m[2], m[3])
			current.Failed = strings.Count(m[4], constant.FailedMarker)
			result = append(result, current)

			continue
		}

		if current == nil {
			continue
		}

		if m := constant.Members.FindStringSubmatch(line); m != nil {
			current.Required = soilStrings.MustToInteger(m[1])
			current.Active = soilStrings.MustToInteger(m[2])
		}

		if constant.Rebuild.MatchString(line) {
			current.Rebuilding = true
		}
	}

	return result
}
