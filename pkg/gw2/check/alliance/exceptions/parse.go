package exceptions

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Parse(path string) []*Exception {
	var result []*Exception
	s := system.ReadFile(path, constant.ExceptionFile)

	if false {
		fmt.Printf("Parsing: %s\n", s)
	}

	notation.MustDecode(s, &result, true)

	return result
}
