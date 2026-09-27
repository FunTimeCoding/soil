package helper

import (
	"github.com/funtimecoding/soil/pkg/netbox/constant"
	"strings"
)

func ToWebLink(v string) string {
	return strings.Replace(v, constant.InterfaceSegment, "/", 1)
}
