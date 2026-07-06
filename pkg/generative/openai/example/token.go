package example

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console/scanner"
	"github.com/funtimecoding/soil/pkg/generative/token"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func Token() {
	s := scanner.New()
	fmt.Println("Paste text and press Ctrl+D to finish:")
	lines := s.Scan()

	if false {
		for _, l := range lines {
			fmt.Printf("Line: %+v\n", l)
		}
	}

	fmt.Printf("Token count: %d\n", token.Count(join.NewLine(lines)))
}
