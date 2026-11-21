package subject

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/constant"
)

func Upgrade(
	release string,
	source string,
) string {
	return join.Empty(
		constant.RemediationUpgrade,
		constant.SubjectSeparator,
		release,
		constant.SubjectSeparator,
		source,
	)
}
