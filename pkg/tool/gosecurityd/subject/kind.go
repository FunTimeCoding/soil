package subject

import (
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
	"github.com/funtimecoding/soil/pkg/vulnerability/identifier"
	"strings"
)

func Kind(s string) string {
	if identifier.Check(s) == nil {
		return constant.RemediationVulnerability
	}

	prefix, rest, okay := strings.Cut(s, constant.SubjectSeparator)

	if !okay || rest == "" {
		return ""
	}

	switch prefix {
	case constant.RemediationRelease,
		constant.RemediationReboot,
		constant.RemediationKernel:

		return prefix
	case constant.RemediationUpgrade:
		release, source, found := strings.Cut(rest, constant.SubjectSeparator)

		if found && release != "" && source != "" {
			return prefix
		}
	}

	return ""
}
