package permission_page

import (
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"strings"
)

func reason(code string) string {
	return strings.TrimPrefix(code, constant.PermissionReasonPrefix)
}
