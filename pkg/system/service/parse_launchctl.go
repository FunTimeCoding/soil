package service

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func ParseLaunchctl(output string) map[string]string {
	result := make(map[string]string)

	for _, line := range strings.Split(output, stringsConstant.Unix) {
		fields := strings.Fields(line)

		if len(fields) != 3 || fields[0] == constant.LaunchctlHeader {
			continue
		}

		result[fields[2]] = launchState(fields[0], fields[1])
	}

	return result
}
