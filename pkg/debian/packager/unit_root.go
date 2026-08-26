package packager

import (
	"github.com/funtimecoding/soil/pkg/debian/constant"
	system "github.com/funtimecoding/soil/pkg/system/constant"
	"github.com/funtimecoding/soil/pkg/system/join"
)

func unitRoot(packageDirectory string) string {
	return join.Absolute(
		packageDirectory,
		system.Library,
		constant.SystemdDirectory,
		constant.SystemDirectory,
	)
}
