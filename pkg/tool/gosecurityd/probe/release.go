package probe

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"strings"
)

func release(lines []string) (string, error) {
	fields := map[string]string{}

	for _, line := range lines {
		if k, v, okay := strings.Cut(line, "="); okay {
			fields[k] = strings.Trim(v, `"'`)
		}
	}

	if len(fields) == 0 && len(lines) > 0 {
		if m := constant.RedHatPattern.FindStringSubmatch(lines[0]); m != nil {
			return join.Empty(strings.ToLower(m[1]), m[2]), nil
		}
	}

	identifier := fields[constant.ReleaseIdentifier]

	if identifier == constant.DebianIdentifier {
		if c := fields[constant.ReleaseCodename]; c != "" {
			return c, nil
		}

		m := constant.CodenamePattern.FindStringSubmatch(
			fields[constant.ReleaseVersion],
		)

		if m != nil {
			return m[1], nil
		}
	} else if identifier != "" {
		major, _, _ := strings.Cut(
			fields[constant.ReleaseNumber],
			constant.VersionSeparator,
		)

		if major != "" {
			return join.Empty(identifier, major), nil
		}
	}

	return "", fmt.Errorf("release not recognized: %s", join.Space(lines...))
}
