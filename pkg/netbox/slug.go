package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"strings"
)

func slug(name string) string {
	result := strings.ToLower(strings.ReplaceAll(name, " ", "-"))

	return constant.NonSlug.ReplaceAllString(result, "")
}
