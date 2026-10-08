package guard

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/tool/goclaude/constant"
	"path/filepath"
	"regexp"
	"strings"
)

func pythonTarget(
	source string,
	token string,
) string {
	if strings.HasPrefix(token, "'") || strings.HasPrefix(token, `"`) {
		return filepath.Clean(token[1 : len(token)-1])
	}

	m := regexp.MustCompile(
		fmt.Sprintf(constant.PythonAssignment, regexp.QuoteMeta(token)),
	).FindStringSubmatch(source)

	if m == nil {
		return ""
	}

	return filepath.Clean(m[1])
}
