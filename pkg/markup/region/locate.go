package region

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/markup/constant"
	"strings"
)

func locate(
	content string,
	name string,
) (int, int, error) {
	start, end := markers(name)
	starts := strings.Count(content, start)
	ends := strings.Count(content, end)

	if starts == 0 && ends == 0 {
		return 0, 0, fmt.Errorf("%q: %w", name, constant.ErrorRegionMissing)
	}

	if starts > 1 {
		return 0, 0, fmt.Errorf(
			constant.RegionRepeated,
			name,
			constant.RegionStartWord,
			starts,
		)
	}

	if ends > 1 {
		return 0, 0, fmt.Errorf(
			constant.RegionRepeated,
			name,
			constant.RegionEndWord,
			ends,
		)
	}

	if starts == 0 {
		return 0, 0, fmt.Errorf(
			constant.RegionUnpaired,
			name,
			constant.RegionEndWord,
		)
	}

	if ends == 0 {
		return 0, 0, fmt.Errorf(
			constant.RegionUnpaired,
			name,
			constant.RegionStartWord,
		)
	}

	inner := strings.Index(content, start) + len(start)
	closing := strings.Index(content, end)

	if closing < inner {
		return 0, 0, fmt.Errorf(constant.RegionReversed, name)
	}

	return inner, closing, nil
}
