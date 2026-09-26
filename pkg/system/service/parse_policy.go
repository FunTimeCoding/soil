package service

import (
	separator "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"strings"
)

func ParsePolicy(output string) map[string]*Policy {
	result := make(map[string]*Policy)
	var name string
	var installed bool

	for _, line := range strings.Split(output, separator.Unix) {
		trimmed := strings.TrimSpace(line)

		if trimmed == "" {
			continue
		}

		if !strings.HasPrefix(line, separator.Space) &&
			strings.HasSuffix(trimmed, separator.Colon) {
			name = strings.TrimSuffix(trimmed, separator.Colon)
			result[name] = &Policy{}
			installed = false

			continue
		}

		if name == "" {
			continue
		}

		if v, okay := strings.CutPrefix(
			trimmed,
			constant.PolicyInstalled,
		); okay {
			result[name].Version = strings.TrimSpace(v)

			continue
		}

		fields := strings.Fields(trimmed)

		if len(fields) < 2 {
			continue
		}

		if !priority(fields[0]) {
			installed = fields[0] == constant.PolicyCurrent

			continue
		}

		if installed &&
			result[name].Origin == "" &&
			fields[1] != constant.PolicyStatusPath {
			result[name].Origin = fields[1]
		}
	}

	return result
}
