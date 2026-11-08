package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"slices"
	"strings"
)

func callerKind(
	path string,
	home string,
	file string,
) string {
	segments := strings.Split(path, "/")

	switch {
	case strings.HasSuffix(file, constant.TestFileSuffix):
		return constant.CallerTest
	case path == home:
		return constant.CallerInternal
	case slices.Contains(segments, constant.UnitSegment),
		slices.Contains(segments, constant.IntegrationSegment):

		return constant.CallerSupport
	}

	return constant.CallerProduction
}
