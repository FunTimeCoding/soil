package goquery

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/terminal"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
)

func printWords(
	t *terminal.Terminal,
	verb string,
	words []string,
) bool {
	if len(words) == 0 {
		return false
	}

	cut := len(words) > constant.WordListLimit

	if cut {
		words = append(words[:constant.WordListLimit], "...")
	}

	t.Line(fmt.Sprintf("    %s: %s", verb, join.Space(words...)))

	return cut
}
