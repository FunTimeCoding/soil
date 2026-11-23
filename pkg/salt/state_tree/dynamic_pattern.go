package state_tree

import (
	"github.com/funtimecoding/soil/pkg/salt/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"regexp"
	"strings"
)

func dynamicPattern(text string) *regexp.Regexp {
	pieces := constant.Expression.Split(text, -1)

	for i, p := range pieces {
		pieces[i] = regexp.QuoteMeta(p)
	}

	return regexp.MustCompile(join.Empty("^", strings.Join(pieces, ".*"), "$"))
}
