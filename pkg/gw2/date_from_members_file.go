package gw2

import (
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func dateFromMembersFile(file string) string {
	last := strings.LastIndex(file, constant.Underscore)
	secondLast := strings.LastIndex(file[:last], constant.Underscore)

	return file[secondLast+1 : last]
}
