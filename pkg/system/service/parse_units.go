package service

import (
	separator "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func ParseUnits(output string) map[string]string {
	result := make(map[string]string)

	for _, line := range strings.Split(output, separator.Unix) {
		fields := strings.Fields(line)

		if len(fields) < 4 || fields[1] != constant.UnitLoaded ||
			!strings.HasSuffix(fields[0], constant.UnitSuffix) {
			continue
		}

		result[fields[0]] = unitState(fields[2])
	}

	return result
}
