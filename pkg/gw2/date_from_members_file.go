package gw2

import (
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"strings"
)

func dateFromMembersFile(file string) string {
	last := strings.LastIndex(file, stringsConstant.Underscore)
	secondLast := strings.LastIndex(file[:last], stringsConstant.Underscore)

	return file[secondLast+1 : last]
}
