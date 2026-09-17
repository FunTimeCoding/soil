package pointer

import (
	"github.com/funtimecoding/soil/pkg/lint/constant"
	"strings"
)

func RouteLiteral(route string) string {
	cut := len(route)

	if i := strings.Index(route, "{"); i != -1 {
		cut = i
	}

	if i := strings.Index(route, constant.RouteEllipsis); i != -1 && i < cut {
		cut = i
	}

	return route[:cut]
}
