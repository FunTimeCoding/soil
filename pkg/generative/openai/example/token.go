package example

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/console/scanner"
	"github.com/funtimecoding/soil/pkg/generative/token"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Token() {
	s := scanner.New()
	console.Line("Paste text and press Ctrl+D to finish:")
	lines := s.Scan()

	if false {
		for _, l := range lines {
			console.Format("Line: %+v\n", l)
		}
	}

	console.Format("Token count: %d\n", token.Count(join.NewLine(lines)))
}
