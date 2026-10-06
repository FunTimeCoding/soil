package goquery

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/tool/goquery/constant"
)

func printWords(
	verb string,
	words []string,
) {
	if len(words) == 0 {
		return
	}

	if len(words) > constant.WordListLimit {
		words = append(words[:constant.WordListLimit], "...")
	}

	console.Format("    %s: %s\n", verb, join.Space(words...))
}
