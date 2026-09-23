package fixture

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	memoryConstant "github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"strings"
)

func Section(
	document string,
	heading string,
) string {
	start := strings.Index(document, heading)

	if start < 0 {
		return ""
	}

	rest := document[start+len(heading):]
	next := strings.Index(
		rest,
		join.Empty(constant.Unix, memoryConstant.SectionPrefix),
	)

	if next < 0 {
		return rest
	}

	return rest[:next]
}
