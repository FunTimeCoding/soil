package census

import (
	"github.com/funtimecoding/soil/pkg/tool/gosourced/constant"
	"strings"
)

func keyPackage(key string) string {
	typed := key[:strings.LastIndex(key, constant.MemberSeparator)]

	return typed[:strings.LastIndex(typed, constant.MemberSeparator)]
}
