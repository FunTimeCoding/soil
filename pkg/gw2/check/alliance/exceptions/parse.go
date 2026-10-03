package exceptions

import (
	"github.com/funtimecoding/soil/pkg/gw2/constant"
	"github.com/funtimecoding/soil/pkg/notation"
	"github.com/funtimecoding/soil/pkg/system"
)

func Parse(path string) []*Exception {
	var result []*Exception
	notation.MustDecode(
		system.ReadFile(path, constant.ExceptionFile),
		&result,
		true,
	)

	return result
}
